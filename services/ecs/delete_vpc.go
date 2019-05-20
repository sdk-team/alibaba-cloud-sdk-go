
package ecs

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

// DeleteVpc invokes the ecs.DeleteVpc API synchronously
// api document: https://help.aliyun.com/api/ecs/deletevpc.html
func (client *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
response = CreateDeleteVpcResponse()
err = client.DoAction(request, response)
return
}

// DeleteVpcWithChan invokes the ecs.DeleteVpc API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletevpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVpcWithChan(request *DeleteVpcRequest) (<-chan *DeleteVpcResponse, <-chan error) {
responseChan := make(chan *DeleteVpcResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteVpc(request)
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

// DeleteVpcWithCallback invokes the ecs.DeleteVpc API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletevpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVpcWithCallback(request *DeleteVpcRequest, callback func(response *DeleteVpcResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteVpcResponse
var err error
defer close(result)
response, err = client.DeleteVpc(request)
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

// DeleteVpcRequest is the request struct for api DeleteVpc
type DeleteVpcRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    VpcId     string `position:"Query" name:"VpcId"`
}


// DeleteVpcResponse is the response struct for api DeleteVpc
type DeleteVpcResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVpcRequest creates a request to invoke DeleteVpc API
func CreateDeleteVpcRequest() (request *DeleteVpcRequest) {
request = &DeleteVpcRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteVpc", "ecs", "openAPI")
return
}

// CreateDeleteVpcResponse creates a response to parse from DeleteVpc response
func CreateDeleteVpcResponse() (response *DeleteVpcResponse) {
response = &DeleteVpcResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


