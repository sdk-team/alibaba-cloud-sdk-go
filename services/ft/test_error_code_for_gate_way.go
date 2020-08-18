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

// TestErrorCodeForGateWay invokes the ft.TestErrorCodeForGateWay API synchronously
// api document: https://help.aliyun.com/api/ft/testerrorcodeforgateway.html
func (client *Client) TestErrorCodeForGateWay(request *TestErrorCodeForGateWayRequest) (response *TestErrorCodeForGateWayResponse, err error) {
	response = CreateTestErrorCodeForGateWayResponse()
	err = client.DoAction(request, response)
	return
}

// TestErrorCodeForGateWayWithChan invokes the ft.TestErrorCodeForGateWay API asynchronously
// api document: https://help.aliyun.com/api/ft/testerrorcodeforgateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TestErrorCodeForGateWayWithChan(request *TestErrorCodeForGateWayRequest) (<-chan *TestErrorCodeForGateWayResponse, <-chan error) {
	responseChan := make(chan *TestErrorCodeForGateWayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestErrorCodeForGateWay(request)
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

// TestErrorCodeForGateWayWithCallback invokes the ft.TestErrorCodeForGateWay API asynchronously
// api document: https://help.aliyun.com/api/ft/testerrorcodeforgateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TestErrorCodeForGateWayWithCallback(request *TestErrorCodeForGateWayRequest, callback func(response *TestErrorCodeForGateWayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestErrorCodeForGateWayResponse
		var err error
		defer close(result)
		response, err = client.TestErrorCodeForGateWay(request)
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

// TestErrorCodeForGateWayRequest is the request struct for api TestErrorCodeForGateWay
type TestErrorCodeForGateWayRequest struct {
	*requests.RpcRequest
	HttpStatusCode string `position:"Query" name:"HttpStatusCode"`
	Code           string `position:"Query" name:"Code"`
	Success        string `position:"Query" name:"Success"`
	Message        string `position:"Query" name:"Message"`
}

// TestErrorCodeForGateWayResponse is the response struct for api TestErrorCodeForGateWay
type TestErrorCodeForGateWayResponse struct {
	*responses.BaseResponse
	Success        string `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateTestErrorCodeForGateWayRequest creates a request to invoke TestErrorCodeForGateWay API
func CreateTestErrorCodeForGateWayRequest() (request *TestErrorCodeForGateWayRequest) {
	request = &TestErrorCodeForGateWayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "TestErrorCodeForGateWay", "", "")
	request.Method = requests.GET
	return
}

// CreateTestErrorCodeForGateWayResponse creates a response to parse from TestErrorCodeForGateWay response
func CreateTestErrorCodeForGateWayResponse() (response *TestErrorCodeForGateWayResponse) {
	response = &TestErrorCodeForGateWayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
