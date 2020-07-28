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


async function publishSdk(task, options) {
  const gitPath = path.join(task.output, task.gitCodePath);
  await publish.publishSdkToGithub(task, gitPath, options, true);
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
  task.codeGenPath = path.join(output, `${task.dirPrefix}-${task.productId.toLowerCase()}`);
  const origin = path.join(output, path.basename(config['sdkCoreDir']));
  // check php syntax
  const phpFiles = await glob(`${task.codeGenPath}/**/*.php`);
  await pMap(phpFiles, async file => {
    return await exec(`php -l ${file}`, options);
  }, { concurrency: config.sdkCheckConcurrency });
  // copy generate code to git path
  await exec(`cp -rf ${task.codeGenPath} ${origin}`, options);
}

async function checkSdkScene(output, options) {
  // TODO: add scene test case
}

testRuning();
