
package yundun

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

// AddCNameWaf invokes the yundun.AddCNameWaf API synchronously
// api document: https://help.aliyun.com/api/yundun/addcnamewaf.html
func (client *Client) AddCNameWaf(request *AddCNameWafRequest) (response *AddCNameWafResponse, err error) {
response = CreateAddCNameWafResponse()
err = client.DoAction(request, response)
return
}

// AddCNameWafWithChan invokes the yundun.AddCNameWaf API asynchronously
// api document: https://help.aliyun.com/api/yundun/addcnamewaf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCNameWafWithChan(request *AddCNameWafRequest) (<-chan *AddCNameWafResponse, <-chan error) {
responseChan := make(chan *AddCNameWafResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AddCNameWaf(request)
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

// AddCNameWafWithCallback invokes the yundun.AddCNameWaf API asynchronously
// api document: https://help.aliyun.com/api/yundun/addcnamewaf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCNameWafWithCallback(request *AddCNameWafRequest, callback func(response *AddCNameWafResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AddCNameWafResponse
var err error
defer close(result)
response, err = client.AddCNameWaf(request)
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

// AddCNameWafRequest is the request struct for api AddCNameWaf
type AddCNameWafRequest struct {
*requests.RpcRequest
                    InstanceId     string `position:"Query" name:"InstanceId"`
                    Domain     string `position:"Query" name:"Domain"`
                    InstanceType     string `position:"Query" name:"InstanceType"`
}


// AddCNameWafResponse is the response struct for api AddCNameWaf
type AddCNameWafResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                    WafInfoList WafInfoListInAddCNameWaf `json:"WafInfoList" xml:"WafInfoList"`
}

// CreateAddCNameWafRequest creates a request to invoke AddCNameWaf API
func CreateAddCNameWafRequest() (request *AddCNameWafRequest) {
request = &AddCNameWafRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Yundun", "2015-04-16", "AddCNameWaf", "yundun", "openAPI")
return
}

// CreateAddCNameWafResponse creates a response to parse from AddCNameWaf response
func CreateAddCNameWafResponse() (response *AddCNameWafResponse) {
response = &AddCNameWafResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


