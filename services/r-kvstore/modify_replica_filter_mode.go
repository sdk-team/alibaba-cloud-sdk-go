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

// ModifyReplicaFilterMode invokes the r_kvstore.ModifyReplicaFilterMode API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyreplicafiltermode.html
func (client *Client) ModifyReplicaFilterMode(request *ModifyReplicaFilterModeRequest) (response *ModifyReplicaFilterModeResponse, err error) {
	response = CreateModifyReplicaFilterModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReplicaFilterModeWithChan invokes the r_kvstore.ModifyReplicaFilterMode API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyreplicafiltermode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaFilterModeWithChan(request *ModifyReplicaFilterModeRequest) (<-chan *ModifyReplicaFilterModeResponse, <-chan error) {
	responseChan := make(chan *ModifyReplicaFilterModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReplicaFilterMode(request)
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

// ModifyReplicaFilterModeWithCallback invokes the r_kvstore.ModifyReplicaFilterMode API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyreplicafiltermode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaFilterModeWithCallback(request *ModifyReplicaFilterModeRequest, callback func(response *ModifyReplicaFilterModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReplicaFilterModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyReplicaFilterMode(request)
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

// ModifyReplicaFilterModeRequest is the request struct for api ModifyReplicaFilterMode
type ModifyReplicaFilterModeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	FillterDDL           string           `position:"Query" name:"FillterDDL"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyReplicaFilterModeResponse is the response struct for api ModifyReplicaFilterMode
type ModifyReplicaFilterModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReplicaFilterModeRequest creates a request to invoke ModifyReplicaFilterMode API
func CreateModifyReplicaFilterModeRequest() (request *ModifyReplicaFilterModeRequest) {
	request = &ModifyReplicaFilterModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyReplicaFilterMode", "", "")
	return
}

// CreateModifyReplicaFilterModeResponse creates a response to parse from ModifyReplicaFilterMode response
func CreateModifyReplicaFilterModeResponse() (response *ModifyReplicaFilterModeResponse) {
	response = &ModifyReplicaFilterModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
