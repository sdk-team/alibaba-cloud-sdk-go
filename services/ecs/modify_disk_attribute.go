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

// ModifyDiskAttribute invokes the ecs.ModifyDiskAttribute API synchronously
// api document: https://help.aliyun.com/api/ecs/modifydiskattribute.html
func (client *Client) ModifyDiskAttribute(request *ModifyDiskAttributeRequest) (response *ModifyDiskAttributeResponse, err error) {
	response = CreateModifyDiskAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDiskAttributeWithChan invokes the ecs.ModifyDiskAttribute API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifydiskattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDiskAttributeWithChan(request *ModifyDiskAttributeRequest) (<-chan *ModifyDiskAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyDiskAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDiskAttribute(request)
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

// ModifyDiskAttributeWithCallback invokes the ecs.ModifyDiskAttribute API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifydiskattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDiskAttributeWithCallback(request *ModifyDiskAttributeRequest, callback func(response *ModifyDiskAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDiskAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDiskAttribute(request)
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

// ModifyDiskAttributeRequest is the request struct for api ModifyDiskAttribute
type ModifyDiskAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	DiskName             string           `position:"Query" name:"DiskName"`
	DeleteAutoSnapshot   requests.Boolean `position:"Query" name:"DeleteAutoSnapshot"`
	DiskIds              *[]string        `position:"Query" name:"DiskIds"  type:"Repeated"`
	DiskId               string           `position:"Query" name:"DiskId"`
	DeleteWithInstance   requests.Boolean `position:"Query" name:"DeleteWithInstance"`
	EnableAutoSnapshot   requests.Boolean `position:"Query" name:"EnableAutoSnapshot"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDiskAttributeResponse is the response struct for api ModifyDiskAttribute
type ModifyDiskAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDiskAttributeRequest creates a request to invoke ModifyDiskAttribute API
func CreateModifyDiskAttributeRequest() (request *ModifyDiskAttributeRequest) {
	request = &ModifyDiskAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDiskAttribute", "", "")
	return
}

// CreateModifyDiskAttributeResponse creates a response to parse from ModifyDiskAttribute response
func CreateModifyDiskAttributeResponse() (response *ModifyDiskAttributeResponse) {
	response = &ModifyDiskAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
