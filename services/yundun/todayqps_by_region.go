
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

// TodayqpsByRegion invokes the yundun.TodayqpsByRegion API synchronously
// api document: https://help.aliyun.com/api/yundun/todayqpsbyregion.html
func (client *Client) TodayqpsByRegion(request *TodayqpsByRegionRequest) (response *TodayqpsByRegionResponse, err error) {
response = CreateTodayqpsByRegionResponse()
err = client.DoAction(request, response)
return
}

// TodayqpsByRegionWithChan invokes the yundun.TodayqpsByRegion API asynchronously
// api document: https://help.aliyun.com/api/yundun/todayqpsbyregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TodayqpsByRegionWithChan(request *TodayqpsByRegionRequest) (<-chan *TodayqpsByRegionResponse, <-chan error) {
responseChan := make(chan *TodayqpsByRegionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.TodayqpsByRegion(request)
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

// TodayqpsByRegionWithCallback invokes the yundun.TodayqpsByRegion API asynchronously
// api document: https://help.aliyun.com/api/yundun/todayqpsbyregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TodayqpsByRegionWithCallback(request *TodayqpsByRegionRequest, callback func(response *TodayqpsByRegionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *TodayqpsByRegionResponse
var err error
defer close(result)
response, err = client.TodayqpsByRegion(request)
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

// TodayqpsByRegionRequest is the request struct for api TodayqpsByRegion
type TodayqpsByRegionRequest struct {
*requests.RpcRequest
}


// TodayqpsByRegionResponse is the response struct for api TodayqpsByRegion
type TodayqpsByRegionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                    Data Data `json:"Data" xml:"Data"`
}

// CreateTodayqpsByRegionRequest creates a request to invoke TodayqpsByRegion API
func CreateTodayqpsByRegionRequest() (request *TodayqpsByRegionRequest) {
request = &TodayqpsByRegionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Yundun", "2015-02-27", "TodayqpsByRegion", "yundun", "openAPI")
return
}

// CreateTodayqpsByRegionResponse creates a response to parse from TodayqpsByRegion response
func CreateTodayqpsByRegionResponse() (response *TodayqpsByRegionResponse) {
response = &TodayqpsByRegionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


