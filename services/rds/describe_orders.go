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

// DescribeOrders invokes the rds.DescribeOrders API synchronously
// api document: https://help.aliyun.com/api/rds/describeorders.html
func (client *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
	response = CreateDescribeOrdersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOrdersWithChan invokes the rds.DescribeOrders API asynchronously
// api document: https://help.aliyun.com/api/rds/describeorders.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrdersWithChan(request *DescribeOrdersRequest) (<-chan *DescribeOrdersResponse, <-chan error) {
	responseChan := make(chan *DescribeOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOrders(request)
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

// DescribeOrdersWithCallback invokes the rds.DescribeOrders API asynchronously
// api document: https://help.aliyun.com/api/rds/describeorders.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrdersWithCallback(request *DescribeOrdersRequest, callback func(response *DescribeOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOrdersResponse
		var err error
		defer close(result)
		response, err = client.DescribeOrders(request)
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

// DescribeOrdersRequest is the request struct for api DescribeOrders
type DescribeOrdersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OrderStatus          string           `position:"Query" name:"OrderStatus"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OrderId              string           `position:"Query" name:"OrderId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

// DescribeOrdersResponse is the response struct for api DescribeOrders
type DescribeOrdersResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
}

// CreateDescribeOrdersRequest creates a request to invoke DescribeOrders API
func CreateDescribeOrdersRequest() (request *DescribeOrdersRequest) {
	request = &DescribeOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOrders", "", "")
	return
}

// CreateDescribeOrdersResponse creates a response to parse from DescribeOrders response
func CreateDescribeOrdersResponse() (response *DescribeOrdersResponse) {
	response = &DescribeOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
