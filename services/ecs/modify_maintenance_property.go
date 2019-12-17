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

// ModifyMaintenanceProperty invokes the ecs.ModifyMaintenanceProperty API synchronously
// api document: https://help.aliyun.com/api/ecs/modifymaintenanceproperty.html
func (client *Client) ModifyMaintenanceProperty(request *ModifyMaintenancePropertyRequest) (response *ModifyMaintenancePropertyResponse, err error) {
	response = CreateModifyMaintenancePropertyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMaintenancePropertyWithChan invokes the ecs.ModifyMaintenanceProperty API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifymaintenanceproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMaintenancePropertyWithChan(request *ModifyMaintenancePropertyRequest) (<-chan *ModifyMaintenancePropertyResponse, <-chan error) {
	responseChan := make(chan *ModifyMaintenancePropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMaintenanceProperty(request)
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

// ModifyMaintenancePropertyWithCallback invokes the ecs.ModifyMaintenanceProperty API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifymaintenanceproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMaintenancePropertyWithCallback(request *ModifyMaintenancePropertyRequest, callback func(response *ModifyMaintenancePropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMaintenancePropertyResponse
		var err error
		defer close(result)
		response, err = client.ModifyMaintenanceProperty(request)
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

// ModifyMaintenancePropertyRequest is the request struct for api ModifyMaintenanceProperty
type ModifyMaintenancePropertyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ActionOnMaintenance  string           `position:"Query" name:"ActionOnMaintenance"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// ModifyMaintenancePropertyResponse is the response struct for api ModifyMaintenanceProperty
type ModifyMaintenancePropertyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMaintenancePropertyRequest creates a request to invoke ModifyMaintenanceProperty API
func CreateModifyMaintenancePropertyRequest() (request *ModifyMaintenancePropertyRequest) {
	request = &ModifyMaintenancePropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyMaintenanceProperty", "", "")
	return
}

// CreateModifyMaintenancePropertyResponse creates a response to parse from ModifyMaintenanceProperty response
func CreateModifyMaintenancePropertyResponse() (response *ModifyMaintenancePropertyResponse) {
	response = &ModifyMaintenancePropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
