package ecs

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

// DescribeNetworkInterfaceMonitorData invokes the ecs.DescribeNetworkInterfaceMonitorData API synchronously
func (client *Client) DescribeNetworkInterfaceMonitorData(request *DescribeNetworkInterfaceMonitorDataRequest) (response *DescribeNetworkInterfaceMonitorDataResponse, err error) {
	response = CreateDescribeNetworkInterfaceMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkInterfaceMonitorDataWithChan invokes the ecs.DescribeNetworkInterfaceMonitorData API asynchronously
func (client *Client) DescribeNetworkInterfaceMonitorDataWithChan(request *DescribeNetworkInterfaceMonitorDataRequest) (<-chan *DescribeNetworkInterfaceMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkInterfaceMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkInterfaceMonitorData(request)
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

// DescribeNetworkInterfaceMonitorDataWithCallback invokes the ecs.DescribeNetworkInterfaceMonitorData API asynchronously
func (client *Client) DescribeNetworkInterfaceMonitorDataWithCallback(request *DescribeNetworkInterfaceMonitorDataRequest, callback func(response *DescribeNetworkInterfaceMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkInterfaceMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkInterfaceMonitorData(request)
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

// DescribeNetworkInterfaceMonitorDataRequest is the request struct for api DescribeNetworkInterfaceMonitorData
type DescribeNetworkInterfaceMonitorDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	EniId                string           `position:"Query" name:"EniId"`
}

// DescribeNetworkInterfaceMonitorDataResponse is the response struct for api DescribeNetworkInterfaceMonitorData
type DescribeNetworkInterfaceMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId   string                                           `json:"RequestId" xml:"RequestId"`
	TotalCount  int                                              `json:"TotalCount" xml:"TotalCount"`
	MonitorData MonitorDataInDescribeNetworkInterfaceMonitorData `json:"MonitorData" xml:"MonitorData"`
}

// CreateDescribeNetworkInterfaceMonitorDataRequest creates a request to invoke DescribeNetworkInterfaceMonitorData API
func CreateDescribeNetworkInterfaceMonitorDataRequest() (request *DescribeNetworkInterfaceMonitorDataRequest) {
	request = &DescribeNetworkInterfaceMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfaceMonitorData", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkInterfaceMonitorDataResponse creates a response to parse from DescribeNetworkInterfaceMonitorData response
func CreateDescribeNetworkInterfaceMonitorDataResponse() (response *DescribeNetworkInterfaceMonitorDataResponse) {
	response = &DescribeNetworkInterfaceMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
