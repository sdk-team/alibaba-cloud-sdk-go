
package rds

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

// DescribeVerificationList invokes the rds.DescribeVerificationList API synchronously
// api document: https://help.aliyun.com/api/rds/describeverificationlist.html
func (client *Client) DescribeVerificationList(request *DescribeVerificationListRequest) (response *DescribeVerificationListResponse, err error) {
response = CreateDescribeVerificationListResponse()
err = client.DoAction(request, response)
return
}

// DescribeVerificationListWithChan invokes the rds.DescribeVerificationList API asynchronously
// api document: https://help.aliyun.com/api/rds/describeverificationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerificationListWithChan(request *DescribeVerificationListRequest) (<-chan *DescribeVerificationListResponse, <-chan error) {
responseChan := make(chan *DescribeVerificationListResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeVerificationList(request)
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

// DescribeVerificationListWithCallback invokes the rds.DescribeVerificationList API asynchronously
// api document: https://help.aliyun.com/api/rds/describeverificationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerificationListWithCallback(request *DescribeVerificationListRequest, callback func(response *DescribeVerificationListResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeVerificationListResponse
var err error
defer close(result)
response, err = client.DescribeVerificationList(request)
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

// DescribeVerificationListRequest is the request struct for api DescribeVerificationList
type DescribeVerificationListRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ReplicaId     string `position:"Query" name:"ReplicaId"`
                    PageSize     requests.Integer `position:"Query" name:"PageSize"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    EndTime     string `position:"Query" name:"EndTime"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeVerificationListResponse is the response struct for api DescribeVerificationList
type DescribeVerificationListResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            ReplicaId     string `json:"ReplicaId" xml:"ReplicaId"`
            PagNumber     int `json:"PagNumber" xml:"PagNumber"`
            PageRecordCount     int `json:"PageRecordCount" xml:"PageRecordCount"`
            TotalRecordCount     int `json:"TotalRecordCount" xml:"TotalRecordCount"`
                    Items []ItemsItem  `json:"Items" xml:"Items"`
}

// CreateDescribeVerificationListRequest creates a request to invoke DescribeVerificationList API
func CreateDescribeVerificationListRequest() (request *DescribeVerificationListRequest) {
request = &DescribeVerificationListRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeVerificationList", "rds", "openAPI")
return
}

// CreateDescribeVerificationListResponse creates a response to parse from DescribeVerificationList response
func CreateDescribeVerificationListResponse() (response *DescribeVerificationListResponse) {
response = &DescribeVerificationListResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


