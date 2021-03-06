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

// ModifyHASwitchConfig invokes the rds.ModifyHASwitchConfig API synchronously
// api document: https://help.aliyun.com/api/rds/modifyhaswitchconfig.html
func (client *Client) ModifyHASwitchConfig(request *ModifyHASwitchConfigRequest) (response *ModifyHASwitchConfigResponse, err error) {
	response = CreateModifyHASwitchConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyHASwitchConfigWithChan invokes the rds.ModifyHASwitchConfig API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyhaswitchconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHASwitchConfigWithChan(request *ModifyHASwitchConfigRequest) (<-chan *ModifyHASwitchConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyHASwitchConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHASwitchConfig(request)
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

// ModifyHASwitchConfigWithCallback invokes the rds.ModifyHASwitchConfig API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyhaswitchconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHASwitchConfigWithCallback(request *ModifyHASwitchConfigRequest, callback func(response *ModifyHASwitchConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHASwitchConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyHASwitchConfig(request)
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

// ModifyHASwitchConfigRequest is the request struct for api ModifyHASwitchConfig
type ModifyHASwitchConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ManualHATime         string           `position:"Query" name:"ManualHATime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	HAConfig             string           `position:"Query" name:"HAConfig"`
}

// ModifyHASwitchConfigResponse is the response struct for api ModifyHASwitchConfig
type ModifyHASwitchConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyHASwitchConfigRequest creates a request to invoke ModifyHASwitchConfig API
func CreateModifyHASwitchConfigRequest() (request *ModifyHASwitchConfigRequest) {
	request = &ModifyHASwitchConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyHASwitchConfig", "", "")
	return
}

// CreateModifyHASwitchConfigResponse creates a response to parse from ModifyHASwitchConfig response
func CreateModifyHASwitchConfigResponse() (response *ModifyHASwitchConfigResponse) {
	response = &ModifyHASwitchConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
