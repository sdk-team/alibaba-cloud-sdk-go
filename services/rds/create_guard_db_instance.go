
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

// CreateGuardDBInstance invokes the rds.CreateGuardDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/createguarddbinstance.html
func (client *Client) CreateGuardDBInstance(request *CreateGuardDBInstanceRequest) (response *CreateGuardDBInstanceResponse, err error) {
response = CreateCreateGuardDBInstanceResponse()
err = client.DoAction(request, response)
return
}

// CreateGuardDBInstanceWithChan invokes the rds.CreateGuardDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createguarddbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGuardDBInstanceWithChan(request *CreateGuardDBInstanceRequest) (<-chan *CreateGuardDBInstanceResponse, <-chan error) {
responseChan := make(chan *CreateGuardDBInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateGuardDBInstance(request)
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

// CreateGuardDBInstanceWithCallback invokes the rds.CreateGuardDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createguarddbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGuardDBInstanceWithCallback(request *CreateGuardDBInstanceRequest, callback func(response *CreateGuardDBInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateGuardDBInstanceResponse
var err error
defer close(result)
response, err = client.CreateGuardDBInstance(request)
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

// CreateGuardDBInstanceRequest is the request struct for api CreateGuardDBInstance
type CreateGuardDBInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ConnectionStringPrefix     string `position:"Query" name:"ConnectionStringPrefix"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    Port     string `position:"Query" name:"Port"`
}


// CreateGuardDBInstanceResponse is the response struct for api CreateGuardDBInstance
type CreateGuardDBInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            GuardDBInstanceId     string `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
            ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
            Port     string `json:"Port" xml:"Port"`
}

// CreateCreateGuardDBInstanceRequest creates a request to invoke CreateGuardDBInstance API
func CreateCreateGuardDBInstanceRequest() (request *CreateGuardDBInstanceRequest) {
request = &CreateGuardDBInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2013-05-28", "CreateGuardDBInstance", "", "")
return
}

// CreateCreateGuardDBInstanceResponse creates a response to parse from CreateGuardDBInstance response
func CreateCreateGuardDBInstanceResponse() (response *CreateGuardDBInstanceResponse) {
response = &CreateGuardDBInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


