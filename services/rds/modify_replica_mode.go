package rds

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

// ModifyReplicaMode invokes the rds.ModifyReplicaMode API synchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicamode.html
func (client *Client) ModifyReplicaMode(request *ModifyReplicaModeRequest) (response *ModifyReplicaModeResponse, err error) {
	response = CreateModifyReplicaModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReplicaModeWithChan invokes the rds.ModifyReplicaMode API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicamode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaModeWithChan(request *ModifyReplicaModeRequest) (<-chan *ModifyReplicaModeResponse, <-chan error) {
	responseChan := make(chan *ModifyReplicaModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReplicaMode(request)
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

// ModifyReplicaModeWithCallback invokes the rds.ModifyReplicaMode API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicamode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaModeWithCallback(request *ModifyReplicaModeRequest, callback func(response *ModifyReplicaModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReplicaModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyReplicaMode(request)
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

// ModifyReplicaModeRequest is the request struct for api ModifyReplicaMode
type ModifyReplicaModeRequest struct {
	*requests.RpcRequest
	DomainMode           string           `position:"Query" name:"DomainMode"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PrimaryInstanceId    string           `position:"Query" name:"PrimaryInstanceId"`
	ReplicaMode          string           `position:"Query" name:"ReplicaMode"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyReplicaModeResponse is the response struct for api ModifyReplicaMode
type ModifyReplicaModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReplicaModeRequest creates a request to invoke ModifyReplicaMode API
func CreateModifyReplicaModeRequest() (request *ModifyReplicaModeRequest) {
	request = &ModifyReplicaModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaMode", "rds", "openAPI")
	return
}

// CreateModifyReplicaModeResponse creates a response to parse from ModifyReplicaMode response
func CreateModifyReplicaModeResponse() (response *ModifyReplicaModeResponse) {
	response = &ModifyReplicaModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
