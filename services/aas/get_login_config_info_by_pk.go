
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

// GetLoginConfigInfoByPK invokes the aas.GetLoginConfigInfoByPK API synchronously
// api document: https://help.aliyun.com/api/aas/getloginconfiginfobypk.html
func (client *Client) GetLoginConfigInfoByPK(request *GetLoginConfigInfoByPKRequest) (response *GetLoginConfigInfoByPKResponse, err error) {
response = CreateGetLoginConfigInfoByPKResponse()
err = client.DoAction(request, response)
return
}

// GetLoginConfigInfoByPKWithChan invokes the aas.GetLoginConfigInfoByPK API asynchronously
// api document: https://help.aliyun.com/api/aas/getloginconfiginfobypk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLoginConfigInfoByPKWithChan(request *GetLoginConfigInfoByPKRequest) (<-chan *GetLoginConfigInfoByPKResponse, <-chan error) {
responseChan := make(chan *GetLoginConfigInfoByPKResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetLoginConfigInfoByPK(request)
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

// GetLoginConfigInfoByPKWithCallback invokes the aas.GetLoginConfigInfoByPK API asynchronously
// api document: https://help.aliyun.com/api/aas/getloginconfiginfobypk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLoginConfigInfoByPKWithCallback(request *GetLoginConfigInfoByPKRequest, callback func(response *GetLoginConfigInfoByPKResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetLoginConfigInfoByPKResponse
var err error
defer close(result)
response, err = client.GetLoginConfigInfoByPK(request)
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

// GetLoginConfigInfoByPKRequest is the request struct for api GetLoginConfigInfoByPK
type GetLoginConfigInfoByPKRequest struct {
*requests.RpcRequest
                    PK     string `position:"Query" name:"PK"`
}


// GetLoginConfigInfoByPKResponse is the response struct for api GetLoginConfigInfoByPK
type GetLoginConfigInfoByPKResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PK     string `json:"PK" xml:"PK"`
            MFADeviceId     string `json:"MFADeviceId" xml:"MFADeviceId"`
            MFAChecked     bool `json:"MFAChecked" xml:"MFAChecked"`
            ForceLoginCheckMFA     bool `json:"ForceLoginCheckMFA" xml:"ForceLoginCheckMFA"`
            SkipMFANotAllowed     bool `json:"SkipMFANotAllowed" xml:"SkipMFANotAllowed"`
}

// CreateGetLoginConfigInfoByPKRequest creates a request to invoke GetLoginConfigInfoByPK API
func CreateGetLoginConfigInfoByPKRequest() (request *GetLoginConfigInfoByPKRequest) {
request = &GetLoginConfigInfoByPKRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Aas", "2015-07-01", "GetLoginConfigInfoByPK", "", "")
return
}

// CreateGetLoginConfigInfoByPKResponse creates a response to parse from GetLoginConfigInfoByPK response
func CreateGetLoginConfigInfoByPKResponse() (response *GetLoginConfigInfoByPKResponse) {
response = &GetLoginConfigInfoByPKResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


