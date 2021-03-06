
package bssopenapi

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

// SetUserQuota invokes the bssopenapi.SetUserQuota API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/setuserquota.html
func (client *Client) SetUserQuota(request *SetUserQuotaRequest) (response *SetUserQuotaResponse, err error) {
response = CreateSetUserQuotaResponse()
err = client.DoAction(request, response)
return
}

// SetUserQuotaWithChan invokes the bssopenapi.SetUserQuota API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/setuserquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserQuotaWithChan(request *SetUserQuotaRequest) (<-chan *SetUserQuotaResponse, <-chan error) {
responseChan := make(chan *SetUserQuotaResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SetUserQuota(request)
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

// SetUserQuotaWithCallback invokes the bssopenapi.SetUserQuota API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/setuserquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserQuotaWithCallback(request *SetUserQuotaRequest, callback func(response *SetUserQuotaResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SetUserQuotaResponse
var err error
defer close(result)
response, err = client.SetUserQuota(request)
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

// SetUserQuotaRequest is the request struct for api SetUserQuota
type SetUserQuotaRequest struct {
*requests.RpcRequest
                    Uid     requests.Integer `position:"Query" name:"Uid"`
                    Amount     string `position:"Query" name:"Amount"`
                    OutBizId     string `position:"Query" name:"OutBizId"`
                    Currency     string `position:"Query" name:"Currency"`
                    Bid     string `position:"Query" name:"Bid"`
}


// SetUserQuotaResponse is the response struct for api SetUserQuota
type SetUserQuotaResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
            Success     bool `json:"Success" xml:"Success"`
            Data     bool `json:"Data" xml:"Data"`
}

// CreateSetUserQuotaRequest creates a request to invoke SetUserQuota API
func CreateSetUserQuotaRequest() (request *SetUserQuotaRequest) {
request = &SetUserQuotaRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("BssOpenApi", "2017-12-14", "SetUserQuota", "", "")
return
}

// CreateSetUserQuotaResponse creates a response to parse from SetUserQuota response
func CreateSetUserQuotaResponse() (response *SetUserQuotaResponse) {
response = &SetUserQuotaResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


