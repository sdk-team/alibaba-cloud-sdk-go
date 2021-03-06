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

// DescribeInstanceKeyPairInfo invokes the ecs.DescribeInstanceKeyPairInfo API synchronously
func (client *Client) DescribeInstanceKeyPairInfo(request *DescribeInstanceKeyPairInfoRequest) (response *DescribeInstanceKeyPairInfoResponse, err error) {
	response = CreateDescribeInstanceKeyPairInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceKeyPairInfoWithChan invokes the ecs.DescribeInstanceKeyPairInfo API asynchronously
func (client *Client) DescribeInstanceKeyPairInfoWithChan(request *DescribeInstanceKeyPairInfoRequest) (<-chan *DescribeInstanceKeyPairInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceKeyPairInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceKeyPairInfo(request)
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

// DescribeInstanceKeyPairInfoWithCallback invokes the ecs.DescribeInstanceKeyPairInfo API asynchronously
func (client *Client) DescribeInstanceKeyPairInfoWithCallback(request *DescribeInstanceKeyPairInfoRequest, callback func(response *DescribeInstanceKeyPairInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceKeyPairInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceKeyPairInfo(request)
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

// DescribeInstanceKeyPairInfoRequest is the request struct for api DescribeInstanceKeyPairInfo
type DescribeInstanceKeyPairInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairName          string           `position:"Query" name:"KeyPairName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeInstanceKeyPairInfoResponse is the response struct for api DescribeInstanceKeyPairInfo
type DescribeInstanceKeyPairInfoResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	TotalCount          int                 `json:"TotalCount" xml:"TotalCount"`
	PageNumber          int                 `json:"PageNumber" xml:"PageNumber"`
	PageSize            int                 `json:"PageSize" xml:"PageSize"`
	KeyPairInstanceList KeyPairInstanceList `json:"KeyPairInstanceList" xml:"KeyPairInstanceList"`
}

// CreateDescribeInstanceKeyPairInfoRequest creates a request to invoke DescribeInstanceKeyPairInfo API
func CreateDescribeInstanceKeyPairInfoRequest() (request *DescribeInstanceKeyPairInfoRequest) {
	request = &DescribeInstanceKeyPairInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceKeyPairInfo", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceKeyPairInfoResponse creates a response to parse from DescribeInstanceKeyPairInfo response
func CreateDescribeInstanceKeyPairInfoResponse() (response *DescribeInstanceKeyPairInfoResponse) {
	response = &DescribeInstanceKeyPairInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
