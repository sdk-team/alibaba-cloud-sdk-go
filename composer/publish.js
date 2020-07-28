const config = require('./setting');
const fs = require('fs');
const path = require('path');
const cp = require('child_process');
const promisify = require('util').promisify;
const exists = promisify(fs.exists);
const exec = promisify(cp.exec);
const readFile = promisify(fs.readFile);
const writeFile = promisify(fs.writeFile);
const mkdir = promisify(fs.mkdir);
const moment = require('moment');
const worker = require('./cpp_worker');
const ossutil = require('./oss_util');

async function msleep(n) {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve();
    }, n);
  });
};

exports.clone = async function (repo, sdk, url) {
  if (await exists(sdk)) {
    // await exec('git pull origin master', { cwd: sdk });
    return;
  }
  await mkdir(repo, { recursive: true });
  console.log('git clone');
  console.log(repo);
  console.log(`git clone ${url} --depth 1 --no-checkout ${sdk}`);
  // clone sdk repo
  await exec(`git clone ${url} --depth 1 --no-checkout ${sdk}`, { cwd: repo });
  console.log('--------');
  // set git config
  await Promise.all([
    //exec('git config user.email "sdk-team@alibabacloud.com"', { cwd: sdk }),
    //exec('git config user.name "Alibaba Cloud SDK"', { cwd: sdk })
  ]);
}

async function preBuild(task, options) {
  const output = task.output;
  const gitCodePath = path.join(output, this.gitCodePath);
  await exec(`rm -rf ${gitCodePath} && cp -rf ${this.coreDir} ${gitCodePath}`, options);
  await exec('git fetch origin master && ' +
    'git reset --hard origin/master && ' +
    'git checkout master',
  Object.assign({ cwd: gitCodePath }, options));
  //composer的生成器，要将output中所有目录里加入 dirname.php 的文件，单独建立一个dir作output避免干扰同时发布的其他语言生成目录
  await exec(`cp -rf ${path.join(gitCodePath, 'src')} ${path.join(output, `${this.gitCodePath}-src`)}`);
}

exports.publishSdkToGithub = async function (task, gitPath, options, isRelease = false) {
  try {
    await console.log('start publishing sdk to github');
    await release(task, gitPath, isRelease, options);
    const result = await exec('git rev-parse HEAD',
      Object.assign({ cwd: gitPath }, options));//commit id
    const commitid = result.stdout.replace('\n', '');
    const addresses = [
      `https://github.com/${task.gitbub_username}/${task.language}/commits/master`,
      `https://github.com/${task.gitbub_username}/${task.language}/commit/${commitid}`
    ];
    console.log(addresses);
    await console.log('end publishing sdk to github');
    // await Promise.all([
    //   await releaseTask.addSuccessPackageReleaseTask(task, addresses[0], task.language, 'github'),
    //   await releaseTask.addSuccessPackageReleaseTask(task, addresses[1], task.language, 'git_diff'),
    //   sdkList.updateReleaseLog({ commit_url: addresses[0] }, task.taskId, task.language),
    // ]);
  } catch (err) {
    // ctx.logger.error(err);
    console.log(err);
    const address = `/task/${task.redisKey}/github?type=history`;
    console.log(address);
    //await this.updatePublishError(task, 'github', address, err);
  }
}

exports.publishSdkToOss = async function (task, targetPath) {
  const targetName = path.basename(targetPath);
  // ctx.logger.info(`push oss [${targetName}]: ${targetPath}`);
  const filepath = path.join(config['ossDir'],'/',targetName);
  await ossutil.upload(targetPath,filepath);
 // await ossutil.upload(targetPath, `${config['ossDir']}/${targetName}`);
  const address = `https://${config['oss_bucket']}.${config['oss_region']}.aliyuncs.com/${config['ossDir']}/${targetName}`;
  console.log(address);
  //await releaseTask.addSuccessPackageReleaseTask(task, address, this.language, 'oss');
}

exports.publishToOssForDownload = async function (task, codeGenPath, options) {
  const productId = task.productId.toLowerCase();
  const targetName = `u-${task.taskId}-${task.dirPrefix}-${productId}.zip`;
  await ossutil.uploadZipDir(codeGenPath, targetName, 'downloads/', options);
  const address = `https://${config['oss_bucket']}.${config['oss_region']}.aliyuncs.com/downloads/${targetName}`;
  console.log(address);
  //await releaseTask.addSuccessPackageReleaseTask(task, address, this.language);
}

