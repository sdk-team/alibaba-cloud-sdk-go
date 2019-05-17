
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

// QueryPriceMultiForDefer invokes the rds.QueryPriceMultiForDefer API synchronously
// api document: https://help.aliyun.com/api/rds/querypricemultifordefer.html
func (client *Client) QueryPriceMultiForDefer(request *QueryPriceMultiForDeferRequest) (response *QueryPriceMultiForDeferResponse, err error) {
response = CreateQueryPriceMultiForDeferResponse()
err = client.DoAction(request, response)
return
}

// QueryPriceMultiForDeferWithChan invokes the rds.QueryPriceMultiForDefer API asynchronously
// api document: https://help.aliyun.com/api/rds/querypricemultifordefer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceMultiForDeferWithChan(request *QueryPriceMultiForDeferRequest) (<-chan *QueryPriceMultiForDeferResponse, <-chan error) {
responseChan := make(chan *QueryPriceMultiForDeferResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryPriceMultiForDefer(request)
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

// QueryPriceMultiForDeferWithCallback invokes the rds.QueryPriceMultiForDefer API asynchronously
// api document: https://help.aliyun.com/api/rds/querypricemultifordefer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceMultiForDeferWithCallback(request *QueryPriceMultiForDeferRequest, callback func(response *QueryPriceMultiForDeferResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryPriceMultiForDeferResponse
var err error
defer close(result)
response, err = client.QueryPriceMultiForDefer(request)
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

// QueryPriceMultiForDeferRequest is the request struct for api QueryPriceMultiForDefer
type QueryPriceMultiForDeferRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    BusinessInfo     string `position:"Query" name:"BusinessInfo"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    Resource     string `position:"Query" name:"Resource"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    PromotionCode     string `position:"Query" name:"PromotionCode"`
}


// QueryPriceMultiForDeferResponse is the response struct for api QueryPriceMultiForDefer
type QueryPriceMultiForDeferResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Order Order  `json:"Order" xml:"Order"`
                    Rules RulesInQueryPriceMultiForDefer `json:"Rules" xml:"Rules"`
}

// CreateQueryPriceMultiForDeferRequest creates a request to invoke QueryPriceMultiForDefer API
func CreateQueryPriceMultiForDeferRequest() (request *QueryPriceMultiForDeferRequest) {
request = &QueryPriceMultiForDeferRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "QueryPriceMultiForDefer", "", "")
return
}

// CreateQueryPriceMultiForDeferResponse creates a response to parse from QueryPriceMultiForDefer response
func CreateQueryPriceMultiForDeferResponse() (response *QueryPriceMultiForDeferResponse) {
response = &QueryPriceMultiForDeferResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


