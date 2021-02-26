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

// FtIpFlowControl invokes the ft.FtIpFlowControl API synchronously
func (client *Client) FtIpFlowControl(request *FtIpFlowControlRequest) (response *FtIpFlowControlResponse, err error) {
	response = CreateFtIpFlowControlResponse()
	err = client.DoAction(request, response)
	return
}

// FtIpFlowControlWithChan invokes the ft.FtIpFlowControl API asynchronously
func (client *Client) FtIpFlowControlWithChan(request *FtIpFlowControlRequest) (<-chan *FtIpFlowControlResponse, <-chan error) {
	responseChan := make(chan *FtIpFlowControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FtIpFlowControl(request)
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

// FtIpFlowControlWithCallback invokes the ft.FtIpFlowControl API asynchronously
func (client *Client) FtIpFlowControlWithCallback(request *FtIpFlowControlRequest, callback func(response *FtIpFlowControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FtIpFlowControlResponse
		var err error
		defer close(result)
		response, err = client.FtIpFlowControl(request)
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

// FtIpFlowControlRequest is the request struct for api FtIpFlowControl
type FtIpFlowControlRequest struct {
	*requests.RpcRequest
	StringList string `position:"Query" name:"StringList"`
}

// FtIpFlowControlResponse is the response struct for api FtIpFlowControl
type FtIpFlowControlResponse struct {
	*responses.BaseResponse
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Names        NamesInFtIpFlowControl `json:"Names" xml:"Names"`
	Names01      Names01                `json:"Names01" xml:"Names01"`
	IntegerList  IntegerList            `json:"IntegerList" xml:"IntegerList"`
	IntegerList1 IntegerList1           `json:"IntegerList1" xml:"IntegerList1"`
}

// CreateFtIpFlowControlRequest creates a request to invoke FtIpFlowControl API
func CreateFtIpFlowControlRequest() (request *FtIpFlowControlRequest) {
	request = &FtIpFlowControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "FtIpFlowControl", "", "")
	request.Method = requests.POST
	return
}

// CreateFtIpFlowControlResponse creates a response to parse from FtIpFlowControl response
func CreateFtIpFlowControlResponse() (response *FtIpFlowControlResponse) {
	response = &FtIpFlowControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
