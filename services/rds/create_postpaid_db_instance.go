
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

// CreatePostpaidDBInstance invokes the rds.CreatePostpaidDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/createpostpaiddbinstance.html
func (client *Client) CreatePostpaidDBInstance(request *CreatePostpaidDBInstanceRequest) (response *CreatePostpaidDBInstanceResponse, err error) {
response = CreateCreatePostpaidDBInstanceResponse()
err = client.DoAction(request, response)
return
}

// CreatePostpaidDBInstanceWithChan invokes the rds.CreatePostpaidDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createpostpaiddbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePostpaidDBInstanceWithChan(request *CreatePostpaidDBInstanceRequest) (<-chan *CreatePostpaidDBInstanceResponse, <-chan error) {
responseChan := make(chan *CreatePostpaidDBInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreatePostpaidDBInstance(request)
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

// CreatePostpaidDBInstanceWithCallback invokes the rds.CreatePostpaidDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createpostpaiddbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePostpaidDBInstanceWithCallback(request *CreatePostpaidDBInstanceRequest, callback func(response *CreatePostpaidDBInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreatePostpaidDBInstanceResponse
var err error
defer close(result)
response, err = client.CreatePostpaidDBInstance(request)
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

// CreatePostpaidDBInstanceRequest is the request struct for api CreatePostpaidDBInstance
type CreatePostpaidDBInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    Engine     string `position:"Query" name:"Engine"`
                    DBInstanceDescription     string `position:"Query" name:"DBInstanceDescription"`
                    DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    SecurityIPList     string `position:"Query" name:"SecurityIPList"`
}


// CreatePostpaidDBInstanceResponse is the response struct for api CreatePostpaidDBInstance
type CreatePostpaidDBInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     string `json:"OrderId" xml:"OrderId"`
            ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
            Port     string `json:"Port" xml:"Port"`
}

// CreateCreatePostpaidDBInstanceRequest creates a request to invoke CreatePostpaidDBInstance API
func CreateCreatePostpaidDBInstanceRequest() (request *CreatePostpaidDBInstanceRequest) {
request = &CreatePostpaidDBInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreatePostpaidDBInstance", "", "")
return
}

// CreateCreatePostpaidDBInstanceResponse creates a response to parse from CreatePostpaidDBInstance response
func CreateCreatePostpaidDBInstanceResponse() (response *CreatePostpaidDBInstanceResponse) {
response = &CreatePostpaidDBInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


