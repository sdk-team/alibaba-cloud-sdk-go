package ecsdemo

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

// CancelInstance invokes the ecsdemo.CancelInstance API synchronously
// api document: https://help.aliyun.com/api/ecsdemo/cancelinstance.html
func (client *Client) CancelInstance(request *CancelInstanceRequest) (response *CancelInstanceResponse, err error) {
	response = CreateCancelInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CancelInstanceWithChan invokes the ecsdemo.CancelInstance API asynchronously
// api document: https://help.aliyun.com/api/ecsdemo/cancelinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelInstanceWithChan(request *CancelInstanceRequest) (<-chan *CancelInstanceResponse, <-chan error) {
	responseChan := make(chan *CancelInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelInstance(request)
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

// CancelInstanceWithCallback invokes the ecsdemo.CancelInstance API asynchronously
// api document: https://help.aliyun.com/api/ecsdemo/cancelinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelInstanceWithCallback(request *CancelInstanceRequest, callback func(response *CancelInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelInstanceResponse
		var err error
		defer close(result)
		response, err = client.CancelInstance(request)
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

// CancelInstanceRequest is the request struct for api CancelInstance
type CancelInstanceRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// CancelInstanceResponse is the response struct for api CancelInstance
type CancelInstanceResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Instance  Instance `json:"Instance" xml:"Instance"`
}

// CreateCancelInstanceRequest creates a request to invoke CancelInstance API
func CreateCancelInstanceRequest() (request *CancelInstanceRequest) {
	request = &CancelInstanceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("EcsDemo", "2019-06-26", "CancelInstance", "/dev/mvp/cancel/instance", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelInstanceResponse creates a response to parse from CancelInstance response
func CreateCancelInstanceResponse() (response *CancelInstanceResponse) {
	response = &CancelInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
