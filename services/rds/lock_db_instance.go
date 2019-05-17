
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

// LockDBInstance invokes the rds.LockDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/lockdbinstance.html
func (client *Client) LockDBInstance(request *LockDBInstanceRequest) (response *LockDBInstanceResponse, err error) {
response = CreateLockDBInstanceResponse()
err = client.DoAction(request, response)
return
}

// LockDBInstanceWithChan invokes the rds.LockDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/lockdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LockDBInstanceWithChan(request *LockDBInstanceRequest) (<-chan *LockDBInstanceResponse, <-chan error) {
responseChan := make(chan *LockDBInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.LockDBInstance(request)
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

// LockDBInstanceWithCallback invokes the rds.LockDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/lockdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LockDBInstanceWithCallback(request *LockDBInstanceRequest, callback func(response *LockDBInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *LockDBInstanceResponse
var err error
defer close(result)
response, err = client.LockDBInstance(request)
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

// LockDBInstanceRequest is the request struct for api LockDBInstance
type LockDBInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    LockReason     string `position:"Query" name:"LockReason"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// LockDBInstanceResponse is the response struct for api LockDBInstance
type LockDBInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateLockDBInstanceRequest creates a request to invoke LockDBInstance API
func CreateLockDBInstanceRequest() (request *LockDBInstanceRequest) {
request = &LockDBInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "LockDBInstance", "", "")
return
}

// CreateLockDBInstanceResponse creates a response to parse from LockDBInstance response
func CreateLockDBInstanceResponse() (response *LockDBInstanceResponse) {
response = &LockDBInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


