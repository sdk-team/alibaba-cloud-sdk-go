
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

// MigrateToOtherZone invokes the rds.MigrateToOtherZone API synchronously
// api document: https://help.aliyun.com/api/rds/migratetootherzone.html
func (client *Client) MigrateToOtherZone(request *MigrateToOtherZoneRequest) (response *MigrateToOtherZoneResponse, err error) {
response = CreateMigrateToOtherZoneResponse()
err = client.DoAction(request, response)
return
}

// MigrateToOtherZoneWithChan invokes the rds.MigrateToOtherZone API asynchronously
// api document: https://help.aliyun.com/api/rds/migratetootherzone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MigrateToOtherZoneWithChan(request *MigrateToOtherZoneRequest) (<-chan *MigrateToOtherZoneResponse, <-chan error) {
responseChan := make(chan *MigrateToOtherZoneResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.MigrateToOtherZone(request)
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

// MigrateToOtherZoneWithCallback invokes the rds.MigrateToOtherZone API asynchronously
// api document: https://help.aliyun.com/api/rds/migratetootherzone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MigrateToOtherZoneWithCallback(request *MigrateToOtherZoneRequest, callback func(response *MigrateToOtherZoneResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *MigrateToOtherZoneResponse
var err error
defer close(result)
response, err = client.MigrateToOtherZone(request)
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

// MigrateToOtherZoneRequest is the request struct for api MigrateToOtherZone
type MigrateToOtherZoneRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ZoneIdSlave1     string `position:"Query" name:"ZoneIdSlave1"`
                    ZoneIdSlave2     string `position:"Query" name:"ZoneIdSlave2"`
                    EffectiveTime     string `position:"Query" name:"EffectiveTime"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    VSwitchId     string `position:"Query" name:"VSwitchId"`
                    VPCId     string `position:"Query" name:"VPCId"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    Category     string `position:"Query" name:"Category"`
}


// MigrateToOtherZoneResponse is the response struct for api MigrateToOtherZone
type MigrateToOtherZoneResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateMigrateToOtherZoneRequest creates a request to invoke MigrateToOtherZone API
func CreateMigrateToOtherZoneRequest() (request *MigrateToOtherZoneRequest) {
request = &MigrateToOtherZoneRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "MigrateToOtherZone", "", "")
return
}

// CreateMigrateToOtherZoneResponse creates a response to parse from MigrateToOtherZone response
func CreateMigrateToOtherZoneResponse() (response *MigrateToOtherZoneResponse) {
response = &MigrateToOtherZoneResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


