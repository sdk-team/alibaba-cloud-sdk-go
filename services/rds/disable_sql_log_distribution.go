
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

// DisableSqlLogDistribution invokes the rds.DisableSqlLogDistribution API synchronously
// api document: https://help.aliyun.com/api/rds/disablesqllogdistribution.html
func (client *Client) DisableSqlLogDistribution(request *DisableSqlLogDistributionRequest) (response *DisableSqlLogDistributionResponse, err error) {
response = CreateDisableSqlLogDistributionResponse()
err = client.DoAction(request, response)
return
}

// DisableSqlLogDistributionWithChan invokes the rds.DisableSqlLogDistribution API asynchronously
// api document: https://help.aliyun.com/api/rds/disablesqllogdistribution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableSqlLogDistributionWithChan(request *DisableSqlLogDistributionRequest) (<-chan *DisableSqlLogDistributionResponse, <-chan error) {
responseChan := make(chan *DisableSqlLogDistributionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DisableSqlLogDistribution(request)
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

// DisableSqlLogDistributionWithCallback invokes the rds.DisableSqlLogDistribution API asynchronously
// api document: https://help.aliyun.com/api/rds/disablesqllogdistribution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableSqlLogDistributionWithCallback(request *DisableSqlLogDistributionRequest, callback func(response *DisableSqlLogDistributionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DisableSqlLogDistributionResponse
var err error
defer close(result)
response, err = client.DisableSqlLogDistribution(request)
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

// DisableSqlLogDistributionRequest is the request struct for api DisableSqlLogDistribution
type DisableSqlLogDistributionRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DisableSqlLogDistributionResponse is the response struct for api DisableSqlLogDistribution
type DisableSqlLogDistributionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceID     int `json:"DBInstanceID" xml:"DBInstanceID"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateDisableSqlLogDistributionRequest creates a request to invoke DisableSqlLogDistribution API
func CreateDisableSqlLogDistributionRequest() (request *DisableSqlLogDistributionRequest) {
request = &DisableSqlLogDistributionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DisableSqlLogDistribution", "", "")
return
}

// CreateDisableSqlLogDistributionResponse creates a response to parse from DisableSqlLogDistribution response
func CreateDisableSqlLogDistributionResponse() (response *DisableSqlLogDistributionResponse) {
response = &DisableSqlLogDistributionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


