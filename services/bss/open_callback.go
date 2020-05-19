package bss

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

// OpenCallback invokes the bss.OpenCallback API synchronously
// api document: https://help.aliyun.com/api/bss/opencallback.html
func (client *Client) OpenCallback(request *OpenCallbackRequest) (response *OpenCallbackResponse, err error) {
	response = CreateOpenCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// OpenCallbackWithChan invokes the bss.OpenCallback API asynchronously
// api document: https://help.aliyun.com/api/bss/opencallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OpenCallbackWithChan(request *OpenCallbackRequest) (<-chan *OpenCallbackResponse, <-chan error) {
	responseChan := make(chan *OpenCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenCallback(request)
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

// OpenCallbackWithCallback invokes the bss.OpenCallback API asynchronously
// api document: https://help.aliyun.com/api/bss/opencallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OpenCallbackWithCallback(request *OpenCallbackRequest, callback func(response *OpenCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenCallbackResponse
		var err error
		defer close(result)
		response, err = client.OpenCallback(request)
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

// OpenCallbackRequest is the request struct for api OpenCallback
type OpenCallbackRequest struct {
	*requests.RpcRequest
	ParamStr string `position:"Query" name:"paramStr"`
}

// OpenCallbackResponse is the response struct for api OpenCallback
type OpenCallbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateOpenCallbackRequest creates a request to invoke OpenCallback API
func CreateOpenCallbackRequest() (request *OpenCallbackRequest) {
	request = &OpenCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Bss", "2014-07-14", "OpenCallback", "", "")
	return
}

// CreateOpenCallbackResponse creates a response to parse from OpenCallback response
func CreateOpenCallbackResponse() (response *OpenCallbackResponse) {
	response = &OpenCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
