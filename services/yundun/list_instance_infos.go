
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

// ListInstanceInfos invokes the yundun.ListInstanceInfos API synchronously
// api document: https://help.aliyun.com/api/yundun/listinstanceinfos.html
func (client *Client) ListInstanceInfos(request *ListInstanceInfosRequest) (response *ListInstanceInfosResponse, err error) {
response = CreateListInstanceInfosResponse()
err = client.DoAction(request, response)
return
}

// ListInstanceInfosWithChan invokes the yundun.ListInstanceInfos API asynchronously
// api document: https://help.aliyun.com/api/yundun/listinstanceinfos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstanceInfosWithChan(request *ListInstanceInfosRequest) (<-chan *ListInstanceInfosResponse, <-chan error) {
responseChan := make(chan *ListInstanceInfosResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ListInstanceInfos(request)
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

// ListInstanceInfosWithCallback invokes the yundun.ListInstanceInfos API asynchronously
// api document: https://help.aliyun.com/api/yundun/listinstanceinfos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstanceInfosWithCallback(request *ListInstanceInfosRequest, callback func(response *ListInstanceInfosResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ListInstanceInfosResponse
var err error
defer close(result)
response, err = client.ListInstanceInfos(request)
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

// ListInstanceInfosRequest is the request struct for api ListInstanceInfos
type ListInstanceInfosRequest struct {
*requests.RpcRequest
                    JstOwnerId     requests.Integer `position:"Query" name:"JstOwnerId"`
                    InstanceName     string `position:"Query" name:"InstanceName"`
                    InstanceIds     string `position:"Query" name:"InstanceIds"`
                    PageSize     requests.Integer `position:"Query" name:"PageSize"`
                    InstanceType     string `position:"Query" name:"InstanceType"`
                    EventType     string `position:"Query" name:"EventType"`
                    Region     string `position:"Query" name:"Region"`
                    PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
}


// ListInstanceInfosResponse is the response struct for api ListInstanceInfos
type ListInstanceInfosResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
                    InfosList InfosList `json:"InfosList" xml:"InfosList"`
}

// CreateListInstanceInfosRequest creates a request to invoke ListInstanceInfos API
func CreateListInstanceInfosRequest() (request *ListInstanceInfosRequest) {
request = &ListInstanceInfosRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Yundun", "2015-04-16", "ListInstanceInfos", "yundun", "openAPI")
return
}

// CreateListInstanceInfosResponse creates a response to parse from ListInstanceInfos response
func CreateListInstanceInfosResponse() (response *ListInstanceInfosResponse) {
response = &ListInstanceInfosResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


