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

// StopAxtTask invokes the ecs.StopAxtTask API synchronously
func (client *Client) StopAxtTask(request *StopAxtTaskRequest) (response *StopAxtTaskResponse, err error) {
	response = CreateStopAxtTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopAxtTaskWithChan invokes the ecs.StopAxtTask API asynchronously
func (client *Client) StopAxtTaskWithChan(request *StopAxtTaskRequest) (<-chan *StopAxtTaskResponse, <-chan error) {
	responseChan := make(chan *StopAxtTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopAxtTask(request)
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

// StopAxtTaskWithCallback invokes the ecs.StopAxtTask API asynchronously
func (client *Client) StopAxtTaskWithCallback(request *StopAxtTaskRequest, callback func(response *StopAxtTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopAxtTaskResponse
		var err error
		defer close(result)
		response, err = client.StopAxtTask(request)
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

// StopAxtTaskRequest is the request struct for api StopAxtTask
type StopAxtTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
}

// StopAxtTaskResponse is the response struct for api StopAxtTask
type StopAxtTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopAxtTaskRequest creates a request to invoke StopAxtTask API
func CreateStopAxtTaskRequest() (request *StopAxtTaskRequest) {
	request = &StopAxtTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "StopAxtTask", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopAxtTaskResponse creates a response to parse from StopAxtTask response
func CreateStopAxtTaskResponse() (response *StopAxtTaskResponse) {
	response = &StopAxtTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
