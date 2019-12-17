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

// DescribeInstanceHistoryEvents invokes the ecs.DescribeInstanceHistoryEvents API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancehistoryevents.html
func (client *Client) DescribeInstanceHistoryEvents(request *DescribeInstanceHistoryEventsRequest) (response *DescribeInstanceHistoryEventsResponse, err error) {
	response = CreateDescribeInstanceHistoryEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceHistoryEventsWithChan invokes the ecs.DescribeInstanceHistoryEvents API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancehistoryevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceHistoryEventsWithChan(request *DescribeInstanceHistoryEventsRequest) (<-chan *DescribeInstanceHistoryEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceHistoryEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceHistoryEvents(request)
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

// DescribeInstanceHistoryEventsWithCallback invokes the ecs.DescribeInstanceHistoryEvents API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancehistoryevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceHistoryEventsWithCallback(request *DescribeInstanceHistoryEventsRequest, callback func(response *DescribeInstanceHistoryEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceHistoryEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceHistoryEvents(request)
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

// DescribeInstanceHistoryEventsRequest is the request struct for api DescribeInstanceHistoryEvents
type DescribeInstanceHistoryEventsRequest struct {
	*requests.RpcRequest
	EventId                  *[]string        `position:"Query" name:"EventId"  type:"Repeated"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EventCycleStatus         string           `position:"Query" name:"EventCycleStatus"`
	PageNumber               requests.Integer `position:"Query" name:"PageNumber"`
	PageSize                 requests.Integer `position:"Query" name:"PageSize"`
	InstanceEventCycleStatus *[]string        `position:"Query" name:"InstanceEventCycleStatus"  type:"Repeated"`
	EventPublishTimeEnd      string           `position:"Query" name:"EventPublishTime.End"`
	InstanceEventType        *[]string        `position:"Query" name:"InstanceEventType"  type:"Repeated"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	NotBeforeStart           string           `position:"Query" name:"NotBefore.Start"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	EventPublishTimeStart    string           `position:"Query" name:"EventPublishTime.Start"`
	InstanceId               string           `position:"Query" name:"InstanceId"`
	NotBeforeEnd             string           `position:"Query" name:"NotBefore.End"`
	EventType                string           `position:"Query" name:"EventType"`
}

// DescribeInstanceHistoryEventsResponse is the response struct for api DescribeInstanceHistoryEvents
type DescribeInstanceHistoryEventsResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	PageNumber             int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize               int                    `json:"PageSize" xml:"PageSize"`
	InstanceSystemEventSet InstanceSystemEventSet `json:"InstanceSystemEventSet" xml:"InstanceSystemEventSet"`
}

// CreateDescribeInstanceHistoryEventsRequest creates a request to invoke DescribeInstanceHistoryEvents API
func CreateDescribeInstanceHistoryEventsRequest() (request *DescribeInstanceHistoryEventsRequest) {
	request = &DescribeInstanceHistoryEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceHistoryEvents", "", "")
	return
}

// CreateDescribeInstanceHistoryEventsResponse creates a response to parse from DescribeInstanceHistoryEvents response
func CreateDescribeInstanceHistoryEventsResponse() (response *DescribeInstanceHistoryEventsResponse) {
	response = &DescribeInstanceHistoryEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
