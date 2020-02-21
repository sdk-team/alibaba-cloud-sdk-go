package r_kvstore

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

// DeactivateInstance invokes the r_kvstore.DeactivateInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/deactivateinstance.html
func (client *Client) DeactivateInstance(request *DeactivateInstanceRequest) (response *DeactivateInstanceResponse, err error) {
	response = CreateDeactivateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DeactivateInstanceWithChan invokes the r_kvstore.DeactivateInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/deactivateinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeactivateInstanceWithChan(request *DeactivateInstanceRequest) (<-chan *DeactivateInstanceResponse, <-chan error) {
	responseChan := make(chan *DeactivateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeactivateInstance(request)
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

// DeactivateInstanceWithCallback invokes the r_kvstore.DeactivateInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/deactivateinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeactivateInstanceWithCallback(request *DeactivateInstanceRequest, callback func(response *DeactivateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeactivateInstanceResponse
		var err error
		defer close(result)
		response, err = client.DeactivateInstance(request)
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

// DeactivateInstanceRequest is the request struct for api DeactivateInstance
type DeactivateInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DeactivateInstanceResponse is the response struct for api DeactivateInstance
type DeactivateInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeactivateInstanceRequest creates a request to invoke DeactivateInstance API
func CreateDeactivateInstanceRequest() (request *DeactivateInstanceRequest) {
	request = &DeactivateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DeactivateInstance", "", "")
	return
}

// CreateDeactivateInstanceResponse creates a response to parse from DeactivateInstance response
func CreateDeactivateInstanceResponse() (response *DeactivateInstanceResponse) {
	response = &DeactivateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
