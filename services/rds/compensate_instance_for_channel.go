
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

// CompensateInstanceForChannel invokes the rds.CompensateInstanceForChannel API synchronously
// api document: https://help.aliyun.com/api/rds/compensateinstanceforchannel.html
func (client *Client) CompensateInstanceForChannel(request *CompensateInstanceForChannelRequest) (response *CompensateInstanceForChannelResponse, err error) {
response = CreateCompensateInstanceForChannelResponse()
err = client.DoAction(request, response)
return
}

// CompensateInstanceForChannelWithChan invokes the rds.CompensateInstanceForChannel API asynchronously
// api document: https://help.aliyun.com/api/rds/compensateinstanceforchannel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CompensateInstanceForChannelWithChan(request *CompensateInstanceForChannelRequest) (<-chan *CompensateInstanceForChannelResponse, <-chan error) {
responseChan := make(chan *CompensateInstanceForChannelResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CompensateInstanceForChannel(request)
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

// CompensateInstanceForChannelWithCallback invokes the rds.CompensateInstanceForChannel API asynchronously
// api document: https://help.aliyun.com/api/rds/compensateinstanceforchannel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CompensateInstanceForChannelWithCallback(request *CompensateInstanceForChannelRequest, callback func(response *CompensateInstanceForChannelResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CompensateInstanceForChannelResponse
var err error
defer close(result)
response, err = client.CompensateInstanceForChannel(request)
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

// CompensateInstanceForChannelRequest is the request struct for api CompensateInstanceForChannel
type CompensateInstanceForChannelRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    SubDomain     string `position:"Query" name:"SubDomain"`
}


// CompensateInstanceForChannelResponse is the response struct for api CompensateInstanceForChannel
type CompensateInstanceForChannelResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateCompensateInstanceForChannelRequest creates a request to invoke CompensateInstanceForChannel API
func CreateCompensateInstanceForChannelRequest() (request *CompensateInstanceForChannelRequest) {
request = &CompensateInstanceForChannelRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CompensateInstanceForChannel", "rds", "openAPI")
return
}

// CreateCompensateInstanceForChannelResponse creates a response to parse from CompensateInstanceForChannel response
func CreateCompensateInstanceForChannelResponse() (response *CompensateInstanceForChannelResponse) {
response = &CompensateInstanceForChannelResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


