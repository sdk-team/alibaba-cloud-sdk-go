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

// DescribeRegionsTest invokes the rds.DescribeRegionsTest API synchronously
// api document: https://help.aliyun.com/api/rds/describeregionstest.html
func (client *Client) DescribeRegionsTest(request *DescribeRegionsTestRequest) (response *DescribeRegionsTestResponse, err error) {
	response = CreateDescribeRegionsTestResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRegionsTestWithChan invokes the rds.DescribeRegionsTest API asynchronously
// api document: https://help.aliyun.com/api/rds/describeregionstest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRegionsTestWithChan(request *DescribeRegionsTestRequest) (<-chan *DescribeRegionsTestResponse, <-chan error) {
	responseChan := make(chan *DescribeRegionsTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRegionsTest(request)
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

// DescribeRegionsTestWithCallback invokes the rds.DescribeRegionsTest API asynchronously
// api document: https://help.aliyun.com/api/rds/describeregionstest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRegionsTestWithCallback(request *DescribeRegionsTestRequest, callback func(response *DescribeRegionsTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRegionsTestResponse
		var err error
		defer close(result)
		response, err = client.DescribeRegionsTest(request)
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

// DescribeRegionsTestRequest is the request struct for api DescribeRegionsTest
type DescribeRegionsTestRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"resourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"resourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"ownerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"ownerId"`
}

// DescribeRegionsTestResponse is the response struct for api DescribeRegionsTest
type DescribeRegionsTestResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Regions   RegionsInDescribeRegionsTest `json:"Regions" xml:"Regions"`
}

// CreateDescribeRegionsTestRequest creates a request to invoke DescribeRegionsTest API
func CreateDescribeRegionsTestRequest() (request *DescribeRegionsTestRequest) {
	request = &DescribeRegionsTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeRegionsTest", "rds", "openAPI")
	return
}

// CreateDescribeRegionsTestResponse creates a response to parse from DescribeRegionsTest response
func CreateDescribeRegionsTestResponse() (response *DescribeRegionsTestResponse) {
	response = &DescribeRegionsTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
