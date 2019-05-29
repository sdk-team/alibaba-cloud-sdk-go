
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

// PurgeDBInstanceLog invokes the rds.PurgeDBInstanceLog API synchronously
// api document: https://help.aliyun.com/api/rds/purgedbinstancelog.html
func (client *Client) PurgeDBInstanceLog(request *PurgeDBInstanceLogRequest) (response *PurgeDBInstanceLogResponse, err error) {
response = CreatePurgeDBInstanceLogResponse()
err = client.DoAction(request, response)
return
}

// PurgeDBInstanceLogWithChan invokes the rds.PurgeDBInstanceLog API asynchronously
// api document: https://help.aliyun.com/api/rds/purgedbinstancelog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PurgeDBInstanceLogWithChan(request *PurgeDBInstanceLogRequest) (<-chan *PurgeDBInstanceLogResponse, <-chan error) {
responseChan := make(chan *PurgeDBInstanceLogResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PurgeDBInstanceLog(request)
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

// PurgeDBInstanceLogWithCallback invokes the rds.PurgeDBInstanceLog API asynchronously
// api document: https://help.aliyun.com/api/rds/purgedbinstancelog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PurgeDBInstanceLogWithCallback(request *PurgeDBInstanceLogRequest, callback func(response *PurgeDBInstanceLogResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PurgeDBInstanceLogResponse
var err error
defer close(result)
response, err = client.PurgeDBInstanceLog(request)
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

// PurgeDBInstanceLogRequest is the request struct for api PurgeDBInstanceLog
type PurgeDBInstanceLogRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// PurgeDBInstanceLogResponse is the response struct for api PurgeDBInstanceLog
type PurgeDBInstanceLogResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePurgeDBInstanceLogRequest creates a request to invoke PurgeDBInstanceLog API
func CreatePurgeDBInstanceLogRequest() (request *PurgeDBInstanceLogRequest) {
request = &PurgeDBInstanceLogRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "PurgeDBInstanceLog", "rds", "openAPI")
return
}

// CreatePurgeDBInstanceLogResponse creates a response to parse from PurgeDBInstanceLog response
func CreatePurgeDBInstanceLogResponse() (response *PurgeDBInstanceLogResponse) {
response = &PurgeDBInstanceLogResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


