
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

// SecureCheck invokes the yundun.SecureCheck API synchronously
// api document: https://help.aliyun.com/api/yundun/securecheck.html
func (client *Client) SecureCheck(request *SecureCheckRequest) (response *SecureCheckResponse, err error) {
response = CreateSecureCheckResponse()
err = client.DoAction(request, response)
return
}

// SecureCheckWithChan invokes the yundun.SecureCheck API asynchronously
// api document: https://help.aliyun.com/api/yundun/securecheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SecureCheckWithChan(request *SecureCheckRequest) (<-chan *SecureCheckResponse, <-chan error) {
responseChan := make(chan *SecureCheckResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SecureCheck(request)
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

// SecureCheckWithCallback invokes the yundun.SecureCheck API asynchronously
// api document: https://help.aliyun.com/api/yundun/securecheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SecureCheckWithCallback(request *SecureCheckRequest, callback func(response *SecureCheckResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SecureCheckResponse
var err error
defer close(result)
response, err = client.SecureCheck(request)
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

// SecureCheckRequest is the request struct for api SecureCheck
type SecureCheckRequest struct {
*requests.RpcRequest
                    JstOwnerId     requests.Integer `position:"Query" name:"JstOwnerId"`
                    InstanceIds     string `position:"Query" name:"InstanceIds"`
}


// SecureCheckResponse is the response struct for api SecureCheck
type SecureCheckResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            RecentInstanceId     string `json:"RecentInstanceId" xml:"RecentInstanceId"`
                    ProblemList ProblemList `json:"ProblemList" xml:"ProblemList"`
                    NoProblemList NoProblemList `json:"NoProblemList" xml:"NoProblemList"`
                    NoScanList NoScanList `json:"NoScanList" xml:"NoScanList"`
                    ScanningList ScanningList `json:"ScanningList" xml:"ScanningList"`
                    InnerIpList InnerIpList `json:"InnerIpList" xml:"InnerIpList"`
}

// CreateSecureCheckRequest creates a request to invoke SecureCheck API
func CreateSecureCheckRequest() (request *SecureCheckRequest) {
request = &SecureCheckRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Yundun", "2015-04-16", "SecureCheck", "yundun", "openAPI")
return
}

// CreateSecureCheckResponse creates a response to parse from SecureCheck response
func CreateSecureCheckResponse() (response *SecureCheckResponse) {
response = &SecureCheckResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


