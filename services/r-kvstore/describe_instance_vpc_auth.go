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

// DescribeInstanceVPCAuth invokes the r_kvstore.DescribeInstanceVPCAuth API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancevpcauth.html
func (client *Client) DescribeInstanceVPCAuth(request *DescribeInstanceVPCAuthRequest) (response *DescribeInstanceVPCAuthResponse, err error) {
	response = CreateDescribeInstanceVPCAuthResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceVPCAuthWithChan invokes the r_kvstore.DescribeInstanceVPCAuth API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancevpcauth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceVPCAuthWithChan(request *DescribeInstanceVPCAuthRequest) (<-chan *DescribeInstanceVPCAuthResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceVPCAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceVPCAuth(request)
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

// DescribeInstanceVPCAuthWithCallback invokes the r_kvstore.DescribeInstanceVPCAuth API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancevpcauth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceVPCAuthWithCallback(request *DescribeInstanceVPCAuthRequest, callback func(response *DescribeInstanceVPCAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceVPCAuthResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceVPCAuth(request)
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

// DescribeInstanceVPCAuthRequest is the request struct for api DescribeInstanceVPCAuth
type DescribeInstanceVPCAuthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeInstanceVPCAuthResponse is the response struct for api DescribeInstanceVPCAuth
type DescribeInstanceVPCAuthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IsVPCAuth bool   `json:"IsVPCAuth" xml:"IsVPCAuth"`
}

// CreateDescribeInstanceVPCAuthRequest creates a request to invoke DescribeInstanceVPCAuth API
func CreateDescribeInstanceVPCAuthRequest() (request *DescribeInstanceVPCAuthRequest) {
	request = &DescribeInstanceVPCAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceVPCAuth", "", "")
	return
}

// CreateDescribeInstanceVPCAuthResponse creates a response to parse from DescribeInstanceVPCAuth response
func CreateDescribeInstanceVPCAuthResponse() (response *DescribeInstanceVPCAuthResponse) {
	response = &DescribeInstanceVPCAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
