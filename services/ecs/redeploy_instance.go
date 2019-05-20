
package ecs

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

// RedeployInstance invokes the ecs.RedeployInstance API synchronously
// api document: https://help.aliyun.com/api/ecs/redeployinstance.html
func (client *Client) RedeployInstance(request *RedeployInstanceRequest) (response *RedeployInstanceResponse, err error) {
response = CreateRedeployInstanceResponse()
err = client.DoAction(request, response)
return
}

// RedeployInstanceWithChan invokes the ecs.RedeployInstance API asynchronously
// api document: https://help.aliyun.com/api/ecs/redeployinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RedeployInstanceWithChan(request *RedeployInstanceRequest) (<-chan *RedeployInstanceResponse, <-chan error) {
responseChan := make(chan *RedeployInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.RedeployInstance(request)
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

// RedeployInstanceWithCallback invokes the ecs.RedeployInstance API asynchronously
// api document: https://help.aliyun.com/api/ecs/redeployinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RedeployInstanceWithCallback(request *RedeployInstanceRequest, callback func(response *RedeployInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *RedeployInstanceResponse
var err error
defer close(result)
response, err = client.RedeployInstance(request)
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

// RedeployInstanceRequest is the request struct for api RedeployInstance
type RedeployInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ForceStop     requests.Boolean `position:"Query" name:"ForceStop"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// RedeployInstanceResponse is the response struct for api RedeployInstance
type RedeployInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TaskId     string `json:"TaskId" xml:"TaskId"`
}

// CreateRedeployInstanceRequest creates a request to invoke RedeployInstance API
func CreateRedeployInstanceRequest() (request *RedeployInstanceRequest) {
request = &RedeployInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "RedeployInstance", "ecs", "openAPI")
return
}

// CreateRedeployInstanceResponse creates a response to parse from RedeployInstance response
func CreateRedeployInstanceResponse() (response *RedeployInstanceResponse) {
response = &RedeployInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


