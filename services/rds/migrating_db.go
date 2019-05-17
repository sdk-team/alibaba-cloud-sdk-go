
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

// MigratingDB invokes the rds.MigratingDB API synchronously
// api document: https://help.aliyun.com/api/rds/migratingdb.html
func (client *Client) MigratingDB(request *MigratingDBRequest) (response *MigratingDBResponse, err error) {
response = CreateMigratingDBResponse()
err = client.DoAction(request, response)
return
}

// MigratingDBWithChan invokes the rds.MigratingDB API asynchronously
// api document: https://help.aliyun.com/api/rds/migratingdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MigratingDBWithChan(request *MigratingDBRequest) (<-chan *MigratingDBResponse, <-chan error) {
responseChan := make(chan *MigratingDBResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.MigratingDB(request)
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

// MigratingDBWithCallback invokes the rds.MigratingDB API asynchronously
// api document: https://help.aliyun.com/api/rds/migratingdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MigratingDBWithCallback(request *MigratingDBRequest, callback func(response *MigratingDBResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *MigratingDBResponse
var err error
defer close(result)
response, err = client.MigratingDB(request)
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

// MigratingDBRequest is the request struct for api MigratingDB
type MigratingDBRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    TargetDBInstanceId     string `position:"Query" name:"TargetDBInstanceId"`
                    DBInfo     string `position:"Query" name:"DBInfo"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// MigratingDBResponse is the response struct for api MigratingDB
type MigratingDBResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            MigrationId     int `json:"MigrationId" xml:"MigrationId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
}

// CreateMigratingDBRequest creates a request to invoke MigratingDB API
func CreateMigratingDBRequest() (request *MigratingDBRequest) {
request = &MigratingDBRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2013-05-28", "MigratingDB", "", "")
return
}

// CreateMigratingDBResponse creates a response to parse from MigratingDB response
func CreateMigratingDBResponse() (response *MigratingDBResponse) {
response = &MigratingDBResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


