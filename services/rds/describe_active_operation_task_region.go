
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

// DescribeActiveOperationTaskRegion invokes the rds.DescribeActiveOperationTaskRegion API synchronously
// api document: https://help.aliyun.com/api/rds/describeactiveoperationtaskregion.html
func (client *Client) DescribeActiveOperationTaskRegion(request *DescribeActiveOperationTaskRegionRequest) (response *DescribeActiveOperationTaskRegionResponse, err error) {
response = CreateDescribeActiveOperationTaskRegionResponse()
err = client.DoAction(request, response)
return
}

// DescribeActiveOperationTaskRegionWithChan invokes the rds.DescribeActiveOperationTaskRegion API asynchronously
// api document: https://help.aliyun.com/api/rds/describeactiveoperationtaskregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveOperationTaskRegionWithChan(request *DescribeActiveOperationTaskRegionRequest) (<-chan *DescribeActiveOperationTaskRegionResponse, <-chan error) {
responseChan := make(chan *DescribeActiveOperationTaskRegionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeActiveOperationTaskRegion(request)
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

// DescribeActiveOperationTaskRegionWithCallback invokes the rds.DescribeActiveOperationTaskRegion API asynchronously
// api document: https://help.aliyun.com/api/rds/describeactiveoperationtaskregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveOperationTaskRegionWithCallback(request *DescribeActiveOperationTaskRegionRequest, callback func(response *DescribeActiveOperationTaskRegionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeActiveOperationTaskRegionResponse
var err error
defer close(result)
response, err = client.DescribeActiveOperationTaskRegion(request)
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

// DescribeActiveOperationTaskRegionRequest is the request struct for api DescribeActiveOperationTaskRegion
type DescribeActiveOperationTaskRegionRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    IsHistory     requests.Integer `position:"Query" name:"IsHistory"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    TaskType     string `position:"Query" name:"TaskType"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeActiveOperationTaskRegionResponse is the response struct for api DescribeActiveOperationTaskRegion
type DescribeActiveOperationTaskRegionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                    RegionList []Items  `json:"RegionList" xml:"RegionList"`
}

// CreateDescribeActiveOperationTaskRegionRequest creates a request to invoke DescribeActiveOperationTaskRegion API
func CreateDescribeActiveOperationTaskRegionRequest() (request *DescribeActiveOperationTaskRegionRequest) {
request = &DescribeActiveOperationTaskRegionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeActiveOperationTaskRegion", "rds", "openAPI")
return
}

// CreateDescribeActiveOperationTaskRegionResponse creates a response to parse from DescribeActiveOperationTaskRegion response
func CreateDescribeActiveOperationTaskRegionResponse() (response *DescribeActiveOperationTaskRegionResponse) {
response = &DescribeActiveOperationTaskRegionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


