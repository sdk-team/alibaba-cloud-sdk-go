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

// DeleteContainerInstance invokes the ecs.DeleteContainerInstance API synchronously
func (client *Client) DeleteContainerInstance(request *DeleteContainerInstanceRequest) (response *DeleteContainerInstanceResponse, err error) {
	response = CreateDeleteContainerInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteContainerInstanceWithChan invokes the ecs.DeleteContainerInstance API asynchronously
func (client *Client) DeleteContainerInstanceWithChan(request *DeleteContainerInstanceRequest) (<-chan *DeleteContainerInstanceResponse, <-chan error) {
	responseChan := make(chan *DeleteContainerInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteContainerInstance(request)
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

// DeleteContainerInstanceWithCallback invokes the ecs.DeleteContainerInstance API asynchronously
func (client *Client) DeleteContainerInstanceWithCallback(request *DeleteContainerInstanceRequest, callback func(response *DeleteContainerInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteContainerInstanceResponse
		var err error
		defer close(result)
		response, err = client.DeleteContainerInstance(request)
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

// DeleteContainerInstanceRequest is the request struct for api DeleteContainerInstance
type DeleteContainerInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TerminateSubscription requests.Boolean `position:"Query" name:"TerminateSubscription"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId            string           `position:"Query" name:"InstanceId"`
	Force                 requests.Boolean `position:"Query" name:"Force"`
}

// DeleteContainerInstanceResponse is the response struct for api DeleteContainerInstance
type DeleteContainerInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteContainerInstanceRequest creates a request to invoke DeleteContainerInstance API
func CreateDeleteContainerInstanceRequest() (request *DeleteContainerInstanceRequest) {
	request = &DeleteContainerInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteContainerInstance", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteContainerInstanceResponse creates a response to parse from DeleteContainerInstance response
func CreateDeleteContainerInstanceResponse() (response *DeleteContainerInstanceResponse) {
	response = &DeleteContainerInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
