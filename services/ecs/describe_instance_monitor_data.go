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

// DescribeInstanceMonitorData invokes the ecs.DescribeInstanceMonitorData API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemonitordata.html
func (client *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
	response = CreateDescribeInstanceMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceMonitorDataWithChan invokes the ecs.DescribeInstanceMonitorData API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemonitordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceMonitorDataWithChan(request *DescribeInstanceMonitorDataRequest) (<-chan *DescribeInstanceMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceMonitorData(request)
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

// DescribeInstanceMonitorDataWithCallback invokes the ecs.DescribeInstanceMonitorData API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemonitordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceMonitorDataWithCallback(request *DescribeInstanceMonitorDataRequest, callback func(response *DescribeInstanceMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceMonitorData(request)
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

// DescribeInstanceMonitorDataRequest is the request struct for api DescribeInstanceMonitorData
type DescribeInstanceMonitorDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeInstanceMonitorDataResponse is the response struct for api DescribeInstanceMonitorData
type DescribeInstanceMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId   string                                   `json:"RequestId" xml:"RequestId"`
	MonitorData MonitorDataInDescribeInstanceMonitorData `json:"MonitorData" xml:"MonitorData"`
}

// CreateDescribeInstanceMonitorDataRequest creates a request to invoke DescribeInstanceMonitorData API
func CreateDescribeInstanceMonitorDataRequest() (request *DescribeInstanceMonitorDataRequest) {
	request = &DescribeInstanceMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceMonitorData", "", "")
	return
}

// CreateDescribeInstanceMonitorDataResponse creates a response to parse from DescribeInstanceMonitorData response
func CreateDescribeInstanceMonitorDataResponse() (response *DescribeInstanceMonitorDataResponse) {
	response = &DescribeInstanceMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
