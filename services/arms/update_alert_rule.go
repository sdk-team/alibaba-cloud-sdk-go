package arms

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

// UpdateAlertRule invokes the arms.UpdateAlertRule API synchronously
// api document: https://help.aliyun.com/api/arms/updatealertrule.html
func (client *Client) UpdateAlertRule(request *UpdateAlertRuleRequest) (response *UpdateAlertRuleResponse, err error) {
	response = CreateUpdateAlertRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAlertRuleWithChan invokes the arms.UpdateAlertRule API asynchronously
// api document: https://help.aliyun.com/api/arms/updatealertrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAlertRuleWithChan(request *UpdateAlertRuleRequest) (<-chan *UpdateAlertRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateAlertRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAlertRule(request)
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

// UpdateAlertRuleWithCallback invokes the arms.UpdateAlertRule API asynchronously
// api document: https://help.aliyun.com/api/arms/updatealertrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAlertRuleWithCallback(request *UpdateAlertRuleRequest, callback func(response *UpdateAlertRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAlertRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateAlertRule(request)
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

// UpdateAlertRuleRequest is the request struct for api UpdateAlertRule
type UpdateAlertRuleRequest struct {
	*requests.RpcRequest
	IsAutoStart         requests.Boolean `position:"Query" name:"IsAutoStart"`
	ProxyUserId         string           `position:"Query" name:"ProxyUserId"`
	ContactGroupIds     string           `position:"Query" name:"ContactGroupIds"`
	AlertId             requests.Integer `position:"Query" name:"AlertId"`
	TemplageAlertConfig string           `position:"Query" name:"TemplageAlertConfig"`
}

// UpdateAlertRuleResponse is the response struct for api UpdateAlertRule
type UpdateAlertRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	AlertId   int64  `json:"AlertId" xml:"AlertId"`
}

// CreateUpdateAlertRuleRequest creates a request to invoke UpdateAlertRule API
func CreateUpdateAlertRuleRequest() (request *UpdateAlertRuleRequest) {
	request = &UpdateAlertRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateAlertRule", "", "")
	return
}

// CreateUpdateAlertRuleResponse creates a response to parse from UpdateAlertRule response
func CreateUpdateAlertRuleResponse() (response *UpdateAlertRuleResponse) {
	response = &UpdateAlertRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
