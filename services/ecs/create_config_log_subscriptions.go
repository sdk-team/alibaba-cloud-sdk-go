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

// CreateConfigLogSubscriptions invokes the ecs.CreateConfigLogSubscriptions API synchronously
// api document: https://help.aliyun.com/api/ecs/createconfiglogsubscriptions.html
func (client *Client) CreateConfigLogSubscriptions(request *CreateConfigLogSubscriptionsRequest) (response *CreateConfigLogSubscriptionsResponse, err error) {
	response = CreateCreateConfigLogSubscriptionsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateConfigLogSubscriptionsWithChan invokes the ecs.CreateConfigLogSubscriptions API asynchronously
// api document: https://help.aliyun.com/api/ecs/createconfiglogsubscriptions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateConfigLogSubscriptionsWithChan(request *CreateConfigLogSubscriptionsRequest) (<-chan *CreateConfigLogSubscriptionsResponse, <-chan error) {
	responseChan := make(chan *CreateConfigLogSubscriptionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateConfigLogSubscriptions(request)
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

// CreateConfigLogSubscriptionsWithCallback invokes the ecs.CreateConfigLogSubscriptions API asynchronously
// api document: https://help.aliyun.com/api/ecs/createconfiglogsubscriptions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateConfigLogSubscriptionsWithCallback(request *CreateConfigLogSubscriptionsRequest, callback func(response *CreateConfigLogSubscriptionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateConfigLogSubscriptionsResponse
		var err error
		defer close(result)
		response, err = client.CreateConfigLogSubscriptions(request)
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

// CreateConfigLogSubscriptionsRequest is the request struct for api CreateConfigLogSubscriptions
type CreateConfigLogSubscriptionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                            `position:"Query" name:"ResourceOwnerId"`
	Subscription         *[]CreateConfigLogSubscriptionsSubscription `position:"Query" name:"Subscription"  type:"Repeated"`
	ResourceOwnerAccount string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                      `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                            `position:"Query" name:"OwnerId"`
}

// CreateConfigLogSubscriptionsSubscription is a repeated param struct in CreateConfigLogSubscriptionsRequest
type CreateConfigLogSubscriptionsSubscription struct {
	Name         string `name:"Name"`
	ResourceType string `name:"ResourceType"`
	MnsQueueArn  string `name:"MnsQueueArn"`
}

// CreateConfigLogSubscriptionsResponse is the response struct for api CreateConfigLogSubscriptions
type CreateConfigLogSubscriptionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateConfigLogSubscriptionsRequest creates a request to invoke CreateConfigLogSubscriptions API
func CreateCreateConfigLogSubscriptionsRequest() (request *CreateConfigLogSubscriptionsRequest) {
	request = &CreateConfigLogSubscriptionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "CreateConfigLogSubscriptions", "ecs", "openAPI")
	return
}

// CreateCreateConfigLogSubscriptionsResponse creates a response to parse from CreateConfigLogSubscriptions response
func CreateCreateConfigLogSubscriptionsResponse() (response *CreateConfigLogSubscriptionsResponse) {
	response = &CreateConfigLogSubscriptionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
