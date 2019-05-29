
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

// DescribeSqlLogTemplatesTimeDistribution invokes the rds.DescribeSqlLogTemplatesTimeDistribution API synchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtemplatestimedistribution.html
func (client *Client) DescribeSqlLogTemplatesTimeDistribution(request *DescribeSqlLogTemplatesTimeDistributionRequest) (response *DescribeSqlLogTemplatesTimeDistributionResponse, err error) {
response = CreateDescribeSqlLogTemplatesTimeDistributionResponse()
err = client.DoAction(request, response)
return
}

// DescribeSqlLogTemplatesTimeDistributionWithChan invokes the rds.DescribeSqlLogTemplatesTimeDistribution API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtemplatestimedistribution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSqlLogTemplatesTimeDistributionWithChan(request *DescribeSqlLogTemplatesTimeDistributionRequest) (<-chan *DescribeSqlLogTemplatesTimeDistributionResponse, <-chan error) {
responseChan := make(chan *DescribeSqlLogTemplatesTimeDistributionResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeSqlLogTemplatesTimeDistribution(request)
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

// DescribeSqlLogTemplatesTimeDistributionWithCallback invokes the rds.DescribeSqlLogTemplatesTimeDistribution API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtemplatestimedistribution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSqlLogTemplatesTimeDistributionWithCallback(request *DescribeSqlLogTemplatesTimeDistributionRequest, callback func(response *DescribeSqlLogTemplatesTimeDistributionResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeSqlLogTemplatesTimeDistributionResponse
var err error
defer close(result)
response, err = client.DescribeSqlLogTemplatesTimeDistribution(request)
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

// DescribeSqlLogTemplatesTimeDistributionRequest is the request struct for api DescribeSqlLogTemplatesTimeDistribution
type DescribeSqlLogTemplatesTimeDistributionRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    JobId     string `position:"Query" name:"JobId"`
                    SortKey     string `position:"Query" name:"SortKey"`
                    TemplateHash     string `position:"Query" name:"TemplateHash"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    EndTime     string `position:"Query" name:"EndTime"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    TimeLevel     string `position:"Query" name:"TimeLevel"`
}


// DescribeSqlLogTemplatesTimeDistributionResponse is the response struct for api DescribeSqlLogTemplatesTimeDistribution
type DescribeSqlLogTemplatesTimeDistributionResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceID     int `json:"DBInstanceID" xml:"DBInstanceID"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
            StartTime     string `json:"StartTime" xml:"StartTime"`
            EndTime     string `json:"EndTime" xml:"EndTime"`
            ItemsNumbers     int `json:"ItemsNumbers" xml:"ItemsNumbers"`
            JobId     string `json:"JobId" xml:"JobId"`
            Finish     string `json:"Finish" xml:"Finish"`
                    Items ItemsInDescribeSqlLogTemplatesTimeDistribution `json:"Items" xml:"Items"`
}

// CreateDescribeSqlLogTemplatesTimeDistributionRequest creates a request to invoke DescribeSqlLogTemplatesTimeDistribution API
func CreateDescribeSqlLogTemplatesTimeDistributionRequest() (request *DescribeSqlLogTemplatesTimeDistributionRequest) {
request = &DescribeSqlLogTemplatesTimeDistributionRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSqlLogTemplatesTimeDistribution", "rds", "openAPI")
return
}

// CreateDescribeSqlLogTemplatesTimeDistributionResponse creates a response to parse from DescribeSqlLogTemplatesTimeDistribution response
func CreateDescribeSqlLogTemplatesTimeDistributionResponse() (response *DescribeSqlLogTemplatesTimeDistributionResponse) {
response = &DescribeSqlLogTemplatesTimeDistributionResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


