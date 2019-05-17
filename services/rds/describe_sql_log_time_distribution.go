
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

// DescribeSqlLogTimeDistribution invokes the rds.DescribeSqlLogTimeDistribution API synchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtimedistribution.html
func (client *Client) DescribeSqlLogTimeDistribution(request *DescribeSqlLogTimeDistributionRequest) (response *DescribeSqlLogTimeDistributionResponse, err error) {
response = CreateDescribeSqlLogTimeDistributionResponse()
err = client.DoAction(request, response)
return
}

// DescribeSqlLogTimeDistributionWithChan invokes the rds.DescribeSqlLogTimeDistribution API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtimedistribution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSqlLogTimeDistributionWithChan(request *DescribeSqlLogTimeDistributionRequest) (<-chan *DescribeSqlLogTimeDistributionResponse, <-chan error) {
responseChan := make(chan *DescribeSqlLogTimeDistributionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeSqlLogTimeDistribution(request)
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

// DescribeSqlLogTimeDistributionWithCallback invokes the rds.DescribeSqlLogTimeDistribution API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtimedistribution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSqlLogTimeDistributionWithCallback(request *DescribeSqlLogTimeDistributionRequest, callback func(response *DescribeSqlLogTimeDistributionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeSqlLogTimeDistributionResponse
var err error
defer close(result)
response, err = client.DescribeSqlLogTimeDistribution(request)
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

// DescribeSqlLogTimeDistributionRequest is the request struct for api DescribeSqlLogTimeDistribution
type DescribeSqlLogTimeDistributionRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    JobId     string `position:"Query" name:"JobId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    EndTime     string `position:"Query" name:"EndTime"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    TimeLevel     string `position:"Query" name:"TimeLevel"`
}


// DescribeSqlLogTimeDistributionResponse is the response struct for api DescribeSqlLogTimeDistribution
type DescribeSqlLogTimeDistributionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceID     int `json:"DBInstanceID" xml:"DBInstanceID"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            TemplateHash     int `json:"TemplateHash" xml:"TemplateHash"`
            JobId     string `json:"JobId" xml:"JobId"`
            Finish     string `json:"Finish" xml:"Finish"`
            ItemsNumbers     int `json:"ItemsNumbers" xml:"ItemsNumbers"`
                    Items ItemsInDescribeSqlLogTimeDistribution `json:"Items" xml:"Items"`
}

// CreateDescribeSqlLogTimeDistributionRequest creates a request to invoke DescribeSqlLogTimeDistribution API
func CreateDescribeSqlLogTimeDistributionRequest() (request *DescribeSqlLogTimeDistributionRequest) {
request = &DescribeSqlLogTimeDistributionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSqlLogTimeDistribution", "", "")
return
}

// CreateDescribeSqlLogTimeDistributionResponse creates a response to parse from DescribeSqlLogTimeDistribution response
func CreateDescribeSqlLogTimeDistributionResponse() (response *DescribeSqlLogTimeDistributionResponse) {
response = &DescribeSqlLogTimeDistributionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


