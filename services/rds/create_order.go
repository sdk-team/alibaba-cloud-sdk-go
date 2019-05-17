
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

// CreateOrder invokes the rds.CreateOrder API synchronously
// api document: https://help.aliyun.com/api/rds/createorder.html
func (client *Client) CreateOrder(request *CreateOrderRequest) (response *CreateOrderResponse, err error) {
response = CreateCreateOrderResponse()
err = client.DoAction(request, response)
return
}

// CreateOrderWithChan invokes the rds.CreateOrder API asynchronously
// api document: https://help.aliyun.com/api/rds/createorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrderWithChan(request *CreateOrderRequest) (<-chan *CreateOrderResponse, <-chan error) {
responseChan := make(chan *CreateOrderResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateOrder(request)
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

// CreateOrderWithCallback invokes the rds.CreateOrder API asynchronously
// api document: https://help.aliyun.com/api/rds/createorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrderWithCallback(request *CreateOrderRequest, callback func(response *CreateOrderResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateOrderResponse
var err error
defer close(result)
response, err = client.CreateOrder(request)
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

// CreateOrderRequest is the request struct for api CreateOrder
type CreateOrderRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    SystemDBCharset     string `position:"Query" name:"SystemDBCharset"`
                    CountryCode     string `position:"Query" name:"CountryCode"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    CurrencyCode     string `position:"Query" name:"CurrencyCode"`
                    ResourceGroupId     string `position:"Query" name:"ResourceGroupId"`
                    DBInstanceDescription     string `position:"Query" name:"DBInstanceDescription"`
                    BusinessInfo     string `position:"Query" name:"BusinessInfo"`
                    AgentId     string `position:"Query" name:"AgentId"`
                    Resource     string `position:"Query" name:"Resource"`
                    BackupId     string `position:"Query" name:"BackupId"`
                    CommodityCode     string `position:"Query" name:"CommodityCode"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    SecurityIPList     string `position:"Query" name:"SecurityIPList"`
                    VSwitchId     string `position:"Query" name:"VSwitchId"`
                    PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
                    AutoRenew     string `position:"Query" name:"AutoRenew"`
                    PromotionCode     string `position:"Query" name:"PromotionCode"`
                    ZoneId     string `position:"Query" name:"ZoneId"`
                    TimeType     string `position:"Query" name:"TimeType"`
                    InstanceNetworkType     string `position:"Query" name:"InstanceNetworkType"`
                    ConnectionMode     string `position:"Query" name:"ConnectionMode"`
                    NodeType     string `position:"Query" name:"NodeType"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    ZoneIdSlave1     string `position:"Query" name:"ZoneIdSlave1"`
                    ZoneIdSlave2     string `position:"Query" name:"ZoneIdSlave2"`
                    Engine     string `position:"Query" name:"Engine"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    DBInstanceStorageType     string `position:"Query" name:"DBInstanceStorageType"`
                    DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
                    RestoreTime     string `position:"Query" name:"RestoreTime"`
                    Quantity     requests.Integer `position:"Query" name:"Quantity"`
                    AutoPay     requests.Boolean `position:"Query" name:"AutoPay"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    UsedTime     string `position:"Query" name:"UsedTime"`
                    InstanceUsedType     requests.Integer `position:"Query" name:"InstanceUsedType"`
                    VPCId     string `position:"Query" name:"VPCId"`
                    Category     string `position:"Query" name:"Category"`
                    PayType     string `position:"Query" name:"PayType"`
}


// CreateOrderResponse is the response struct for api CreateOrder
type CreateOrderResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     int `json:"OrderId" xml:"OrderId"`
}

// CreateCreateOrderRequest creates a request to invoke CreateOrder API
func CreateCreateOrderRequest() (request *CreateOrderRequest) {
request = &CreateOrderRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreateOrder", "", "")
return
}

// CreateCreateOrderResponse creates a response to parse from CreateOrder response
func CreateCreateOrderResponse() (response *CreateOrderResponse) {
response = &CreateOrderResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


