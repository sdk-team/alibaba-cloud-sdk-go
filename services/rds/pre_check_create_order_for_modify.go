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

// PreCheckCreateOrderForModify invokes the rds.PreCheckCreateOrderForModify API synchronously
// api document: https://help.aliyun.com/api/rds/precheckcreateorderformodify.html
func (client *Client) PreCheckCreateOrderForModify(request *PreCheckCreateOrderForModifyRequest) (response *PreCheckCreateOrderForModifyResponse, err error) {
	response = CreatePreCheckCreateOrderForModifyResponse()
	err = client.DoAction(request, response)
	return
}

// PreCheckCreateOrderForModifyWithChan invokes the rds.PreCheckCreateOrderForModify API asynchronously
// api document: https://help.aliyun.com/api/rds/precheckcreateorderformodify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreCheckCreateOrderForModifyWithChan(request *PreCheckCreateOrderForModifyRequest) (<-chan *PreCheckCreateOrderForModifyResponse, <-chan error) {
	responseChan := make(chan *PreCheckCreateOrderForModifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreCheckCreateOrderForModify(request)
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

// PreCheckCreateOrderForModifyWithCallback invokes the rds.PreCheckCreateOrderForModify API asynchronously
// api document: https://help.aliyun.com/api/rds/precheckcreateorderformodify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreCheckCreateOrderForModifyWithCallback(request *PreCheckCreateOrderForModifyRequest, callback func(response *PreCheckCreateOrderForModifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreCheckCreateOrderForModifyResponse
		var err error
		defer close(result)
		response, err = client.PreCheckCreateOrderForModify(request)
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

// PreCheckCreateOrderForModifyRequest is the request struct for api PreCheckCreateOrderForModify
type PreCheckCreateOrderForModifyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	NodeType              string           `position:"Query" name:"NodeType"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	EffectiveTime         string           `position:"Query" name:"EffectiveTime"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	SwitchTime            string           `position:"Query" name:"SwitchTime"`
	DBInstanceStorageType string           `position:"Query" name:"DBInstanceStorageType"`
	BusinessInfo          string           `position:"Query" name:"BusinessInfo"`
	AutoPay               requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	Resource              string           `position:"Query" name:"Resource"`
	CommodityCode         string           `position:"Query" name:"CommodityCode"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime              string           `position:"Query" name:"UsedTime"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PromotionCode         string           `position:"Query" name:"PromotionCode"`
	VpcId                 string           `position:"Query" name:"VpcId"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	TimeType              string           `position:"Query" name:"TimeType"`
	PayType               string           `position:"Query" name:"PayType"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

// PreCheckCreateOrderForModifyResponse is the response struct for api PreCheckCreateOrderForModify
type PreCheckCreateOrderForModifyResponse struct {
	*responses.BaseResponse
	RequestId      string                                 `json:"RequestId" xml:"RequestId"`
	PreCheckResult bool                                   `json:"PreCheckResult" xml:"PreCheckResult"`
	Failures       FailuresInPreCheckCreateOrderForModify `json:"Failures" xml:"Failures"`
}

// CreatePreCheckCreateOrderForModifyRequest creates a request to invoke PreCheckCreateOrderForModify API
func CreatePreCheckCreateOrderForModifyRequest() (request *PreCheckCreateOrderForModifyRequest) {
	request = &PreCheckCreateOrderForModifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "PreCheckCreateOrderForModify", "", "")
	return
}

// CreatePreCheckCreateOrderForModifyResponse creates a response to parse from PreCheckCreateOrderForModify response
func CreatePreCheckCreateOrderForModifyResponse() (response *PreCheckCreateOrderForModifyResponse) {
	response = &PreCheckCreateOrderForModifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
