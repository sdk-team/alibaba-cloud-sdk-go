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

// DescribeKMSKeyAttribute invokes the ecs.DescribeKMSKeyAttribute API synchronously
// api document: https://help.aliyun.com/api/ecs/describekmskeyattribute.html
func (client *Client) DescribeKMSKeyAttribute(request *DescribeKMSKeyAttributeRequest) (response *DescribeKMSKeyAttributeResponse, err error) {
	response = CreateDescribeKMSKeyAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeKMSKeyAttributeWithChan invokes the ecs.DescribeKMSKeyAttribute API asynchronously
// api document: https://help.aliyun.com/api/ecs/describekmskeyattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKMSKeyAttributeWithChan(request *DescribeKMSKeyAttributeRequest) (<-chan *DescribeKMSKeyAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeKMSKeyAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeKMSKeyAttribute(request)
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

// DescribeKMSKeyAttributeWithCallback invokes the ecs.DescribeKMSKeyAttribute API asynchronously
// api document: https://help.aliyun.com/api/ecs/describekmskeyattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKMSKeyAttributeWithCallback(request *DescribeKMSKeyAttributeRequest, callback func(response *DescribeKMSKeyAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeKMSKeyAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeKMSKeyAttribute(request)
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

// DescribeKMSKeyAttributeRequest is the request struct for api DescribeKMSKeyAttribute
type DescribeKMSKeyAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Channel              string           `position:"Query" name:"channel"`
	Operator             string           `position:"Query" name:"operator"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Token                string           `position:"Query" name:"token"`
	AppKey               string           `position:"Query" name:"appKey"`
	KMSKeyId             string           `position:"Query" name:"KMSKeyId"`
}

// DescribeKMSKeyAttributeResponse is the response struct for api DescribeKMSKeyAttribute
type DescribeKMSKeyAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	KMSKeyId  string `json:"KMSKeyId" xml:"KMSKeyId"`
	Status    string `json:"Status" xml:"Status"`
	EcsTagged bool   `json:"EcsTagged" xml:"EcsTagged"`
	Creator   string `json:"Creator" xml:"Creator"`
	Alias     string `json:"Alias" xml:"Alias"`
}

// CreateDescribeKMSKeyAttributeRequest creates a request to invoke DescribeKMSKeyAttribute API
func CreateDescribeKMSKeyAttributeRequest() (request *DescribeKMSKeyAttributeRequest) {
	request = &DescribeKMSKeyAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "DescribeKMSKeyAttribute", "ecs", "openAPI")
	return
}

// CreateDescribeKMSKeyAttributeResponse creates a response to parse from DescribeKMSKeyAttribute response
func CreateDescribeKMSKeyAttributeResponse() (response *DescribeKMSKeyAttributeResponse) {
	response = &DescribeKMSKeyAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
