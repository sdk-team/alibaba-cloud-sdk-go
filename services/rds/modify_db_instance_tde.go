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

// ModifyDBInstanceTDE invokes the rds.ModifyDBInstanceTDE API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancetde.html
func (client *Client) ModifyDBInstanceTDE(request *ModifyDBInstanceTDERequest) (response *ModifyDBInstanceTDEResponse, err error) {
	response = CreateModifyDBInstanceTDEResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceTDEWithChan invokes the rds.ModifyDBInstanceTDE API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancetde.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceTDEWithChan(request *ModifyDBInstanceTDERequest) (<-chan *ModifyDBInstanceTDEResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceTDEResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceTDE(request)
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

// ModifyDBInstanceTDEWithCallback invokes the rds.ModifyDBInstanceTDE API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancetde.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceTDEWithCallback(request *ModifyDBInstanceTDERequest, callback func(response *ModifyDBInstanceTDEResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceTDEResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceTDE(request)
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

// ModifyDBInstanceTDERequest is the request struct for api ModifyDBInstanceTDE
type ModifyDBInstanceTDERequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TDEStatus            string           `position:"Query" name:"TDEStatus"`
}

// ModifyDBInstanceTDEResponse is the response struct for api ModifyDBInstanceTDE
type ModifyDBInstanceTDEResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceTDERequest creates a request to invoke ModifyDBInstanceTDE API
func CreateModifyDBInstanceTDERequest() (request *ModifyDBInstanceTDERequest) {
	request = &ModifyDBInstanceTDERequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceTDE", "rds", "openAPI")
	return
}

// CreateModifyDBInstanceTDEResponse creates a response to parse from ModifyDBInstanceTDE response
func CreateModifyDBInstanceTDEResponse() (response *ModifyDBInstanceTDEResponse) {
	response = &ModifyDBInstanceTDEResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
