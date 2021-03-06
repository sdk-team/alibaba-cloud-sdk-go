
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

// OpenPortScan invokes the yundun.OpenPortScan API synchronously
// api document: https://help.aliyun.com/api/yundun/openportscan.html
func (client *Client) OpenPortScan(request *OpenPortScanRequest) (response *OpenPortScanResponse, err error) {
response = CreateOpenPortScanResponse()
err = client.DoAction(request, response)
return
}

// OpenPortScanWithChan invokes the yundun.OpenPortScan API asynchronously
// api document: https://help.aliyun.com/api/yundun/openportscan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OpenPortScanWithChan(request *OpenPortScanRequest) (<-chan *OpenPortScanResponse, <-chan error) {
responseChan := make(chan *OpenPortScanResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.OpenPortScan(request)
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

// OpenPortScanWithCallback invokes the yundun.OpenPortScan API asynchronously
// api document: https://help.aliyun.com/api/yundun/openportscan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OpenPortScanWithCallback(request *OpenPortScanRequest, callback func(response *OpenPortScanResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *OpenPortScanResponse
var err error
defer close(result)
response, err = client.OpenPortScan(request)
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

// OpenPortScanRequest is the request struct for api OpenPortScan
type OpenPortScanRequest struct {
*requests.RpcRequest
                    InstanceId     string `position:"Query" name:"InstanceId"`
}


// OpenPortScanResponse is the response struct for api OpenPortScan
type OpenPortScanResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateOpenPortScanRequest creates a request to invoke OpenPortScan API
func CreateOpenPortScanRequest() (request *OpenPortScanRequest) {
request = &OpenPortScanRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Yundun", "2015-04-16", "OpenPortScan", "yundun", "openAPI")
return
}

// CreateOpenPortScanResponse creates a response to parse from OpenPortScan response
func CreateOpenPortScanResponse() (response *OpenPortScanResponse) {
response = &OpenPortScanResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


