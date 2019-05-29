
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

// CheckInstanceExsit invokes the rds.CheckInstanceExsit API synchronously
// api document: https://help.aliyun.com/api/rds/checkinstanceexsit.html
func (client *Client) CheckInstanceExsit(request *CheckInstanceExsitRequest) (response *CheckInstanceExsitResponse, err error) {
response = CreateCheckInstanceExsitResponse()
err = client.DoAction(request, response)
return
}

// CheckInstanceExsitWithChan invokes the rds.CheckInstanceExsit API asynchronously
// api document: https://help.aliyun.com/api/rds/checkinstanceexsit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckInstanceExsitWithChan(request *CheckInstanceExsitRequest) (<-chan *CheckInstanceExsitResponse, <-chan error) {
responseChan := make(chan *CheckInstanceExsitResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CheckInstanceExsit(request)
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

// CheckInstanceExsitWithCallback invokes the rds.CheckInstanceExsit API asynchronously
// api document: https://help.aliyun.com/api/rds/checkinstanceexsit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckInstanceExsitWithCallback(request *CheckInstanceExsitRequest, callback func(response *CheckInstanceExsitResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CheckInstanceExsitResponse
var err error
defer close(result)
response, err = client.CheckInstanceExsit(request)
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

// CheckInstanceExsitRequest is the request struct for api CheckInstanceExsit
type CheckInstanceExsitRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// CheckInstanceExsitResponse is the response struct for api CheckInstanceExsit
type CheckInstanceExsitResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            IsExistInstance     bool `json:"IsExistInstance" xml:"IsExistInstance"`
}

// CreateCheckInstanceExsitRequest creates a request to invoke CheckInstanceExsit API
func CreateCheckInstanceExsitRequest() (request *CheckInstanceExsitRequest) {
request = &CheckInstanceExsitRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CheckInstanceExsit", "rds", "openAPI")
return
}

// CreateCheckInstanceExsitResponse creates a response to parse from CheckInstanceExsit response
func CreateCheckInstanceExsitResponse() (response *CheckInstanceExsitResponse) {
response = &CheckInstanceExsitResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


