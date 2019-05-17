
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

// ModifyReplicaVerificationMode invokes the rds.ModifyReplicaVerificationMode API synchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicaverificationmode.html
func (client *Client) ModifyReplicaVerificationMode(request *ModifyReplicaVerificationModeRequest) (response *ModifyReplicaVerificationModeResponse, err error) {
response = CreateModifyReplicaVerificationModeResponse()
err = client.DoAction(request, response)
return
}

// ModifyReplicaVerificationModeWithChan invokes the rds.ModifyReplicaVerificationMode API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicaverificationmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaVerificationModeWithChan(request *ModifyReplicaVerificationModeRequest) (<-chan *ModifyReplicaVerificationModeResponse, <-chan error) {
responseChan := make(chan *ModifyReplicaVerificationModeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyReplicaVerificationMode(request)
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

// ModifyReplicaVerificationModeWithCallback invokes the rds.ModifyReplicaVerificationMode API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicaverificationmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaVerificationModeWithCallback(request *ModifyReplicaVerificationModeRequest, callback func(response *ModifyReplicaVerificationModeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyReplicaVerificationModeResponse
var err error
defer close(result)
response, err = client.ModifyReplicaVerificationMode(request)
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

// ModifyReplicaVerificationModeRequest is the request struct for api ModifyReplicaVerificationMode
type ModifyReplicaVerificationModeRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ReplicaId     string `position:"Query" name:"ReplicaId"`
                    VerificationMode     string `position:"Query" name:"VerificationMode"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// ModifyReplicaVerificationModeResponse is the response struct for api ModifyReplicaVerificationMode
type ModifyReplicaVerificationModeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReplicaVerificationModeRequest creates a request to invoke ModifyReplicaVerificationMode API
func CreateModifyReplicaVerificationModeRequest() (request *ModifyReplicaVerificationModeRequest) {
request = &ModifyReplicaVerificationModeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaVerificationMode", "", "")
return
}

// CreateModifyReplicaVerificationModeResponse creates a response to parse from ModifyReplicaVerificationMode response
func CreateModifyReplicaVerificationModeResponse() (response *ModifyReplicaVerificationModeResponse) {
response = &ModifyReplicaVerificationModeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


