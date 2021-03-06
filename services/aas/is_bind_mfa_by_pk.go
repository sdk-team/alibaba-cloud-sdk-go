
package aas

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

// IsBindMfaByPk invokes the aas.IsBindMfaByPk API synchronously
// api document: https://help.aliyun.com/api/aas/isbindmfabypk.html
func (client *Client) IsBindMfaByPk(request *IsBindMfaByPkRequest) (response *IsBindMfaByPkResponse, err error) {
response = CreateIsBindMfaByPkResponse()
err = client.DoAction(request, response)
return
}

// IsBindMfaByPkWithChan invokes the aas.IsBindMfaByPk API asynchronously
// api document: https://help.aliyun.com/api/aas/isbindmfabypk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IsBindMfaByPkWithChan(request *IsBindMfaByPkRequest) (<-chan *IsBindMfaByPkResponse, <-chan error) {
responseChan := make(chan *IsBindMfaByPkResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.IsBindMfaByPk(request)
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

// IsBindMfaByPkWithCallback invokes the aas.IsBindMfaByPk API asynchronously
// api document: https://help.aliyun.com/api/aas/isbindmfabypk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IsBindMfaByPkWithCallback(request *IsBindMfaByPkRequest, callback func(response *IsBindMfaByPkResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *IsBindMfaByPkResponse
var err error
defer close(result)
response, err = client.IsBindMfaByPk(request)
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

// IsBindMfaByPkRequest is the request struct for api IsBindMfaByPk
type IsBindMfaByPkRequest struct {
*requests.RpcRequest
                    PK     string `position:"Query" name:"PK"`
}


// IsBindMfaByPkResponse is the response struct for api IsBindMfaByPk
type IsBindMfaByPkResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            IsBindMfa     bool `json:"IsBindMfa" xml:"IsBindMfa"`
}

// CreateIsBindMfaByPkRequest creates a request to invoke IsBindMfaByPk API
func CreateIsBindMfaByPkRequest() (request *IsBindMfaByPkRequest) {
request = &IsBindMfaByPkRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Aas", "2015-07-01", "IsBindMfaByPk", "", "")
return
}

// CreateIsBindMfaByPkResponse creates a response to parse from IsBindMfaByPk response
func CreateIsBindMfaByPkResponse() (response *IsBindMfaByPkResponse) {
response = &IsBindMfaByPkResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


