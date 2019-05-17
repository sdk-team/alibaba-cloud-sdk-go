
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

// QueryPrice invokes the rds.QueryPrice API synchronously
// api document: https://help.aliyun.com/api/rds/queryprice.html
func (client *Client) QueryPrice(request *QueryPriceRequest) (response *QueryPriceResponse, err error) {
response = CreateQueryPriceResponse()
err = client.DoAction(request, response)
return
}

// QueryPriceWithChan invokes the rds.QueryPrice API asynchronously
// api document: https://help.aliyun.com/api/rds/queryprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceWithChan(request *QueryPriceRequest) (<-chan *QueryPriceResponse, <-chan error) {
responseChan := make(chan *QueryPriceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.QueryPrice(request)
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

// QueryPriceWithCallback invokes the rds.QueryPrice API asynchronously
// api document: https://help.aliyun.com/api/rds/queryprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceWithCallback(request *QueryPriceRequest, callback func(response *QueryPriceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *QueryPriceResponse
var err error
defer close(result)
response, err = client.QueryPrice(request)
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

// QueryPriceRequest is the request struct for api QueryPrice
type QueryPriceRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    SystemDBCharset     string `position:"Query" name:"SystemDBCharset"`
                    CountryCode     string `position:"Query" name:"CountryCode"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    CurrencyCode     string `position:"Query" name:"CurrencyCode"`
                    DBInstanceDescription     string `position:"Query" name:"DBInstanceDescription"`
                    BusinessInfo     string `position:"Query" name:"BusinessInfo"`
                    Resource     string `position:"Query" name:"Resource"`
                    CommodityCode     string `position:"Query" name:"CommodityCode"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    SecurityIPList     string `position:"Query" name:"SecurityIPList"`
                    VSwitchId     string `position:"Query" name:"VSwitchId"`
                    PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
                    PromotionCode     string `position:"Query" name:"PromotionCode"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    TimeType     string `position:"Query" name:"TimeType"`
                    InstanceNetworkType     string `position:"Query" name:"InstanceNetworkType"`
                    OrderType     string `position:"Query" name:"OrderType"`
                    ConnectionMode     string `position:"Query" name:"ConnectionMode"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    Engine     string `position:"Query" name:"Engine"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    DBInstanceStorageType     string `position:"Query" name:"DBInstanceStorageType"`
                    DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
                    Quantity     requests.Integer `position:"Query" name:"Quantity"`
                    AutoPay     requests.Boolean `position:"Query" name:"AutoPay"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OrderParamOut     string `position:"Query" name:"OrderParamOut"`
                    UsedTime     string `position:"Query" name:"UsedTime"`
                    InstanceUsedType     requests.Integer `position:"Query" name:"InstanceUsedType"`
                    VPCId     string `position:"Query" name:"VPCId"`
                    PayType     string `position:"Query" name:"PayType"`
}


// QueryPriceResponse is the response struct for api QueryPrice
type QueryPriceResponse struct {
*responses.BaseResponse
            OrderParams     string `json:"OrderParams" xml:"OrderParams"`
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Order Order  `json:"Order" xml:"Order"`
                    Rules RulesInQueryPrice `json:"Rules" xml:"Rules"`
}

// CreateQueryPriceRequest creates a request to invoke QueryPrice API
func CreateQueryPriceRequest() (request *QueryPriceRequest) {
request = &QueryPriceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "QueryPrice", "", "")
return
}

// CreateQueryPriceResponse creates a response to parse from QueryPrice response
func CreateQueryPriceResponse() (response *QueryPriceResponse) {
response = &QueryPriceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


