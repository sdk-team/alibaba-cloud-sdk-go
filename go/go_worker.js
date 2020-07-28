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
const moment = require('moment');

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
    await preBuild(task, config['sdkCoreDir'], options);
    // Step2: 校验代码


    await checkSdkSyntax(task, output, options);
    // Step3: 场景化测试（问题场景发起实际的调用测试验证）

    await checkSdkScene(output, options);
    // Step4: 发布至 Github, php composer 钩子在 github，因此仅需推送至 github 即可

    if (task.type === 'download') {
      await publish.publishToOssForDownload(task, output, options);
    } else {
      await publishSdk(task, output, options);
    }
    ossutil.uploadDir(output, task.output, options);
    await console.log('success!');
    console.log('-'.repeat(20));
  }
}


async function publishSdk(task, options) {
  const gitPath = path.join(config['output'], task.taskId,'github.com/aliyun/alibaba-cloud-sdk-go');
  await console.log(gitPath);
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
  //const { ctx } = this;
  await console.log('start checksdksyntax......');
  const output = task.output;
  const productId = task.productId.toLowerCase();
  task.codeGenPath = path.join(output, 'github.com/aliyun/alibaba-cloud-sdk-go/services', productId);
  // format code
  await exec('gofmt -s -w [!e]*', { cwd: `${path.join(output, 'github.com')}` });
  // copy generate code to git path
  try {
    await exec(`rm -rf ${path.join(output, task.gitCodePath)}services/${productId}/*`, options);
  } catch (err) {
    //不存在目录的错误只记录即可
    ctx.logger.info(err);
  }
  await exec(`cp -rf ${task.codeGenPath} ${path.join(output,`${this.gitCodePath}services/`)}`);
  await console.log('end checksdksyntax......');
  // const goSdkPath = path.join(output, 'go/src/github.com/aliyun/alibaba-cloud-sdk-go/');
  // const data = await exec(`go vet ./services/${productDir}`,
  //   Object.assign({
  //     cwd: goSdkPath,
  //     env: {
  //       GOCACHE: path.join(output, 'Library/Caches/go-build'),
  //       HOME: output,
  //       GOPATH: path.join(output, 'go'),
  //       PATH: process.env.PATH,
  //     }
  //   })
  // );
}

async function preBuild(task, options) {
  await console.log('start prebuild......');
  const output = config['output']+task.taskId;
  await console.log(output);
  await exec(`cp -rf ${config['sdkCoreDir']} ${output}`, options);
  const cwd = path.join(output, path.basename(config['sdkCoreDir']));
  await console.log(cwd);
  await exec('git fetch origin master && ' +
    'git reset --hard origin/master && ' +
    'git checkout master',
  Object.assign({ cwd }, options));
  await console.log('end prebuild......');
}


async function checkSdkScene(output, options) {
  // TODO: add scene test case
}

testRuning();
