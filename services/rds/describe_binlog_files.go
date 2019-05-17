
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

// DescribeBinlogFiles invokes the rds.DescribeBinlogFiles API synchronously
// api document: https://help.aliyun.com/api/rds/describebinlogfiles.html
func (client *Client) DescribeBinlogFiles(request *DescribeBinlogFilesRequest) (response *DescribeBinlogFilesResponse, err error) {
response = CreateDescribeBinlogFilesResponse()
err = client.DoAction(request, response)
return
}

// DescribeBinlogFilesWithChan invokes the rds.DescribeBinlogFiles API asynchronously
// api document: https://help.aliyun.com/api/rds/describebinlogfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBinlogFilesWithChan(request *DescribeBinlogFilesRequest) (<-chan *DescribeBinlogFilesResponse, <-chan error) {
responseChan := make(chan *DescribeBinlogFilesResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeBinlogFiles(request)
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

// DescribeBinlogFilesWithCallback invokes the rds.DescribeBinlogFiles API asynchronously
// api document: https://help.aliyun.com/api/rds/describebinlogfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBinlogFilesWithCallback(request *DescribeBinlogFilesRequest, callback func(response *DescribeBinlogFilesResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeBinlogFilesResponse
var err error
defer close(result)
response, err = client.DescribeBinlogFiles(request)
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

// DescribeBinlogFilesRequest is the request struct for api DescribeBinlogFiles
type DescribeBinlogFilesRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
                    PageSize     requests.Integer `position:"Query" name:"PageSize"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    EndTime     string `position:"Query" name:"EndTime"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeBinlogFilesResponse is the response struct for api DescribeBinlogFiles
type DescribeBinlogFilesResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalRecordCount     int `json:"TotalRecordCount" xml:"TotalRecordCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageRecordCount     int `json:"PageRecordCount" xml:"PageRecordCount"`
            TotalFileSize     int `json:"TotalFileSize" xml:"TotalFileSize"`
                    Items ItemsInDescribeBinlogFiles `json:"Items" xml:"Items"`
}

// CreateDescribeBinlogFilesRequest creates a request to invoke DescribeBinlogFiles API
func CreateDescribeBinlogFilesRequest() (request *DescribeBinlogFilesRequest) {
request = &DescribeBinlogFilesRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeBinlogFiles", "", "")
return
}

// CreateDescribeBinlogFilesResponse creates a response to parse from DescribeBinlogFiles response
func CreateDescribeBinlogFilesResponse() (response *DescribeBinlogFilesResponse) {
response = &DescribeBinlogFilesResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


