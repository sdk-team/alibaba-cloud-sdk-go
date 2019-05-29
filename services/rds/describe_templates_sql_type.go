
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

// DescribeTemplatesSqlType invokes the rds.DescribeTemplatesSqlType API synchronously
// api document: https://help.aliyun.com/api/rds/describetemplatessqltype.html
func (client *Client) DescribeTemplatesSqlType(request *DescribeTemplatesSqlTypeRequest) (response *DescribeTemplatesSqlTypeResponse, err error) {
response = CreateDescribeTemplatesSqlTypeResponse()
err = client.DoAction(request, response)
return
}

// DescribeTemplatesSqlTypeWithChan invokes the rds.DescribeTemplatesSqlType API asynchronously
// api document: https://help.aliyun.com/api/rds/describetemplatessqltype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTemplatesSqlTypeWithChan(request *DescribeTemplatesSqlTypeRequest) (<-chan *DescribeTemplatesSqlTypeResponse, <-chan error) {
responseChan := make(chan *DescribeTemplatesSqlTypeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeTemplatesSqlType(request)
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

// DescribeTemplatesSqlTypeWithCallback invokes the rds.DescribeTemplatesSqlType API asynchronously
// api document: https://help.aliyun.com/api/rds/describetemplatessqltype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTemplatesSqlTypeWithCallback(request *DescribeTemplatesSqlTypeRequest, callback func(response *DescribeTemplatesSqlTypeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeTemplatesSqlTypeResponse
var err error
defer close(result)
response, err = client.DescribeTemplatesSqlType(request)
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

// DescribeTemplatesSqlTypeRequest is the request struct for api DescribeTemplatesSqlType
type DescribeTemplatesSqlTypeRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    MinScanRows     string `position:"Query" name:"MinScanRows"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    SortKey     string `position:"Query" name:"SortKey"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    PagingId     string `position:"Query" name:"PagingId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    MinConsume     string `position:"Query" name:"MinConsume"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    EndTime     string `position:"Query" name:"EndTime"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    MaxConsume     string `position:"Query" name:"MaxConsume"`
                    SortMethod     string `position:"Query" name:"SortMethod"`
                    MaxScanRows     string `position:"Query" name:"MaxScanRows"`
}


// DescribeTemplatesSqlTypeResponse is the response struct for api DescribeTemplatesSqlType
type DescribeTemplatesSqlTypeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceID     int `json:"DBInstanceID" xml:"DBInstanceID"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            TotalRecords     int `json:"TotalRecords" xml:"TotalRecords"`
            PagingID     string `json:"PagingID" xml:"PagingID"`
            ItemsNumbers     int `json:"ItemsNumbers" xml:"ItemsNumbers"`
                    Items ItemsInDescribeTemplatesSqlType `json:"Items" xml:"Items"`
}

// CreateDescribeTemplatesSqlTypeRequest creates a request to invoke DescribeTemplatesSqlType API
func CreateDescribeTemplatesSqlTypeRequest() (request *DescribeTemplatesSqlTypeRequest) {
request = &DescribeTemplatesSqlTypeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeTemplatesSqlType", "rds", "openAPI")
return
}

// CreateDescribeTemplatesSqlTypeResponse creates a response to parse from DescribeTemplatesSqlType response
func CreateDescribeTemplatesSqlTypeResponse() (response *DescribeTemplatesSqlTypeResponse) {
response = &DescribeTemplatesSqlTypeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


