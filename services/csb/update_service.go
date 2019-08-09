package csb

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

// UpdateService invokes the csb.UpdateService API synchronously
// api document: https://help.aliyun.com/api/csb/updateservice.html
func (client *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
	response = CreateUpdateServiceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServiceWithChan invokes the csb.UpdateService API asynchronously
// api document: https://help.aliyun.com/api/csb/updateservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateServiceWithChan(request *UpdateServiceRequest) (<-chan *UpdateServiceResponse, <-chan error) {
	responseChan := make(chan *UpdateServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateService(request)
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

// UpdateServiceWithCallback invokes the csb.UpdateService API asynchronously
// api document: https://help.aliyun.com/api/csb/updateservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateServiceWithCallback(request *UpdateServiceRequest, callback func(response *UpdateServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServiceResponse
		var err error
		defer close(result)
		response, err = client.UpdateService(request)
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

// UpdateServiceRequest is the request struct for api UpdateService
type UpdateServiceRequest struct {
	*requests.RpcRequest
	Data  string           `position:"Body" name:"Data"`
	CsbId requests.Integer `position:"Query" name:"CsbId"`
}

// UpdateServiceResponse is the response struct for api UpdateService
type UpdateServiceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateServiceRequest creates a request to invoke UpdateService API
func CreateUpdateServiceRequest() (request *UpdateServiceRequest) {
	request = &UpdateServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "UpdateService", "csb", "openAPI")
	return
}

// CreateUpdateServiceResponse creates a response to parse from UpdateService response
func CreateUpdateServiceResponse() (response *UpdateServiceResponse) {
	response = &UpdateServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
