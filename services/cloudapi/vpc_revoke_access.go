package cloudapi

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

// VpcRevokeAccess invokes the cloudapi.VpcRevokeAccess API synchronously
// api document: https://help.aliyun.com/api/cloudapi/vpcrevokeaccess.html
func (client *Client) VpcRevokeAccess(request *VpcRevokeAccessRequest) (response *VpcRevokeAccessResponse, err error) {
	response = CreateVpcRevokeAccessResponse()
	err = client.DoAction(request, response)
	return
}

// VpcRevokeAccessWithChan invokes the cloudapi.VpcRevokeAccess API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/vpcrevokeaccess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VpcRevokeAccessWithChan(request *VpcRevokeAccessRequest) (<-chan *VpcRevokeAccessResponse, <-chan error) {
	responseChan := make(chan *VpcRevokeAccessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VpcRevokeAccess(request)
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

// VpcRevokeAccessWithCallback invokes the cloudapi.VpcRevokeAccess API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/vpcrevokeaccess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VpcRevokeAccessWithCallback(request *VpcRevokeAccessRequest, callback func(response *VpcRevokeAccessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VpcRevokeAccessResponse
		var err error
		defer close(result)
		response, err = client.VpcRevokeAccess(request)
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

// VpcRevokeAccessRequest is the request struct for api VpcRevokeAccess
type VpcRevokeAccessRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	Port       requests.Integer `position:"Query" name:"Port"`
	VpcId      string           `position:"Query" name:"VpcId"`
}

// VpcRevokeAccessResponse is the response struct for api VpcRevokeAccess
type VpcRevokeAccessResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateVpcRevokeAccessRequest creates a request to invoke VpcRevokeAccess API
func CreateVpcRevokeAccessRequest() (request *VpcRevokeAccessRequest) {
	request = &VpcRevokeAccessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "VpcRevokeAccess", "apigateway", "openAPI")
	return
}

// CreateVpcRevokeAccessResponse creates a response to parse from VpcRevokeAccess response
func CreateVpcRevokeAccessResponse() (response *VpcRevokeAccessResponse) {
	response = &VpcRevokeAccessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
