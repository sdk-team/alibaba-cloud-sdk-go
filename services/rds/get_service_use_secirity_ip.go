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

// GetServiceUseSecirityIP invokes the rds.GetServiceUseSecirityIP API synchronously
// api document: https://help.aliyun.com/api/rds/getserviceusesecirityip.html
func (client *Client) GetServiceUseSecirityIP(request *GetServiceUseSecirityIPRequest) (response *GetServiceUseSecirityIPResponse, err error) {
	response = CreateGetServiceUseSecirityIPResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceUseSecirityIPWithChan invokes the rds.GetServiceUseSecirityIP API asynchronously
// api document: https://help.aliyun.com/api/rds/getserviceusesecirityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceUseSecirityIPWithChan(request *GetServiceUseSecirityIPRequest) (<-chan *GetServiceUseSecirityIPResponse, <-chan error) {
	responseChan := make(chan *GetServiceUseSecirityIPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceUseSecirityIP(request)
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

// GetServiceUseSecirityIPWithCallback invokes the rds.GetServiceUseSecirityIP API asynchronously
// api document: https://help.aliyun.com/api/rds/getserviceusesecirityip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceUseSecirityIPWithCallback(request *GetServiceUseSecirityIPRequest, callback func(response *GetServiceUseSecirityIPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceUseSecirityIPResponse
		var err error
		defer close(result)
		response, err = client.GetServiceUseSecirityIP(request)
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

// GetServiceUseSecirityIPRequest is the request struct for api GetServiceUseSecirityIP
type GetServiceUseSecirityIPRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Id                   string           `position:"Query" name:"Id"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GetServiceUseSecirityIPResponse is the response struct for api GetServiceUseSecirityIP
type GetServiceUseSecirityIPResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Service   Service `json:"Service" xml:"Service"`
}

// CreateGetServiceUseSecirityIPRequest creates a request to invoke GetServiceUseSecirityIP API
func CreateGetServiceUseSecirityIPRequest() (request *GetServiceUseSecirityIPRequest) {
	request = &GetServiceUseSecirityIPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "GetServiceUseSecirityIP", "", "")
	return
}

// CreateGetServiceUseSecirityIPResponse creates a response to parse from GetServiceUseSecirityIP response
func CreateGetServiceUseSecirityIPResponse() (response *GetServiceUseSecirityIPResponse) {
	response = &GetServiceUseSecirityIPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
