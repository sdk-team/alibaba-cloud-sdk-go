
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

// CloneDBInstance invokes the rds.CloneDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/clonedbinstance.html
func (client *Client) CloneDBInstance(request *CloneDBInstanceRequest) (response *CloneDBInstanceResponse, err error) {
response = CreateCloneDBInstanceResponse()
err = client.DoAction(request, response)
return
}

// CloneDBInstanceWithChan invokes the rds.CloneDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/clonedbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloneDBInstanceWithChan(request *CloneDBInstanceRequest) (<-chan *CloneDBInstanceResponse, <-chan error) {
responseChan := make(chan *CloneDBInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CloneDBInstance(request)
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

// CloneDBInstanceWithCallback invokes the rds.CloneDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/clonedbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloneDBInstanceWithCallback(request *CloneDBInstanceRequest, callback func(response *CloneDBInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CloneDBInstanceResponse
var err error
defer close(result)
response, err = client.CloneDBInstance(request)
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

// CloneDBInstanceRequest is the request struct for api CloneDBInstance
type CloneDBInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    ResourceGroupId     string `position:"Query" name:"ResourceGroupId"`
                    DBInstanceDescription     string `position:"Query" name:"DBInstanceDescription"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    RestoreTime     string `position:"Query" name:"RestoreTime"`
                    Period     string `position:"Query" name:"Period"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    BackupId     string `position:"Query" name:"BackupId"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    UsedTime     string `position:"Query" name:"UsedTime"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    DbNames     string `position:"Query" name:"DbNames"`
                    VSwitchId     string `position:"Query" name:"VSwitchId"`
                    PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
                    VPCId     string `position:"Query" name:"VPCId"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    PayType     string `position:"Query" name:"PayType"`
                    InstanceNetworkType     string `position:"Query" name:"InstanceNetworkType"`
}


// CloneDBInstanceResponse is the response struct for api CloneDBInstance
type CloneDBInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     string `json:"OrderId" xml:"OrderId"`
            ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
            Port     string `json:"Port" xml:"Port"`
}

// CreateCloneDBInstanceRequest creates a request to invoke CloneDBInstance API
func CreateCloneDBInstanceRequest() (request *CloneDBInstanceRequest) {
request = &CloneDBInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CloneDBInstance", "rds", "openAPI")
return
}

// CreateCloneDBInstanceResponse creates a response to parse from CloneDBInstance response
func CreateCloneDBInstanceResponse() (response *CloneDBInstanceResponse) {
response = &CloneDBInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


