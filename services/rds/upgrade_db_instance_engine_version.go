
package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpgradeDBInstanceEngineVersion invokes the rds.UpgradeDBInstanceEngineVersion API synchronously
// api document: https://help.aliyun.com/api/rds/upgradedbinstanceengineversion.html
func (client *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
response = CreateUpgradeDBInstanceEngineVersionResponse()
err = client.DoAction(request, response)
return
}

// UpgradeDBInstanceEngineVersionWithChan invokes the rds.UpgradeDBInstanceEngineVersion API asynchronously
// api document: https://help.aliyun.com/api/rds/upgradedbinstanceengineversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeDBInstanceEngineVersionWithChan(request *UpgradeDBInstanceEngineVersionRequest) (<-chan *UpgradeDBInstanceEngineVersionResponse, <-chan error) {
responseChan := make(chan *UpgradeDBInstanceEngineVersionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.UpgradeDBInstanceEngineVersion(request)
if err != nil {
errChan <- err
} else {
responseChan <- response
}
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

// UpgradeDBInstanceEngineVersionWithCallback invokes the rds.UpgradeDBInstanceEngineVersion API asynchronously
// api document: https://help.aliyun.com/api/rds/upgradedbinstanceengineversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeDBInstanceEngineVersionWithCallback(request *UpgradeDBInstanceEngineVersionRequest, callback func(response *UpgradeDBInstanceEngineVersionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *UpgradeDBInstanceEngineVersionResponse
var err error
defer close(result)
response, err = client.UpgradeDBInstanceEngineVersion(request)
callback(response, err)
result <- 1
})
if err != nil {
defer close(result)
callback(nil, err)
result <- 0
}
return result
}

// UpgradeDBInstanceEngineVersionRequest is the request struct for api UpgradeDBInstanceEngineVersion
type UpgradeDBInstanceEngineVersionRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    EffectiveTime     string `position:"Query" name:"EffectiveTime"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// UpgradeDBInstanceEngineVersionResponse is the response struct for api UpgradeDBInstanceEngineVersion
type UpgradeDBInstanceEngineVersionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TaskId     string `json:"TaskId" xml:"TaskId"`
}

// CreateUpgradeDBInstanceEngineVersionRequest creates a request to invoke UpgradeDBInstanceEngineVersion API
func CreateUpgradeDBInstanceEngineVersionRequest() (request *UpgradeDBInstanceEngineVersionRequest) {
request = &UpgradeDBInstanceEngineVersionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceEngineVersion", "", "")
return
}

// CreateUpgradeDBInstanceEngineVersionResponse creates a response to parse from UpgradeDBInstanceEngineVersion response
func CreateUpgradeDBInstanceEngineVersionResponse() (response *UpgradeDBInstanceEngineVersionResponse) {
response = &UpgradeDBInstanceEngineVersionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


