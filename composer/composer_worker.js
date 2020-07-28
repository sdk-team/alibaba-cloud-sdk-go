const config = require('./setting');
const publish = require('./publish');
const consumer = require('./consumer');
const producer = require('./producer');
const ossutil = require('./oss_util');
const fs = require('fs');
const path = require('path');
const cp = require('child_process');
const promisify = require('util').promisify;
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
function camlCase(name) {
  var arr = name.match(/./g);
  var flag = false;
  for (let i = 0; i < arr.length; i++) {
    var c1 = arr[i];
    if (/[a-z]/.test(c1) && flag) {
      arr[i] = arr[i].toUpperCase();
      flag = false;
    } else if (c1 === '-' || c1 === '_') {
      flag = true;
      arr[i] = '';
    }
  }
  name = arr.join('');
  return name[0].toUpperCase() + name.substring(1);
}

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



exports.conflictDeal = async function (task, options, err) {
  await console.log('start conflictdealing.....');
  await exec('git reset HEAD^', options);
  await exec('git checkout CHANGELOG VERSION', options);
  await console.log('end conflictdealing......');
  return;
}




async function checkSdkSyntax(task, options) {
  const output = task.output;
  const origin = path.join(output, task.gitCodePath, '/src');
  const productId = camlCase(task.productId);
  const codeGenDir = path.join(output, `${task.gitCodePath}-src`);
  task.codeGenPath = path.join(codeGenDir, productId);
  // copy generate code to git path
  await exec(`rm -rf ${origin}`, options);
  await exec(`cp -rf ${codeGenDir} ${origin}`, options);
  // check sdk syntax
  const phpFiles = await glob(`${origin}/**/*.php`);
  await pMap(phpFiles, async file => {
    return await exec(`php -l ${file}`, options);
  }, { concurrency: config['sdkCheckConcurrency'] });
}

async function publishSdk(task, options) {
  await publish.publishSdkToGithub(task, options, true);
}

async function getReleaseTag(task) {
  const origin = path.join(task.output, task.gitCodePath, '/src');
  const releaseFile = path.join(origin, 'Release.php');
  const releaseContent = await readFile(releaseFile, 'utf8');
  const version = releaseContent.match(/const VERSION = '(\S*)';/)[1];
  const newVersionArray = version.split('.');
  ++newVersionArray[2];
  const newVersion = newVersionArray.join('.');
  await writeFile(
    releaseFile,
    releaseContent.replace(`'${version}'`, '\'' + newVersion + '\''),
    'utf8'
  );
  task.version = newVersion;
  return task.version;
}

async function conflictDeal(task,options) {
  await exec('git reset HEAD^', options);
  await exec('git checkout CHANGELOG.md src/Release.php', options);
}

async function updateChangeLog(task, gitPath) {
  const changelogFile = gitPath + '/CHANGELOG.md';
  const log = task.changeLog || '- Generated ' + task.apiVersions.join(', ') + ' for `' + task.productId + '`.';
  const changelogFileContent = await readFile(changelogFile, 'utf8');
  const date = new Date();
  const month = date.getMonth() + 1;
  const dataString = date.getFullYear() + '-' + month + '-' + date.getDate();
  let content = '# CHANGELOG\n'
    + '\n'
    + '## ' + task.version + ' - ' + dataString + '\n'
    + log
    + '\n';
  await writeFile(
    changelogFile,
    changelogFileContent.replace(/# CHANGELOG/g, content),
    'utf8'
  );
  return log;
}

async function checkSdkScene(output, options) {
  // TODO: add scene test case
}

testRuning();
