package acs

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

// ModifyApiNameInDaily invokes the acs.ModifyApiNameInDaily API synchronously
// api document: https://help.aliyun.com/api/acs/modifyapinameindaily.html
func (client *Client) ModifyApiNameInDaily(request *ModifyApiNameInDailyRequest) (response *ModifyApiNameInDailyResponse, err error) {
	response = CreateModifyApiNameInDailyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyApiNameInDailyWithChan invokes the acs.ModifyApiNameInDaily API asynchronously
// api document: https://help.aliyun.com/api/acs/modifyapinameindaily.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyApiNameInDailyWithChan(request *ModifyApiNameInDailyRequest) (<-chan *ModifyApiNameInDailyResponse, <-chan error) {
	responseChan := make(chan *ModifyApiNameInDailyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApiNameInDaily(request)
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

// ModifyApiNameInDailyWithCallback invokes the acs.ModifyApiNameInDaily API asynchronously
// api document: https://help.aliyun.com/api/acs/modifyapinameindaily.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyApiNameInDailyWithCallback(request *ModifyApiNameInDailyRequest, callback func(response *ModifyApiNameInDailyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiNameInDailyResponse
		var err error
		defer close(result)
		response, err = client.ModifyApiNameInDaily(request)
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

// ModifyApiNameInDailyRequest is the request struct for api ModifyApiNameInDaily
type ModifyApiNameInDailyRequest struct {
	*requests.RoaRequest
	ModifyName  string `position:"Query" name:"ModifyName"`
	Name        string `position:"Query" name:"Name"`
	ProductName string `position:"Query" name:"ProductName"`
	ChangeId    string `position:"Query" name:"ChangeId"`
	VersionName string `position:"Query" name:"VersionName"`
}

// ModifyApiNameInDailyResponse is the response struct for api ModifyApiNameInDaily
type ModifyApiNameInDailyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyApiNameInDailyRequest creates a request to invoke ModifyApiNameInDaily API
func CreateModifyApiNameInDailyRequest() (request *ModifyApiNameInDailyRequest) {
	request = &ModifyApiNameInDailyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Acs", "2015-01-01", "ModifyApiNameInDaily", "/modifyApiNameInDaily", "12334", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyApiNameInDailyResponse creates a response to parse from ModifyApiNameInDaily response
func CreateModifyApiNameInDailyResponse() (response *ModifyApiNameInDailyResponse) {
	response = &ModifyApiNameInDailyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
