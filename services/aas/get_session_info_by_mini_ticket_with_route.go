
package aas

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

// GetSessionInfoByMiniTicketWithRoute invokes the aas.GetSessionInfoByMiniTicketWithRoute API synchronously
// api document: https://help.aliyun.com/api/aas/getsessioninfobyminiticketwithroute.html
func (client *Client) GetSessionInfoByMiniTicketWithRoute(request *GetSessionInfoByMiniTicketWithRouteRequest) (response *GetSessionInfoByMiniTicketWithRouteResponse, err error) {
response = CreateGetSessionInfoByMiniTicketWithRouteResponse()
err = client.DoAction(request, response)
return
}

// GetSessionInfoByMiniTicketWithRouteWithChan invokes the aas.GetSessionInfoByMiniTicketWithRoute API asynchronously
// api document: https://help.aliyun.com/api/aas/getsessioninfobyminiticketwithroute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSessionInfoByMiniTicketWithRouteWithChan(request *GetSessionInfoByMiniTicketWithRouteRequest) (<-chan *GetSessionInfoByMiniTicketWithRouteResponse, <-chan error) {
responseChan := make(chan *GetSessionInfoByMiniTicketWithRouteResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetSessionInfoByMiniTicketWithRoute(request)
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

// GetSessionInfoByMiniTicketWithRouteWithCallback invokes the aas.GetSessionInfoByMiniTicketWithRoute API asynchronously
// api document: https://help.aliyun.com/api/aas/getsessioninfobyminiticketwithroute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSessionInfoByMiniTicketWithRouteWithCallback(request *GetSessionInfoByMiniTicketWithRouteRequest, callback func(response *GetSessionInfoByMiniTicketWithRouteResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetSessionInfoByMiniTicketWithRouteResponse
var err error
defer close(result)
response, err = client.GetSessionInfoByMiniTicketWithRoute(request)
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

// GetSessionInfoByMiniTicketWithRouteRequest is the request struct for api GetSessionInfoByMiniTicketWithRoute
type GetSessionInfoByMiniTicketWithRouteRequest struct {
*requests.RpcRequest
                    Ticket     string `position:"Query" name:"Ticket"`
}


// GetSessionInfoByMiniTicketWithRouteResponse is the response struct for api GetSessionInfoByMiniTicketWithRoute
type GetSessionInfoByMiniTicketWithRouteResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            SessionInfo     string `json:"SessionInfo" xml:"SessionInfo"`
}

// CreateGetSessionInfoByMiniTicketWithRouteRequest creates a request to invoke GetSessionInfoByMiniTicketWithRoute API
func CreateGetSessionInfoByMiniTicketWithRouteRequest() (request *GetSessionInfoByMiniTicketWithRouteRequest) {
request = &GetSessionInfoByMiniTicketWithRouteRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Aas", "2015-07-01", "GetSessionInfoByMiniTicketWithRoute", "", "")
return
}

// CreateGetSessionInfoByMiniTicketWithRouteResponse creates a response to parse from GetSessionInfoByMiniTicketWithRoute response
func CreateGetSessionInfoByMiniTicketWithRouteResponse() (response *GetSessionInfoByMiniTicketWithRouteResponse) {
response = &GetSessionInfoByMiniTicketWithRouteResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


