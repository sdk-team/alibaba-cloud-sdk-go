package slb

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

// DeleteRules invokes the slb.DeleteRules API synchronously
// api document: https://help.aliyun.com/api/slb/deleterules.html
func (client *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
	response = CreateDeleteRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRulesWithChan invokes the slb.DeleteRules API asynchronously
// api document: https://help.aliyun.com/api/slb/deleterules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRulesWithChan(request *DeleteRulesRequest) (<-chan *DeleteRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRules(request)
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

// DeleteRulesWithCallback invokes the slb.DeleteRules API asynchronously
// api document: https://help.aliyun.com/api/slb/deleterules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRulesWithCallback(request *DeleteRulesRequest, callback func(response *DeleteRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteRules(request)
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

// DeleteRulesRequest is the request struct for api DeleteRules
type DeleteRulesRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RuleIds              string           `position:"Query" name:"RuleIds"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteRulesResponse is the response struct for api DeleteRules
type DeleteRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRulesRequest creates a request to invoke DeleteRules API
func CreateDeleteRulesRequest() (request *DeleteRulesRequest) {
	request = &DeleteRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteRules", "slb", "openAPI")
	return
}

// CreateDeleteRulesResponse creates a response to parse from DeleteRules response
func CreateDeleteRulesResponse() (response *DeleteRulesResponse) {
	response = &DeleteRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
