
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

// DescribeDBInstanceNetInfoForChannel invokes the rds.DescribeDBInstanceNetInfoForChannel API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancenetinfoforchannel.html
func (client *Client) DescribeDBInstanceNetInfoForChannel(request *DescribeDBInstanceNetInfoForChannelRequest) (response *DescribeDBInstanceNetInfoForChannelResponse, err error) {
response = CreateDescribeDBInstanceNetInfoForChannelResponse()
err = client.DoAction(request, response)
return
}

// DescribeDBInstanceNetInfoForChannelWithChan invokes the rds.DescribeDBInstanceNetInfoForChannel API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancenetinfoforchannel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceNetInfoForChannelWithChan(request *DescribeDBInstanceNetInfoForChannelRequest) (<-chan *DescribeDBInstanceNetInfoForChannelResponse, <-chan error) {
responseChan := make(chan *DescribeDBInstanceNetInfoForChannelResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDBInstanceNetInfoForChannel(request)
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

// DescribeDBInstanceNetInfoForChannelWithCallback invokes the rds.DescribeDBInstanceNetInfoForChannel API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancenetinfoforchannel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceNetInfoForChannelWithCallback(request *DescribeDBInstanceNetInfoForChannelRequest, callback func(response *DescribeDBInstanceNetInfoForChannelResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDBInstanceNetInfoForChannelResponse
var err error
defer close(result)
response, err = client.DescribeDBInstanceNetInfoForChannel(request)
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

// DescribeDBInstanceNetInfoForChannelRequest is the request struct for api DescribeDBInstanceNetInfoForChannel
type DescribeDBInstanceNetInfoForChannelRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    Flag     string `position:"Query" name:"Flag"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceNetRWSplitType     string `position:"Query" name:"DBInstanceNetRWSplitType"`
}


// DescribeDBInstanceNetInfoForChannelResponse is the response struct for api DescribeDBInstanceNetInfoForChannel
type DescribeDBInstanceNetInfoForChannelResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            InstanceNetworkType     string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
                    DBInstanceNetInfos DBInstanceNetInfosInDescribeDBInstanceNetInfoForChannel `json:"DBInstanceNetInfos" xml:"DBInstanceNetInfos"`
}

// CreateDescribeDBInstanceNetInfoForChannelRequest creates a request to invoke DescribeDBInstanceNetInfoForChannel API
func CreateDescribeDBInstanceNetInfoForChannelRequest() (request *DescribeDBInstanceNetInfoForChannelRequest) {
request = &DescribeDBInstanceNetInfoForChannelRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceNetInfoForChannel", "", "")
return
}

// CreateDescribeDBInstanceNetInfoForChannelResponse creates a response to parse from DescribeDBInstanceNetInfoForChannel response
func CreateDescribeDBInstanceNetInfoForChannelResponse() (response *DescribeDBInstanceNetInfoForChannelResponse) {
response = &DescribeDBInstanceNetInfoForChannelResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


