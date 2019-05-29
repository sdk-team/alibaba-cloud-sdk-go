
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

// CheckSqlLogHistoryEnable invokes the rds.CheckSqlLogHistoryEnable API synchronously
// api document: https://help.aliyun.com/api/rds/checksqlloghistoryenable.html
func (client *Client) CheckSqlLogHistoryEnable(request *CheckSqlLogHistoryEnableRequest) (response *CheckSqlLogHistoryEnableResponse, err error) {
response = CreateCheckSqlLogHistoryEnableResponse()
err = client.DoAction(request, response)
return
}

// CheckSqlLogHistoryEnableWithChan invokes the rds.CheckSqlLogHistoryEnable API asynchronously
// api document: https://help.aliyun.com/api/rds/checksqlloghistoryenable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckSqlLogHistoryEnableWithChan(request *CheckSqlLogHistoryEnableRequest) (<-chan *CheckSqlLogHistoryEnableResponse, <-chan error) {
responseChan := make(chan *CheckSqlLogHistoryEnableResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CheckSqlLogHistoryEnable(request)
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

// CheckSqlLogHistoryEnableWithCallback invokes the rds.CheckSqlLogHistoryEnable API asynchronously
// api document: https://help.aliyun.com/api/rds/checksqlloghistoryenable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckSqlLogHistoryEnableWithCallback(request *CheckSqlLogHistoryEnableRequest, callback func(response *CheckSqlLogHistoryEnableResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CheckSqlLogHistoryEnableResponse
var err error
defer close(result)
response, err = client.CheckSqlLogHistoryEnable(request)
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

// CheckSqlLogHistoryEnableRequest is the request struct for api CheckSqlLogHistoryEnable
type CheckSqlLogHistoryEnableRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// CheckSqlLogHistoryEnableResponse is the response struct for api CheckSqlLogHistoryEnable
type CheckSqlLogHistoryEnableResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceID     int `json:"DBInstanceID" xml:"DBInstanceID"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
            HistoryEnable     string `json:"HistoryEnable" xml:"HistoryEnable"`
}

// CreateCheckSqlLogHistoryEnableRequest creates a request to invoke CheckSqlLogHistoryEnable API
func CreateCheckSqlLogHistoryEnableRequest() (request *CheckSqlLogHistoryEnableRequest) {
request = &CheckSqlLogHistoryEnableRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CheckSqlLogHistoryEnable", "rds", "openAPI")
return
}

// CreateCheckSqlLogHistoryEnableResponse creates a response to parse from CheckSqlLogHistoryEnable response
func CreateCheckSqlLogHistoryEnableResponse() (response *CheckSqlLogHistoryEnableResponse) {
response = &CheckSqlLogHistoryEnableResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


