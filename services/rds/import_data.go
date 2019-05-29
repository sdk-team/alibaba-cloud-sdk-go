
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

// ImportData invokes the rds.ImportData API synchronously
// api document: https://help.aliyun.com/api/rds/importdata.html
func (client *Client) ImportData(request *ImportDataRequest) (response *ImportDataResponse, err error) {
response = CreateImportDataResponse()
err = client.DoAction(request, response)
return
}

// ImportDataWithChan invokes the rds.ImportData API asynchronously
// api document: https://help.aliyun.com/api/rds/importdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportDataWithChan(request *ImportDataRequest) (<-chan *ImportDataResponse, <-chan error) {
responseChan := make(chan *ImportDataResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ImportData(request)
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

// ImportDataWithCallback invokes the rds.ImportData API asynchronously
// api document: https://help.aliyun.com/api/rds/importdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportDataWithCallback(request *ImportDataRequest, callback func(response *ImportDataResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ImportDataResponse
var err error
defer close(result)
response, err = client.ImportData(request)
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

// ImportDataRequest is the request struct for api ImportData
type ImportDataRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    FileName     string `position:"Query" name:"FileName"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// ImportDataResponse is the response struct for api ImportData
type ImportDataResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateImportDataRequest creates a request to invoke ImportData API
func CreateImportDataRequest() (request *ImportDataRequest) {
request = &ImportDataRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2013-05-28", "ImportData", "rds", "openAPI")
return
}

// CreateImportDataResponse creates a response to parse from ImportData response
func CreateImportDataResponse() (response *ImportDataResponse) {
response = &ImportDataResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