async function release(task, gitPath, release = false, options) {
  const { ctx } = this;
  options = Object.assign({ cwd: gitPath }, options);
  const repo = path.join(task.gitbub_username, task.repository);
  // Make remote address to token style for auth, token should from KMS.
  const { stdout: remote } = await exec('git remote -v', options);
  if (!~remote.indexOf(task.token)) {

    if (~remote.indexOf('origin')) {
      await exec('git remote rm origin', options);
    }
    await exec('git config user.name sdk-team && git config user.email sdk-team@alibabacloud.com', options);
    await exec(`git remote add origin https://sdk-team:${task.token}@github.com/${repo}.git`, options);
  }
  let tag, log;
  let i = 5;
  while (i-- > 0) {
    try {
      //确保代码最新
      try {
        await exec('git stash', options);
        await exec('git fetch origin master && git rebase origin/master', options);
        await exec('git stash pop', options);
      } catch (e) {
        // this.ctx.logger.error(e);
      }
      tag = await getReleaseTag(task);
      log = await updateChangeLog(task, gitPath);
      // For commit message invalid
      let commit = log.split('\n')[0].toString().replace(/`/g, '').replace('- ', '');
      // Push
      await gitPushCode(commit, options, task);
      break;
    } catch (err) {
      console.log(err);
      // ctx.logger.error(err);
      if (i === 0) {
        throw err;
      }
      if (err.stderr && ~err.stderr.indexOf('.git/index.lock')) {
        // 因为index.lock报错,暂停30s
        await msleep(30 * 1000);
        continue;
      }
      await worker.conflictDeal(task, options, err);
      // 处理完冲突等15s再试
      await msleep(15 * 1000);
    }
  }
  //await console.log(release);
  if (!release) { return; }

  try {
    // Delete local tag
    await exec(`git tag -d ${tag}`, options);
    // Delete remote tag for generate success every time
    await exec(`git push origin --delete ${tag}`, options);
    // Delete remote release
    await exec('curl -v -i -X DELETE -H "Content-Type:application/json" -H' +
      `"Authorization: token ${task.token}" https://api.github.com/repos/${repo}/releases/${tag}'`, options);
  } catch (e) {
    console.log(e);
    // this.ctx.logger.error(e);
  }

  // Release with tag
  i = 6;
  while (i-- > 0) {
    try {
      //确保能打上tag
     // await console.log(ctx);
     await exec('curl  -d"{"tag_name":"tag","target_commitish":"master","name":"tag","body":"log","draft":"false","prerelease":"false"}" ' +
     `-H "Content-Type:application/json" -H "Authorization: token ${task.token}" -X POST  "https://api.github.com/repos/${repo}/releases"`, options);
    //  await ctx.curl(`https://api.github.com/repos/${repo}/releases`, {
    //    method: 'POST',
    //    headers: {
    //      'Authorization': `token ${task.token}`,
    //      'Content-Type': 'application/json'
    //    },
    //    data: {
    //      tag_name: tag,
    //      target_commitish: 'master',
    //      name: tag,
    //      body: log,
    //      draft: false,
    //      prerelease: false
    //    }
     //});
      break;
    } catch (err) {
      console.log(err);
      if (i === 0) {
        throw err;
      }
      continue;
    }
  }
}

async function gitPushCode(commit, options) {
  await console.log('start gitpushcode');
  await exec('git add .', options);
  await exec(`git commit -m "${commit}"`, options);
  await exec('git push origin master', options);
  await console.log('end gitpushcode');
}

async function updateChangeLog(task, gitPath) {
  await console.log('start updatechangelog..');
  const changelog = path.join(gitPath, 'CHANGELOG');
  const log = task.changeLog || '- Generated ' + task.apiVersions.join(', ') + ' for `' + task.productId + '`.';
  const changelogFileContent = await readFile(changelog, 'utf8');
  let content = `${moment().format('YYYY-MM-DD')} Version: ${task.version}\n`
    + log + '\n\n'
    + changelogFileContent;
  await writeFile(changelog, content, 'utf8');
  await console.log('end updatechangelog');
  return log;
}

async function getReleaseTag(task) {
  await console.log('start getreleasetag..');
  return `${task.productId}-${task.version}`;
}
