
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

// ModifyFCTrigger invokes the rds.ModifyFCTrigger API synchronously
// api document: https://help.aliyun.com/api/rds/modifyfctrigger.html
func (client *Client) ModifyFCTrigger(request *ModifyFCTriggerRequest) (response *ModifyFCTriggerResponse, err error) {
response = CreateModifyFCTriggerResponse()
err = client.DoAction(request, response)
return
}

// ModifyFCTriggerWithChan invokes the rds.ModifyFCTrigger API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyfctrigger.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFCTriggerWithChan(request *ModifyFCTriggerRequest) (<-chan *ModifyFCTriggerResponse, <-chan error) {
responseChan := make(chan *ModifyFCTriggerResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyFCTrigger(request)
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

// ModifyFCTriggerWithCallback invokes the rds.ModifyFCTrigger API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyfctrigger.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFCTriggerWithCallback(request *ModifyFCTriggerRequest, callback func(response *ModifyFCTriggerResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyFCTriggerResponse
var err error
defer close(result)
response, err = client.ModifyFCTrigger(request)
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

// ModifyFCTriggerRequest is the request struct for api ModifyFCTrigger
type ModifyFCTriggerRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    TriggerId     string `position:"Query" name:"TriggerId"`
                    SubscriptionObjects  *[]string `position:"Query" name:"SubscriptionObjects"  type:"Repeated"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    EventFormat     string `position:"Query" name:"EventFormat"`
                    Retry     requests.Integer `position:"Query" name:"Retry"`
                    TriggerArn     string `position:"Query" name:"TriggerArn"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    Concurrency     requests.Integer `position:"Query" name:"Concurrency"`
                    InvocationRoleArn     string `position:"Query" name:"InvocationRoleArn"`
                    FunctionArn     string `position:"Query" name:"FunctionArn"`
}


// ModifyFCTriggerResponse is the response struct for api ModifyFCTrigger
type ModifyFCTriggerResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyFCTriggerRequest creates a request to invoke ModifyFCTrigger API
func CreateModifyFCTriggerRequest() (request *ModifyFCTriggerRequest) {
request = &ModifyFCTriggerRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "ModifyFCTrigger", "rds", "openAPI")
return
}

// CreateModifyFCTriggerResponse creates a response to parse from ModifyFCTrigger response
func CreateModifyFCTriggerResponse() (response *ModifyFCTriggerResponse) {
response = &ModifyFCTriggerResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


