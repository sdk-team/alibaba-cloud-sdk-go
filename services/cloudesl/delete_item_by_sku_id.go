package cloudesl

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

// DeleteItemBySkuId invokes the cloudesl.DeleteItemBySkuId API synchronously
// api document: https://help.aliyun.com/api/cloudesl/deleteitembyskuid.html
func (client *Client) DeleteItemBySkuId(request *DeleteItemBySkuIdRequest) (response *DeleteItemBySkuIdResponse, err error) {
	response = CreateDeleteItemBySkuIdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteItemBySkuIdWithChan invokes the cloudesl.DeleteItemBySkuId API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/deleteitembyskuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteItemBySkuIdWithChan(request *DeleteItemBySkuIdRequest) (<-chan *DeleteItemBySkuIdResponse, <-chan error) {
	responseChan := make(chan *DeleteItemBySkuIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteItemBySkuId(request)
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

// DeleteItemBySkuIdWithCallback invokes the cloudesl.DeleteItemBySkuId API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/deleteitembyskuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteItemBySkuIdWithCallback(request *DeleteItemBySkuIdRequest, callback func(response *DeleteItemBySkuIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteItemBySkuIdResponse
		var err error
		defer close(result)
		response, err = client.DeleteItemBySkuId(request)
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

// DeleteItemBySkuIdRequest is the request struct for api DeleteItemBySkuId
type DeleteItemBySkuIdRequest struct {
	*requests.RpcRequest
	StoreId string `position:"Query" name:"StoreId"`
	SkuId   string `position:"Query" name:"SkuId"`
}

// DeleteItemBySkuIdResponse is the response struct for api DeleteItemBySkuId
type DeleteItemBySkuIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateDeleteItemBySkuIdRequest creates a request to invoke DeleteItemBySkuId API
func CreateDeleteItemBySkuIdRequest() (request *DeleteItemBySkuIdRequest) {
	request = &DeleteItemBySkuIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2018-08-01", "DeleteItemBySkuId", "", "")
	return
}

// CreateDeleteItemBySkuIdResponse creates a response to parse from DeleteItemBySkuId response
func CreateDeleteItemBySkuIdResponse() (response *DeleteItemBySkuIdResponse) {
	response = &DeleteItemBySkuIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
