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

// ModifyReplicaRecoverMode invokes the r_kvstore.ModifyReplicaRecoverMode API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyreplicarecovermode.html
func (client *Client) ModifyReplicaRecoverMode(request *ModifyReplicaRecoverModeRequest) (response *ModifyReplicaRecoverModeResponse, err error) {
	response = CreateModifyReplicaRecoverModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReplicaRecoverModeWithChan invokes the r_kvstore.ModifyReplicaRecoverMode API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyreplicarecovermode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaRecoverModeWithChan(request *ModifyReplicaRecoverModeRequest) (<-chan *ModifyReplicaRecoverModeResponse, <-chan error) {
	responseChan := make(chan *ModifyReplicaRecoverModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReplicaRecoverMode(request)
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

// ModifyReplicaRecoverModeWithCallback invokes the r_kvstore.ModifyReplicaRecoverMode API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyreplicarecovermode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaRecoverModeWithCallback(request *ModifyReplicaRecoverModeRequest, callback func(response *ModifyReplicaRecoverModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReplicaRecoverModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyReplicaRecoverMode(request)
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

// ModifyReplicaRecoverModeRequest is the request struct for api ModifyReplicaRecoverMode
type ModifyReplicaRecoverModeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	RecoverMode          string           `position:"Query" name:"RecoverMode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyReplicaRecoverModeResponse is the response struct for api ModifyReplicaRecoverMode
type ModifyReplicaRecoverModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReplicaRecoverModeRequest creates a request to invoke ModifyReplicaRecoverMode API
func CreateModifyReplicaRecoverModeRequest() (request *ModifyReplicaRecoverModeRequest) {
	request = &ModifyReplicaRecoverModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyReplicaRecoverMode", "", "")
	return
}

// CreateModifyReplicaRecoverModeResponse creates a response to parse from ModifyReplicaRecoverMode response
func CreateModifyReplicaRecoverModeResponse() (response *ModifyReplicaRecoverModeResponse) {
	response = &ModifyReplicaRecoverModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
