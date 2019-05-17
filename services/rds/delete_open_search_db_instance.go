
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

// DeleteOpenSearchDBInstance invokes the rds.DeleteOpenSearchDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/deleteopensearchdbinstance.html
func (client *Client) DeleteOpenSearchDBInstance(request *DeleteOpenSearchDBInstanceRequest) (response *DeleteOpenSearchDBInstanceResponse, err error) {
response = CreateDeleteOpenSearchDBInstanceResponse()
err = client.DoAction(request, response)
return
}

// DeleteOpenSearchDBInstanceWithChan invokes the rds.DeleteOpenSearchDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/deleteopensearchdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteOpenSearchDBInstanceWithChan(request *DeleteOpenSearchDBInstanceRequest) (<-chan *DeleteOpenSearchDBInstanceResponse, <-chan error) {
responseChan := make(chan *DeleteOpenSearchDBInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteOpenSearchDBInstance(request)
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

// DeleteOpenSearchDBInstanceWithCallback invokes the rds.DeleteOpenSearchDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/deleteopensearchdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteOpenSearchDBInstanceWithCallback(request *DeleteOpenSearchDBInstanceRequest, callback func(response *DeleteOpenSearchDBInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteOpenSearchDBInstanceResponse
var err error
defer close(result)
response, err = client.DeleteOpenSearchDBInstance(request)
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

// DeleteOpenSearchDBInstanceRequest is the request struct for api DeleteOpenSearchDBInstance
type DeleteOpenSearchDBInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// DeleteOpenSearchDBInstanceResponse is the response struct for api DeleteOpenSearchDBInstance
type DeleteOpenSearchDBInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TaskId     string `json:"TaskId" xml:"TaskId"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateDeleteOpenSearchDBInstanceRequest creates a request to invoke DeleteOpenSearchDBInstance API
func CreateDeleteOpenSearchDBInstanceRequest() (request *DeleteOpenSearchDBInstanceRequest) {
request = &DeleteOpenSearchDBInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DeleteOpenSearchDBInstance", "", "")
return
}

// CreateDeleteOpenSearchDBInstanceResponse creates a response to parse from DeleteOpenSearchDBInstance response
func CreateDeleteOpenSearchDBInstanceResponse() (response *DeleteOpenSearchDBInstanceResponse) {
response = &DeleteOpenSearchDBInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


