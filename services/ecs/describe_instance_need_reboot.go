
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

// DescribeInstanceNeedReboot invokes the ecs.DescribeInstanceNeedReboot API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstanceneedreboot.html
func (client *Client) DescribeInstanceNeedReboot(request *DescribeInstanceNeedRebootRequest) (response *DescribeInstanceNeedRebootResponse, err error) {
response = CreateDescribeInstanceNeedRebootResponse()
err = client.DoAction(request, response)
return
}

// DescribeInstanceNeedRebootWithChan invokes the ecs.DescribeInstanceNeedReboot API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstanceneedreboot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceNeedRebootWithChan(request *DescribeInstanceNeedRebootRequest) (<-chan *DescribeInstanceNeedRebootResponse, <-chan error) {
responseChan := make(chan *DescribeInstanceNeedRebootResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeInstanceNeedReboot(request)
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

// DescribeInstanceNeedRebootWithCallback invokes the ecs.DescribeInstanceNeedReboot API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstanceneedreboot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceNeedRebootWithCallback(request *DescribeInstanceNeedRebootRequest, callback func(response *DescribeInstanceNeedRebootResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeInstanceNeedRebootResponse
var err error
defer close(result)
response, err = client.DescribeInstanceNeedReboot(request)
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

// DescribeInstanceNeedRebootRequest is the request struct for api DescribeInstanceNeedReboot
type DescribeInstanceNeedRebootRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
                    PageSize     requests.Integer `position:"Query" name:"PageSize"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceIds  *[]string `position:"Query" name:"InstanceIds"  type:"Repeated"`
}


// DescribeInstanceNeedRebootResponse is the response struct for api DescribeInstanceNeedReboot
type DescribeInstanceNeedRebootResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
                InstanceInfo InstanceInfo `json:"InstanceInfo" xml:"InstanceInfo"`
}

// CreateDescribeInstanceNeedRebootRequest creates a request to invoke DescribeInstanceNeedReboot API
func CreateDescribeInstanceNeedRebootRequest() (request *DescribeInstanceNeedRebootRequest) {
request = &DescribeInstanceNeedRebootRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2016-03-14", "DescribeInstanceNeedReboot", "ecs", "openAPI")
return
}

// CreateDescribeInstanceNeedRebootResponse creates a response to parse from DescribeInstanceNeedReboot response
func CreateDescribeInstanceNeedRebootResponse() (response *DescribeInstanceNeedRebootResponse) {
response = &DescribeInstanceNeedRebootResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


