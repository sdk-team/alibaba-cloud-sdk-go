
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

// DescribeOptimizeAdviceByDBA invokes the rds.DescribeOptimizeAdviceByDBA API synchronously
// api document: https://help.aliyun.com/api/rds/describeoptimizeadvicebydba.html
func (client *Client) DescribeOptimizeAdviceByDBA(request *DescribeOptimizeAdviceByDBARequest) (response *DescribeOptimizeAdviceByDBAResponse, err error) {
response = CreateDescribeOptimizeAdviceByDBAResponse()
err = client.DoAction(request, response)
return
}

// DescribeOptimizeAdviceByDBAWithChan invokes the rds.DescribeOptimizeAdviceByDBA API asynchronously
// api document: https://help.aliyun.com/api/rds/describeoptimizeadvicebydba.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOptimizeAdviceByDBAWithChan(request *DescribeOptimizeAdviceByDBARequest) (<-chan *DescribeOptimizeAdviceByDBAResponse, <-chan error) {
responseChan := make(chan *DescribeOptimizeAdviceByDBAResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeOptimizeAdviceByDBA(request)
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

// DescribeOptimizeAdviceByDBAWithCallback invokes the rds.DescribeOptimizeAdviceByDBA API asynchronously
// api document: https://help.aliyun.com/api/rds/describeoptimizeadvicebydba.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOptimizeAdviceByDBAWithCallback(request *DescribeOptimizeAdviceByDBARequest, callback func(response *DescribeOptimizeAdviceByDBAResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeOptimizeAdviceByDBAResponse
var err error
defer close(result)
response, err = client.DescribeOptimizeAdviceByDBA(request)
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

// DescribeOptimizeAdviceByDBARequest is the request struct for api DescribeOptimizeAdviceByDBA
type DescribeOptimizeAdviceByDBARequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
                    PageSize     requests.Integer `position:"Query" name:"PageSize"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeOptimizeAdviceByDBAResponse is the response struct for api DescribeOptimizeAdviceByDBA
type DescribeOptimizeAdviceByDBAResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalRecordsCount     int `json:"TotalRecordsCount" xml:"TotalRecordsCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageRecordCount     int `json:"PageRecordCount" xml:"PageRecordCount"`
                    Items ItemsInDescribeOptimizeAdviceByDBA `json:"Items" xml:"Items"`
}

// CreateDescribeOptimizeAdviceByDBARequest creates a request to invoke DescribeOptimizeAdviceByDBA API
func CreateDescribeOptimizeAdviceByDBARequest() (request *DescribeOptimizeAdviceByDBARequest) {
request = &DescribeOptimizeAdviceByDBARequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceByDBA", "rds", "openAPI")
return
}

// CreateDescribeOptimizeAdviceByDBAResponse creates a response to parse from DescribeOptimizeAdviceByDBA response
func CreateDescribeOptimizeAdviceByDBAResponse() (response *DescribeOptimizeAdviceByDBAResponse) {
response = &DescribeOptimizeAdviceByDBAResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


