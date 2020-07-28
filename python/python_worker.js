const config = require('./setting');
const publish = require('./publish');
const consumer = require('./consumer');
const producer = require('./producer');
const ossutil = require('./oss_util');
const fs = require('fs');
const path = require('path');
const cp = require('child_process');
const promisify = require('util').promisify;
const readdir = promisify(fs.readdir);
const exec = promisify(cp.exec);
const readFile = promisify(fs.readFile);
const writeFile = promisify(fs.writeFile);
const exists = promisify(fs.exists);
const glob = promisify(require('glob'));
const pMap = require('p-map');
// publishStatus: {
//   waitCheck: 'wait_check',
//   getMetaData: 'get_meta_data',
//   enqueue: 'in_queue',
//   preBuilding: 'pre_building',
//   building: 'building',
//   checkSyntax: 'checking_syntax',
//   sceneTest: 'scene_test',
//   publishSdk: 'publish_sdk',
//   done: 'done',
//   error: 'error'
// }


async function testRuning() {
  while (true) {
    const task = await consumer.getTask();
    console.log('-'.repeat(20));
    console.log(task);
    console.log('-'.repeat(20));
    let options = task.baseExecOptions;
    options.env = Object.assign(options.env, process.env);
    let output = `${config['output']}`;
    await ossutil.downloadDir(task.output, output, options);
    output = `${output}${task.taskId}`;
    await console.log(output);
    producer.produce(JSON.stringify({
    type: 'status',
    redisKey: task.redisKey,
    status: 'pre_building',     }));
    //let options = task.baseExecOptions;
    options.env = Object.assign(options.env, process.env);
    await publish.clone(config['repo'], config['sdkCoreDir'], `https://github.com/${task.gitbub_username}/${task.repository}.git`);
    // Step1：进行预构建
    await publish.preBuild(task, config['sdkCoreDir'], options);
    // Step2: 校验代码
    producer.produce(JSON.stringify({
              type: 'status',
              redisKey: task.redisKey,
              status: 'checking_syntax',
            }));

    await checkSdkSyntax(task, output, options);
    // Step3: 场景化测试（问题场景发起实际的调用测试验证）
    producer.produce(JSON.stringify({
              type: 'status',
              redisKey: task.redisKey,
              status: 'scene_test',
            }));

    await checkSdkScene(output, options);
    // Step4: 发布至 Github, php composer 钩子在 github，因此仅需推送至 github 即可
    producer.produce(JSON.stringify({
              type: 'status',
              redisKey: task.redisKey,
              status: 'publish_sdk',
            }));

    if (task.type === 'download') {
      await publish.publishToOssForDownload(task, output, options);
    } else {
      await publishSdk(task, output, options);
    }
    producer.produce(JSON.stringify({
              type: 'status',
              redisKey: task.redisKey,
              status: 'done',
            }));
    task.status = 2;
    producer.produce(JSON.stringify({
              type: 'result',
              task,
            }));
    ossutil.uploadDir(output, task.output, options);
    await console.log('success!');
    console.log('-'.repeat(20));
  }
}


async function publishSdk(task, output,options) {
  await console.log('start publish sdk');
  //const output = task.output;
  await console.log(output);
  //task.codeGenPath = path.join(output,task.codeGenPath);
  await console.log(task.codeGenPath);
  await exec(`cp -rf ${task.codeGenPath} ${task.codeGenPath}-pypi`, Object.assign({ cwd: output }, options));
  await Promise.all([
    publish.publishSdkToGithub(task, options),
    publishSdkToPypi(task, options),
  ]);
}

async function publishSdkToPypi(task, options) {
  //const { ctx } = this;
  //const { sdkList } = ctx.service;
  //const { releaseTask } = ctx.service;
  try {
    let name = `${this.dirPrefix}-${task.productId.toLowerCase()}`;
    const gitPath = `${task.codeGenPath}-pypi`;
    if (config['publishToTestRepo']) {
      const tmp = path.join(gitPath, 'setup.py');
      let setup = await readFile(tmp, 'utf8');
      setup = setup.replace(name, `${name}-test`);
      name = `${name}-test`;
      await writeFile(tmp, setup);
    }
    const pack = 'python3 setup.py sdist 2>&1';
    await exec(pack, Object.assign({ cwd: gitPath }, options));
    const push = `twine upload dist/* -u ${config.sdkToken.pypi}`;
    //ctx.logger.info(`push ${gitPath}: ${push}`);
    await exec(push, Object.assign({ cwd: gitPath }, options));
    //const address = `https://pypi.org/project/${name}`;
    //await releaseTask.addSuccessPackageReleaseTask(task, address, this.language, 'pypi');
    //sdkList.updateReleaseLog({url: address}, task.taskId, this.language);
  } catch (err) {
    //ctx.logger.error(err);
    //const address = `/task/${task.redisKey}/pypi?type=history`;
    //await this.updatePublishError(task, 'pypi', address, err);
  }
}


exports.conflictDeal = async function (task, options, err) {
  await console.log('start conflictdealing.....');
  await exec('git reset HEAD^', options);
  await exec('git checkout CHANGELOG VERSION', options);
  await console.log('end conflictdealing......');
  return;
}




async function checkSdkSyntax(task, output, options) {
  await console.log('start checksdksyntax.....');
  //const output = task.output;
  const productId = task.productId.toLowerCase();
  task.codeGenPath = path.join(output, `${task.dirPrefix}-${productId}`);
  const origin = path.join(output, path.basename(config['sdkCoreDir']));
  // apply sdk version
  let dirs = await readdir(task.codeGenPath);
  let sdkDir = dirs.filter(dir => /aliyunsdk\S+/.test(dir))[0];
  let versionFile = path.join(task.codeGenPath, sdkDir, '__init__.py');
  await writeFile(versionFile, `__version__ = '${task.version}'`);
  // check python syntax
  const pythonFiles = await glob(`${task.codeGenPath}/**/*.py`);
  await pMap(pythonFiles, async file => {
    return await exec(`python3 -m py_compile ${file}`, options);
  }, { concurrency: config['sdkCheckConcurrency'] });
  // copy generate code to git path
  await exec(`rm -rf ${origin}/${this.dirPrefix}-${productId}/aliyunsdk${productId.replace(/-/g, '_')}/*`, options);
  await exec(`cp -rf ${task.codeGenPath} ${origin}`, options);
  await console.log('end checksdksyntax.....');

}




async function checkSdkScene(output, options) {
  // TODO: add scene test case
}

testRuning();
