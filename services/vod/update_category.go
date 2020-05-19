package vod

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

// UpdateCategory invokes the vod.UpdateCategory API synchronously
// api document: https://help.aliyun.com/api/vod/updatecategory.html
func (client *Client) UpdateCategory(request *UpdateCategoryRequest) (response *UpdateCategoryResponse, err error) {
	response = CreateUpdateCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCategoryWithChan invokes the vod.UpdateCategory API asynchronously
// api document: https://help.aliyun.com/api/vod/updatecategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCategoryWithChan(request *UpdateCategoryRequest) (<-chan *UpdateCategoryResponse, <-chan error) {
	responseChan := make(chan *UpdateCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCategory(request)
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

// UpdateCategoryWithCallback invokes the vod.UpdateCategory API asynchronously
// api document: https://help.aliyun.com/api/vod/updatecategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCategoryWithCallback(request *UpdateCategoryRequest, callback func(response *UpdateCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCategoryResponse
		var err error
		defer close(result)
		response, err = client.UpdateCategory(request)
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

// UpdateCategoryRequest is the request struct for api UpdateCategory
type UpdateCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	CateId               requests.Integer `position:"Query" name:"CateId"`
	CateName             string           `position:"Query" name:"CateName"`
}

// UpdateCategoryResponse is the response struct for api UpdateCategory
type UpdateCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCategoryRequest creates a request to invoke UpdateCategory API
func CreateUpdateCategoryRequest() (request *UpdateCategoryRequest) {
	request = &UpdateCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateCategory", "vod", "openAPI")
	return
}

// CreateUpdateCategoryResponse creates a response to parse from UpdateCategory response
func CreateUpdateCategoryResponse() (response *UpdateCategoryResponse) {
	response = &UpdateCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
