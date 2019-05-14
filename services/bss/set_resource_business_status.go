
package bss

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

// SetResourceBusinessStatus invokes the bss.SetResourceBusinessStatus API synchronously
// api document: https://help.aliyun.com/api/bss/setresourcebusinessstatus.html
func (client *Client) SetResourceBusinessStatus(request *SetResourceBusinessStatusRequest) (response *SetResourceBusinessStatusResponse, err error) {
response = CreateSetResourceBusinessStatusResponse()
err = client.DoAction(request, response)
return
}

// SetResourceBusinessStatusWithChan invokes the bss.SetResourceBusinessStatus API asynchronously
// api document: https://help.aliyun.com/api/bss/setresourcebusinessstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetResourceBusinessStatusWithChan(request *SetResourceBusinessStatusRequest) (<-chan *SetResourceBusinessStatusResponse, <-chan error) {
responseChan := make(chan *SetResourceBusinessStatusResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SetResourceBusinessStatus(request)
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

// SetResourceBusinessStatusWithCallback invokes the bss.SetResourceBusinessStatus API asynchronously
// api document: https://help.aliyun.com/api/bss/setresourcebusinessstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetResourceBusinessStatusWithCallback(request *SetResourceBusinessStatusRequest, callback func(response *SetResourceBusinessStatusResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SetResourceBusinessStatusResponse
var err error
defer close(result)
response, err = client.SetResourceBusinessStatus(request)
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

// SetResourceBusinessStatusRequest is the request struct for api SetResourceBusinessStatus
type SetResourceBusinessStatusRequest struct {
*requests.RpcRequest
                    BusinessStatus     string `position:"Query" name:"BusinessStatus"`
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceId     string `position:"Query" name:"ResourceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    ResourceType     string `position:"Query" name:"ResourceType"`
}


// SetResourceBusinessStatusResponse is the response struct for api SetResourceBusinessStatus
type SetResourceBusinessStatusResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateSetResourceBusinessStatusRequest creates a request to invoke SetResourceBusinessStatus API
func CreateSetResourceBusinessStatusRequest() (request *SetResourceBusinessStatusRequest) {
request = &SetResourceBusinessStatusRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Bss", "2014-07-14", "SetResourceBusinessStatus", "bss", "openAPI")
return
}

// CreateSetResourceBusinessStatusResponse creates a response to parse from SetResourceBusinessStatus response
func CreateSetResourceBusinessStatusResponse() (response *SetResourceBusinessStatusResponse) {
response = &SetResourceBusinessStatusResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


