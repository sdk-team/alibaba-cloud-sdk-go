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
const xml2js = require('xml2js');
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


async function publishSdk(task, output,options) {
  //const output = task.output;
  await console.log('start publish sdk');
  await console.log(output);
  console.log(task.codeGenPath);
  await exec(`cp -rf ${output} ${output}-oss`, Object.assign({ cwd: output }, options));
  await exec(`cp -rf ${output} ${output}-nuget`, Object.assign({ cwd: output }, options));
  await Promise.all([
    await publish.publishSdkToGithub(task, options),
    await publishSdkToNuget(task, options),
    await publishSdkToOss(task, options)
  ]);
  await console.log('end publish sdk');
}


exports.conflictDeal = async function (task, options, err) {
  await console.log('start conflictdealing.....');
  await exec('git reset HEAD^', options);
  await exec('git checkout CHANGELOG VERSION', options);
  await console.log('end conflictdealing......');
  return;
}




async function checkSdkSyntax(task, options) {
  await console.log('start checksdksyntax........');
  const output = task.output;
  const productId = task.productId.toLowerCase();
  task.codeGenPath = path.join(output, `${task.dirPrefix}-${productId}`);
  const origin = path.join(output, path.basename(config['sdkCoreDir']));
  // apply version
  const csproj = path.join(task.codeGenPath, `${task.dirPrefix}-${productId}.vs2017.csproj`);
  const json = await parse(await readFile(csproj));
  if (config.publishToTestRepo) {
    json.Project.PropertyGroup[0].AssemblyName[0] = `${json.Project.PropertyGroup[0].AssemblyName[0]}-test`;
  }
  
  json.Project.PropertyGroup[0].Version[0] = task.version;
  task.assemblyName = json.Project.PropertyGroup[0].AssemblyName[0];
  await console.log('task.assemblyName = '+task.assemblyName);
  const builder = new xml2js.Builder();
  const newCsproj = builder.buildObject(json);
  await writeFile(csproj, newCsproj);
  // copy generate code to git path
  await exec(`rm -rf ${origin}/${this.dirPrefix}-${productId}/${task.productId.replace(/-/g, '_')}/*`, options);
  await exec(`cp -rf ${task.codeGenPath} ${origin}`, options);
  // check c# build status
  await exec(`dotnet build ${task.dirPrefix}-${productId}.vs2017.csproj`,
    Object.assign({ cwd: path.join(origin, `aliyun-net-sdk-${productId}`) }, options));
  await console.log('end checksdksyntax..........');
}

async function publishSdkToNuget(task, options) {
 //const { ctx } = this;
 // const { sdkList } = ctx.service;
 // const { releaseTask } = ctx.service;
  try {
    //const encrypt = await releaseTask.getToken(config.sdkToken.nuget);
    //const token = (await kms.decrypt(encrypt)).Plaintext;
    await console.log('start publish sdk to nuget');
    const productId = task.productId.toLowerCase();
    const nugetPath = `${task.codeGenPath}-nuget`;
    await console.log('nugetPath:'+nugetPath);
    const remove = `dotnet remove aliyun-net-sdk-${productId}.vs2017.csproj reference ../${task.gitCodePath}/aliyun-net-sdk-core/aliyun-net-sdk-core.vs2017.csproj`;
    await console.log('exec (remove, Object.assign({cwd : nugetPath},options))');
    await exec(remove, Object.assign({ cwd: nugetPath }, options));
    const add = `dotnet add aliyun-net-sdk-${productId}.vs2017.csproj package aliyun-net-sdk-core`;
    await console.log('exec(add, Object.assign({ cwd: nugetPath }, options))');
    await exec(add, Object.assign({ cwd: nugetPath }, options));
    const packModule = `dotnet add aliyun-net-sdk-${productId}.vs2017.csproj package NuGet.Build.Tasks.Pack --version 5.1.0`;
    await console.log('await exec(packModule, Object.assign({ cwd: nugetPath }, options));');
    await exec(packModule, Object.assign({ cwd: nugetPath }, options));
    const pack = 'dotnet pack --configuration Release --output ./ aliyun-net-sdk-*.vs2017.csproj';
    await console.log('await exec(pack, Object.assign({ cwd: nugetPath }, options));');
    await exec(pack, Object.assign({ cwd: nugetPath }, options));
    const push = `dotnet nuget push aliyun-net-sdk-*.*.nupkg --source nuget.org -k ${task.token}`;
    //ctx.logger.info(`push ${nugetPath}: ${push}`);
    await console.log('await exec(push, Object.assign({ cwd: nugetPath }, options));');
    await exec(push, Object.assign({ cwd: nugetPath }, options));
    //const address = `https://www.nuget.org/packages/${task.assemblyName}/`;
    //await releaseTask.addSuccessPackageReleaseTask(task, address, task.language, 'nuget');
    //sdkList.updateReleaseLog({ url: address }, task.taskId, task.language);
  } catch (err) {
    //ctx.logger.error(err);
    await console.log(err);
    const address = `/task/${task.redisKey}/nuget?type=history`;
    //await this.updatePublishError(task, 'nuget', address, err);

  }
  await console.log('end publish sdk to nuget');

}

async function publishSdkToOss(task, options) {
  await console.log('start publish sdk to oss');
  //const { ctx } = this;
  try {
    const productId = task.productId.toLowerCase();
    const targetName = `aliyun-net-sdk-${productId}-${task.version}.zip`;
    const dllFile = `${task.assemblyName}.dll`;
    const pdbFile = `${task.assemblyName}.pdb`;
    const ossPath = `${task.codeGenPath}-oss`;
    const targetPath = path.join(ossPath, targetName);
    const remove = `dotnet remove aliyun-net-sdk-${productId}.vs2017.csproj reference ../${task.gitCodePath}/aliyun-net-sdk-core/aliyun-net-sdk-core.vs2017.csproj`;
    await exec(remove, Object.assign({ cwd: ossPath }, options));
    const add = `dotnet add aliyun-net-sdk-${productId}.vs2017.csproj package aliyun-net-sdk-core`;
    await exec(add, Object.assign({ cwd: ossPath }, options));
    const pack = 'dotnet build --configuration Release --output ./ aliyun-net-sdk-*.vs2017.csproj'
      + ` && zip -r ${targetName} ${dllFile} ${pdbFile}`;
    //ctx.logger.info(`pack c# sdk: ${pack}`);
    await exec(pack, Object.assign({ cwd: ossPath }, options));
    //await super.publishSdkToOss(task, targetPath);
  } catch (err) {
    //ctx.logger.error(err);
    await console.log(err);
    const address = `/task/${task.redisKey}/oss?type=history`;
    //await this.updatePublishError(task, 'oss', address, err);
  }
  await console.log('end publish sdk to nuget');

}


async function checkSdkScene(output, options) {
  // TODO: add scene test case
}

async function parse(xml) {
  return new Promise((resolve, reject) => {
    xml2js.parseString(xml, function (err, result) {
      if (err) {
        return reject(err);
      }
      return resolve(result);
    });
  });
}



testRuning();
