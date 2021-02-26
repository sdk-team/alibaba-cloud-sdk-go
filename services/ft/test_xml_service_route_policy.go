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

// TestXmlServiceRoutePolicy invokes the ft.TestXmlServiceRoutePolicy API synchronously
func (client *Client) TestXmlServiceRoutePolicy(request *TestXmlServiceRoutePolicyRequest) (response *TestXmlServiceRoutePolicyResponse, err error) {
	response = CreateTestXmlServiceRoutePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// TestXmlServiceRoutePolicyWithChan invokes the ft.TestXmlServiceRoutePolicy API asynchronously
func (client *Client) TestXmlServiceRoutePolicyWithChan(request *TestXmlServiceRoutePolicyRequest) (<-chan *TestXmlServiceRoutePolicyResponse, <-chan error) {
	responseChan := make(chan *TestXmlServiceRoutePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestXmlServiceRoutePolicy(request)
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

// TestXmlServiceRoutePolicyWithCallback invokes the ft.TestXmlServiceRoutePolicy API asynchronously
func (client *Client) TestXmlServiceRoutePolicyWithCallback(request *TestXmlServiceRoutePolicyRequest, callback func(response *TestXmlServiceRoutePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestXmlServiceRoutePolicyResponse
		var err error
		defer close(result)
		response, err = client.TestXmlServiceRoutePolicy(request)
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

// TestXmlServiceRoutePolicyRequest is the request struct for api TestXmlServiceRoutePolicy
type TestXmlServiceRoutePolicyRequest struct {
	*requests.RpcRequest
	Api string `position:"Query" name:"Api"`
}

// TestXmlServiceRoutePolicyResponse is the response struct for api TestXmlServiceRoutePolicy
type TestXmlServiceRoutePolicyResponse struct {
	*responses.BaseResponse
	Ban []int `json:"Ban" xml:"Ban"`
	Add Add   `json:"Add" xml:"Add"`
	Db  []Sdw `json:"Db" xml:"Db"`
}

// CreateTestXmlServiceRoutePolicyRequest creates a request to invoke TestXmlServiceRoutePolicy API
func CreateTestXmlServiceRoutePolicyRequest() (request *TestXmlServiceRoutePolicyRequest) {
	request = &TestXmlServiceRoutePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "TestXmlServiceRoutePolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateTestXmlServiceRoutePolicyResponse creates a response to parse from TestXmlServiceRoutePolicy response
func CreateTestXmlServiceRoutePolicyResponse() (response *TestXmlServiceRoutePolicyResponse) {
	response = &TestXmlServiceRoutePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
