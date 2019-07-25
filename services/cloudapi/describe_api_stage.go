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

// DescribeApiStage invokes the cloudapi.DescribeApiStage API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapistage.html
func (client *Client) DescribeApiStage(request *DescribeApiStageRequest) (response *DescribeApiStageResponse, err error) {
	response = CreateDescribeApiStageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiStageWithChan invokes the cloudapi.DescribeApiStage API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapistage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiStageWithChan(request *DescribeApiStageRequest) (<-chan *DescribeApiStageResponse, <-chan error) {
	responseChan := make(chan *DescribeApiStageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiStage(request)
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

// DescribeApiStageWithCallback invokes the cloudapi.DescribeApiStage API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapistage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiStageWithCallback(request *DescribeApiStageRequest, callback func(response *DescribeApiStageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiStageResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiStage(request)
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

// DescribeApiStageRequest is the request struct for api DescribeApiStage
type DescribeApiStageRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	StageId       string `position:"Query" name:"StageId"`
}

// DescribeApiStageResponse is the response struct for api DescribeApiStage
type DescribeApiStageResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	GroupId      string                      `json:"GroupId" xml:"GroupId"`
	StageId      string                      `json:"StageId" xml:"StageId"`
	StageName    string                      `json:"StageName" xml:"StageName"`
	Description  string                      `json:"Description" xml:"Description"`
	CreatedTime  string                      `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime string                      `json:"ModifiedTime" xml:"ModifiedTime"`
	Variables    VariablesInDescribeApiStage `json:"Variables" xml:"Variables"`
}

// CreateDescribeApiStageRequest creates a request to invoke DescribeApiStage API
func CreateDescribeApiStageRequest() (request *DescribeApiStageRequest) {
	request = &DescribeApiStageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiStage", "apigateway", "openAPI")
	return
}

// CreateDescribeApiStageResponse creates a response to parse from DescribeApiStage response
func CreateDescribeApiStageResponse() (response *DescribeApiStageResponse) {
	response = &DescribeApiStageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
