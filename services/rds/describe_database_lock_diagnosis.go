
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

// DescribeDatabaseLockDiagnosis invokes the rds.DescribeDatabaseLockDiagnosis API synchronously
// api document: https://help.aliyun.com/api/rds/describedatabaselockdiagnosis.html
func (client *Client) DescribeDatabaseLockDiagnosis(request *DescribeDatabaseLockDiagnosisRequest) (response *DescribeDatabaseLockDiagnosisResponse, err error) {
response = CreateDescribeDatabaseLockDiagnosisResponse()
err = client.DoAction(request, response)
return
}

// DescribeDatabaseLockDiagnosisWithChan invokes the rds.DescribeDatabaseLockDiagnosis API asynchronously
// api document: https://help.aliyun.com/api/rds/describedatabaselockdiagnosis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDatabaseLockDiagnosisWithChan(request *DescribeDatabaseLockDiagnosisRequest) (<-chan *DescribeDatabaseLockDiagnosisResponse, <-chan error) {
responseChan := make(chan *DescribeDatabaseLockDiagnosisResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDatabaseLockDiagnosis(request)
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

// DescribeDatabaseLockDiagnosisWithCallback invokes the rds.DescribeDatabaseLockDiagnosis API asynchronously
// api document: https://help.aliyun.com/api/rds/describedatabaselockdiagnosis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDatabaseLockDiagnosisWithCallback(request *DescribeDatabaseLockDiagnosisRequest, callback func(response *DescribeDatabaseLockDiagnosisResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDatabaseLockDiagnosisResponse
var err error
defer close(result)
response, err = client.DescribeDatabaseLockDiagnosis(request)
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

// DescribeDatabaseLockDiagnosisRequest is the request struct for api DescribeDatabaseLockDiagnosis
type DescribeDatabaseLockDiagnosisRequest struct {
*requests.RpcRequest
                    EndTime     string `position:"Query" name:"EndTime"`
                    StartTime     string `position:"Query" name:"StartTime"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// DescribeDatabaseLockDiagnosisResponse is the response struct for api DescribeDatabaseLockDiagnosis
type DescribeDatabaseLockDiagnosisResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                DeadLockList DeadLockList `json:"DeadLockList" xml:"DeadLockList"`
}

// CreateDescribeDatabaseLockDiagnosisRequest creates a request to invoke DescribeDatabaseLockDiagnosis API
func CreateDescribeDatabaseLockDiagnosisRequest() (request *DescribeDatabaseLockDiagnosisRequest) {
request = &DescribeDatabaseLockDiagnosisRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDatabaseLockDiagnosis", "rds", "openAPI")
return
}

// CreateDescribeDatabaseLockDiagnosisResponse creates a response to parse from DescribeDatabaseLockDiagnosis response
func CreateDescribeDatabaseLockDiagnosisResponse() (response *DescribeDatabaseLockDiagnosisResponse) {
response = &DescribeDatabaseLockDiagnosisResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


