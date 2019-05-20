
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

// ModifyInstanceVncPasswd_GatedLaunch invokes the ecs.ModifyInstanceVncPasswd_GatedLaunch API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancevncpasswd_gatedlaunch.html
func (client *Client) ModifyInstanceVncPasswd_GatedLaunch(request *ModifyInstanceVncPasswd_GatedLaunchRequest) (response *ModifyInstanceVncPasswd_GatedLaunchResponse, err error) {
response = CreateModifyInstanceVncPasswd_GatedLaunchResponse()
err = client.DoAction(request, response)
return
}

// ModifyInstanceVncPasswd_GatedLaunchWithChan invokes the ecs.ModifyInstanceVncPasswd_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancevncpasswd_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceVncPasswd_GatedLaunchWithChan(request *ModifyInstanceVncPasswd_GatedLaunchRequest) (<-chan *ModifyInstanceVncPasswd_GatedLaunchResponse, <-chan error) {
responseChan := make(chan *ModifyInstanceVncPasswd_GatedLaunchResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyInstanceVncPasswd_GatedLaunch(request)
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

// ModifyInstanceVncPasswd_GatedLaunchWithCallback invokes the ecs.ModifyInstanceVncPasswd_GatedLaunch API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancevncpasswd_gatedlaunch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceVncPasswd_GatedLaunchWithCallback(request *ModifyInstanceVncPasswd_GatedLaunchRequest, callback func(response *ModifyInstanceVncPasswd_GatedLaunchResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyInstanceVncPasswd_GatedLaunchResponse
var err error
defer close(result)
response, err = client.ModifyInstanceVncPasswd_GatedLaunch(request)
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

// ModifyInstanceVncPasswd_GatedLaunchRequest is the request struct for api ModifyInstanceVncPasswd_GatedLaunch
type ModifyInstanceVncPasswd_GatedLaunchRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
                    VncPassword     string `position:"Query" name:"VncPassword"`
}


// ModifyInstanceVncPasswd_GatedLaunchResponse is the response struct for api ModifyInstanceVncPasswd_GatedLaunch
type ModifyInstanceVncPasswd_GatedLaunchResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceVncPasswd_GatedLaunchRequest creates a request to invoke ModifyInstanceVncPasswd_GatedLaunch API
func CreateModifyInstanceVncPasswd_GatedLaunchRequest() (request *ModifyInstanceVncPasswd_GatedLaunchRequest) {
request = &ModifyInstanceVncPasswd_GatedLaunchRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceVncPasswd_GatedLaunch", "ecs", "openAPI")
return
}

// CreateModifyInstanceVncPasswd_GatedLaunchResponse creates a response to parse from ModifyInstanceVncPasswd_GatedLaunch response
func CreateModifyInstanceVncPasswd_GatedLaunchResponse() (response *ModifyInstanceVncPasswd_GatedLaunchResponse) {
response = &ModifyInstanceVncPasswd_GatedLaunchResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


