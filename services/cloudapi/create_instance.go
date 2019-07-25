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

// CreateInstance invokes the cloudapi.CreateInstance API synchronously
// api document: https://help.aliyun.com/api/cloudapi/createinstance.html
func (client *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	response = CreateCreateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInstanceWithChan invokes the cloudapi.CreateInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateInstanceWithChan(request *CreateInstanceRequest) (<-chan *CreateInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstance(request)
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

// CreateInstanceWithCallback invokes the cloudapi.CreateInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateInstanceWithCallback(request *CreateInstanceRequest, callback func(response *CreateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateInstance(request)
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

// CreateInstanceRequest is the request struct for api CreateInstance
type CreateInstanceRequest struct {
	*requests.RpcRequest
	Duration     requests.Integer `position:"Query" name:"Duration"`
	InstanceName string           `position:"Query" name:"InstanceName"`
	AutoPay      requests.Boolean `position:"Query" name:"AutoPay"`
	InstanceSpec string           `position:"Query" name:"InstanceSpec"`
	ZoneId       string           `position:"Query" name:"ZoneId"`
	ChargeType   string           `position:"Query" name:"ChargeType"`
	HttpsPolicy  string           `position:"Query" name:"HttpsPolicy"`
	PricingCycle string           `position:"Query" name:"PricingCycle"`
	Token        string           `position:"Query" name:"Token"`
}

// CreateInstanceResponse is the response struct for api CreateInstance
type CreateInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
}

// CreateCreateInstanceRequest creates a request to invoke CreateInstance API
func CreateCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateInstance", "apigateway", "openAPI")
	return
}

// CreateCreateInstanceResponse creates a response to parse from CreateInstance response
func CreateCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
