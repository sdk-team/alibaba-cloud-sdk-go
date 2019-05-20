
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

// DescribeKMSKeys invokes the ecs.DescribeKMSKeys API synchronously
// api document: https://help.aliyun.com/api/ecs/describekmskeys.html
func (client *Client) DescribeKMSKeys(request *DescribeKMSKeysRequest) (response *DescribeKMSKeysResponse, err error) {
response = CreateDescribeKMSKeysResponse()
err = client.DoAction(request, response)
return
}

// DescribeKMSKeysWithChan invokes the ecs.DescribeKMSKeys API asynchronously
// api document: https://help.aliyun.com/api/ecs/describekmskeys.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKMSKeysWithChan(request *DescribeKMSKeysRequest) (<-chan *DescribeKMSKeysResponse, <-chan error) {
responseChan := make(chan *DescribeKMSKeysResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeKMSKeys(request)
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

// DescribeKMSKeysWithCallback invokes the ecs.DescribeKMSKeys API asynchronously
// api document: https://help.aliyun.com/api/ecs/describekmskeys.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKMSKeysWithCallback(request *DescribeKMSKeysRequest, callback func(response *DescribeKMSKeysResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeKMSKeysResponse
var err error
defer close(result)
response, err = client.DescribeKMSKeys(request)
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

// DescribeKMSKeysRequest is the request struct for api DescribeKMSKeys
type DescribeKMSKeysRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Channel     string `position:"Query" name:"channel"`
                    Operator     string `position:"Query" name:"operator"`
                    PageNumber     string `position:"Query" name:"PageNumber"`
                    PageSize     string `position:"Query" name:"PageSize"`
                    ProxyId     string `position:"Query" name:"proxyId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    Token     string `position:"Query" name:"token"`
                    AppKey     string `position:"Query" name:"appKey"`
}


// DescribeKMSKeysResponse is the response struct for api DescribeKMSKeys
type DescribeKMSKeysResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
                KMSKeyIds KMSKeyIds `json:"KMSKeyIds" xml:"KMSKeyIds"`
}

// CreateDescribeKMSKeysRequest creates a request to invoke DescribeKMSKeys API
func CreateDescribeKMSKeysRequest() (request *DescribeKMSKeysRequest) {
request = &DescribeKMSKeysRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "DescribeKMSKeys", "ecs", "openAPI")
return
}

// CreateDescribeKMSKeysResponse creates a response to parse from DescribeKMSKeys response
func CreateDescribeKMSKeysResponse() (response *DescribeKMSKeysResponse) {
response = &DescribeKMSKeysResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


