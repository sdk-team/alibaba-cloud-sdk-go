
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

// DescribeSQLServerUpgradeVersions invokes the rds.DescribeSQLServerUpgradeVersions API synchronously
// api document: https://help.aliyun.com/api/rds/describesqlserverupgradeversions.html
func (client *Client) DescribeSQLServerUpgradeVersions(request *DescribeSQLServerUpgradeVersionsRequest) (response *DescribeSQLServerUpgradeVersionsResponse, err error) {
response = CreateDescribeSQLServerUpgradeVersionsResponse()
err = client.DoAction(request, response)
return
}

// DescribeSQLServerUpgradeVersionsWithChan invokes the rds.DescribeSQLServerUpgradeVersions API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqlserverupgradeversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLServerUpgradeVersionsWithChan(request *DescribeSQLServerUpgradeVersionsRequest) (<-chan *DescribeSQLServerUpgradeVersionsResponse, <-chan error) {
responseChan := make(chan *DescribeSQLServerUpgradeVersionsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeSQLServerUpgradeVersions(request)
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

// DescribeSQLServerUpgradeVersionsWithCallback invokes the rds.DescribeSQLServerUpgradeVersions API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqlserverupgradeversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLServerUpgradeVersionsWithCallback(request *DescribeSQLServerUpgradeVersionsRequest, callback func(response *DescribeSQLServerUpgradeVersionsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeSQLServerUpgradeVersionsResponse
var err error
defer close(result)
response, err = client.DescribeSQLServerUpgradeVersions(request)
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

// DescribeSQLServerUpgradeVersionsRequest is the request struct for api DescribeSQLServerUpgradeVersions
type DescribeSQLServerUpgradeVersionsRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// DescribeSQLServerUpgradeVersionsResponse is the response struct for api DescribeSQLServerUpgradeVersions
type DescribeSQLServerUpgradeVersionsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                    Items ItemsInDescribeSQLServerUpgradeVersions `json:"Items" xml:"Items"`
}

// CreateDescribeSQLServerUpgradeVersionsRequest creates a request to invoke DescribeSQLServerUpgradeVersions API
func CreateDescribeSQLServerUpgradeVersionsRequest() (request *DescribeSQLServerUpgradeVersionsRequest) {
request = &DescribeSQLServerUpgradeVersionsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLServerUpgradeVersions", "", "")
return
}

// CreateDescribeSQLServerUpgradeVersionsResponse creates a response to parse from DescribeSQLServerUpgradeVersions response
func CreateDescribeSQLServerUpgradeVersionsResponse() (response *DescribeSQLServerUpgradeVersionsResponse) {
response = &DescribeSQLServerUpgradeVersionsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


