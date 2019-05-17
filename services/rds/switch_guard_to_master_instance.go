
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

// SwitchGuardToMasterInstance invokes the rds.SwitchGuardToMasterInstance API synchronously
// api document: https://help.aliyun.com/api/rds/switchguardtomasterinstance.html
func (client *Client) SwitchGuardToMasterInstance(request *SwitchGuardToMasterInstanceRequest) (response *SwitchGuardToMasterInstanceResponse, err error) {
response = CreateSwitchGuardToMasterInstanceResponse()
err = client.DoAction(request, response)
return
}

// SwitchGuardToMasterInstanceWithChan invokes the rds.SwitchGuardToMasterInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/switchguardtomasterinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchGuardToMasterInstanceWithChan(request *SwitchGuardToMasterInstanceRequest) (<-chan *SwitchGuardToMasterInstanceResponse, <-chan error) {
responseChan := make(chan *SwitchGuardToMasterInstanceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.SwitchGuardToMasterInstance(request)
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

// SwitchGuardToMasterInstanceWithCallback invokes the rds.SwitchGuardToMasterInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/switchguardtomasterinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchGuardToMasterInstanceWithCallback(request *SwitchGuardToMasterInstanceRequest, callback func(response *SwitchGuardToMasterInstanceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *SwitchGuardToMasterInstanceResponse
var err error
defer close(result)
response, err = client.SwitchGuardToMasterInstance(request)
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

// SwitchGuardToMasterInstanceRequest is the request struct for api SwitchGuardToMasterInstance
type SwitchGuardToMasterInstanceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// SwitchGuardToMasterInstanceResponse is the response struct for api SwitchGuardToMasterInstance
type SwitchGuardToMasterInstanceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
}

// CreateSwitchGuardToMasterInstanceRequest creates a request to invoke SwitchGuardToMasterInstance API
func CreateSwitchGuardToMasterInstanceRequest() (request *SwitchGuardToMasterInstanceRequest) {
request = &SwitchGuardToMasterInstanceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "SwitchGuardToMasterInstance", "", "")
return
}

// CreateSwitchGuardToMasterInstanceResponse creates a response to parse from SwitchGuardToMasterInstance response
func CreateSwitchGuardToMasterInstanceResponse() (response *SwitchGuardToMasterInstanceResponse) {
response = &SwitchGuardToMasterInstanceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


