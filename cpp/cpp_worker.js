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
    await publish.clone(config['repo'], config['sdkCoreDir'], `https://github.com/${task.gitbub_username}/${task.repository}.git`);
    // Step1：进行预构建
    await publish.preBuild(task, config['sdkCoreDir'], options);
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
  const productId = task.productId.toLowerCase();
  const origin = path.join(output, path.basename(config['sdkCoreDir']));
  task.codeGenPath = path.join(output, productId);
  // copy generate code to git path
  await exec(`rm -rf ${origin}/${productId}/src/model/*`, options);
  await exec(`rm -rf ${origin}/${productId}/include/alibabacloud/${productId}/model/*`, options);
  await exec(`cp -rf ${task.codeGenPath} ${origin}`, options);
}

async function checkSdkScene(output, options) {
  // TODO: add scene test case
}



testRuning();