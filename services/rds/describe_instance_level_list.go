
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

// DescribeInstanceLevelList invokes the rds.DescribeInstanceLevelList API synchronously
// api document: https://help.aliyun.com/api/rds/describeinstancelevellist.html
func (client *Client) DescribeInstanceLevelList(request *DescribeInstanceLevelListRequest) (response *DescribeInstanceLevelListResponse, err error) {
response = CreateDescribeInstanceLevelListResponse()
err = client.DoAction(request, response)
return
}

// DescribeInstanceLevelListWithChan invokes the rds.DescribeInstanceLevelList API asynchronously
// api document: https://help.aliyun.com/api/rds/describeinstancelevellist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceLevelListWithChan(request *DescribeInstanceLevelListRequest) (<-chan *DescribeInstanceLevelListResponse, <-chan error) {
responseChan := make(chan *DescribeInstanceLevelListResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeInstanceLevelList(request)
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

// DescribeInstanceLevelListWithCallback invokes the rds.DescribeInstanceLevelList API asynchronously
// api document: https://help.aliyun.com/api/rds/describeinstancelevellist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceLevelListWithCallback(request *DescribeInstanceLevelListRequest, callback func(response *DescribeInstanceLevelListResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeInstanceLevelListResponse
var err error
defer close(result)
response, err = client.DescribeInstanceLevelList(request)
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

// DescribeInstanceLevelListRequest is the request struct for api DescribeInstanceLevelList
type DescribeInstanceLevelListRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeInstanceLevelListResponse is the response struct for api DescribeInstanceLevelList
type DescribeInstanceLevelListResponse struct {
*responses.BaseResponse
            Data     string `json:"Data" xml:"Data"`
            ClassCodeList     string `json:"ClassCodeList" xml:"ClassCodeList"`
            Category     string `json:"Category" xml:"Category"`
            ClassCode     string `json:"ClassCode" xml:"ClassCode"`
            Engine     string `json:"Engine" xml:"Engine"`
            EngineVersion     string `json:"EngineVersion" xml:"EngineVersion"`
            HostType     string `json:"HostType" xml:"HostType"`
            ResultCode     string `json:"ResultCode" xml:"ResultCode"`
            Success     string `json:"Success" xml:"Success"`
}

// CreateDescribeInstanceLevelListRequest creates a request to invoke DescribeInstanceLevelList API
func CreateDescribeInstanceLevelListRequest() (request *DescribeInstanceLevelListRequest) {
request = &DescribeInstanceLevelListRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceLevelList", "", "")
return
}

// CreateDescribeInstanceLevelListResponse creates a response to parse from DescribeInstanceLevelList response
func CreateDescribeInstanceLevelListResponse() (response *DescribeInstanceLevelListResponse) {
response = &DescribeInstanceLevelListResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


