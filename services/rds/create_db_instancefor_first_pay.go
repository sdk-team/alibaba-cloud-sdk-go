
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

// CreateDBInstanceforFirstPay invokes the rds.CreateDBInstanceforFirstPay API synchronously
// api document: https://help.aliyun.com/api/rds/createdbinstanceforfirstpay.html
func (client *Client) CreateDBInstanceforFirstPay(request *CreateDBInstanceforFirstPayRequest) (response *CreateDBInstanceforFirstPayResponse, err error) {
response = CreateCreateDBInstanceforFirstPayResponse()
err = client.DoAction(request, response)
return
}

// CreateDBInstanceforFirstPayWithChan invokes the rds.CreateDBInstanceforFirstPay API asynchronously
// api document: https://help.aliyun.com/api/rds/createdbinstanceforfirstpay.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceforFirstPayWithChan(request *CreateDBInstanceforFirstPayRequest) (<-chan *CreateDBInstanceforFirstPayResponse, <-chan error) {
responseChan := make(chan *CreateDBInstanceforFirstPayResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateDBInstanceforFirstPay(request)
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

// CreateDBInstanceforFirstPayWithCallback invokes the rds.CreateDBInstanceforFirstPay API asynchronously
// api document: https://help.aliyun.com/api/rds/createdbinstanceforfirstpay.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceforFirstPayWithCallback(request *CreateDBInstanceforFirstPayRequest, callback func(response *CreateDBInstanceforFirstPayResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateDBInstanceforFirstPayResponse
var err error
defer close(result)
response, err = client.CreateDBInstanceforFirstPay(request)
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

// CreateDBInstanceforFirstPayRequest is the request struct for api CreateDBInstanceforFirstPay
type CreateDBInstanceforFirstPayRequest struct {
*requests.RpcRequest
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    EngineVersion     string `position:"Query" name:"EngineVersion"`
                    Uid     requests.Integer `position:"Query" name:"uid"`
                    BidLoginEmail     string `position:"Query" name:"bidLoginEmail"`
                    Engine     string `position:"Query" name:"Engine"`
                    UidLoginEmail     string `position:"Query" name:"uidLoginEmail"`
                    DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    DBInstanceRemarks     string `position:"Query" name:"DBInstanceRemarks"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    Bid     string `position:"Query" name:"bid"`
                    CharacterSetName     string `position:"Query" name:"CharacterSetName"`
}


// CreateDBInstanceforFirstPayResponse is the response struct for api CreateDBInstanceforFirstPay
type CreateDBInstanceforFirstPayResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
            Port     string `json:"Port" xml:"Port"`
}

// CreateCreateDBInstanceforFirstPayRequest creates a request to invoke CreateDBInstanceforFirstPay API
func CreateCreateDBInstanceforFirstPayRequest() (request *CreateDBInstanceforFirstPayRequest) {
request = &CreateDBInstanceforFirstPayRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstanceforFirstPay", "rds", "openAPI")
return
}

// CreateCreateDBInstanceforFirstPayResponse creates a response to parse from CreateDBInstanceforFirstPay response
func CreateCreateDBInstanceforFirstPayResponse() (response *CreateDBInstanceforFirstPayResponse) {
response = &CreateDBInstanceforFirstPayResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


