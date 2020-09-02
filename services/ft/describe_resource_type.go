package ft

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

// DescribeResourceType invokes the ft.DescribeResourceType API synchronously
// api document: https://help.aliyun.com/api/ft/describeresourcetype.html
func (client *Client) DescribeResourceType(request *DescribeResourceTypeRequest) (response *DescribeResourceTypeResponse, err error) {
	response = CreateDescribeResourceTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceTypeWithChan invokes the ft.DescribeResourceType API asynchronously
// api document: https://help.aliyun.com/api/ft/describeresourcetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceTypeWithChan(request *DescribeResourceTypeRequest) (<-chan *DescribeResourceTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceType(request)
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

// DescribeResourceTypeWithCallback invokes the ft.DescribeResourceType API asynchronously
// api document: https://help.aliyun.com/api/ft/describeresourcetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceTypeWithCallback(request *DescribeResourceTypeRequest, callback func(response *DescribeResourceTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceType(request)
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

// DescribeResourceTypeRequest is the request struct for api DescribeResourceType
type DescribeResourceTypeRequest struct {
	*requests.RpcRequest
	Product        string `position:"Query" name:"Product"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	Env            string `position:"Query" name:"Env"`
	ResourceType   string `position:"Query" name:"ResourceType"`
}

// DescribeResourceTypeResponse is the response struct for api DescribeResourceType
type DescribeResourceTypeResponse struct {
	*responses.BaseResponse
	RequestId                     string                        `json:"RequestId" xml:"RequestId"`
	Product                       string                        `json:"Product" xml:"Product"`
	ResourceTypeCode              string                        `json:"ResourceTypeCode" xml:"ResourceTypeCode"`
	ResourceTypeVersion           string                        `json:"ResourceTypeVersion" xml:"ResourceTypeVersion"`
	Title                         string                        `json:"Title" xml:"Title"`
	UniqueIdentifier              string                        `json:"UniqueIdentifier" xml:"UniqueIdentifier"`
	Description                   string                        `json:"Description" xml:"Description"`
	ResourceIdField               string                        `json:"ResourceIdField" xml:"ResourceIdField"`
	Definitions                   map[string]interface{}        `json:"Definitions" xml:"Definitions"`
	Schema                        Schema                        `json:"Schema" xml:"Schema"`
	OperationAPI                  OperationAPI                  `json:"OperationAPI" xml:"OperationAPI"`
	ResourceTypeAttributeMappings ResourceTypeAttributeMappings `json:"ResourceTypeAttributeMappings" xml:"ResourceTypeAttributeMappings"`
	States                        []State                       `json:"States" xml:"States"`
}

// CreateDescribeResourceTypeRequest creates a request to invoke DescribeResourceType API
func CreateDescribeResourceTypeRequest() (request *DescribeResourceTypeRequest) {
	request = &DescribeResourceTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "DescribeResourceType", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeResourceTypeResponse creates a response to parse from DescribeResourceType response
func CreateDescribeResourceTypeResponse() (response *DescribeResourceTypeResponse) {
	response = &DescribeResourceTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
