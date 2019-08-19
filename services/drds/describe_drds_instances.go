package drds

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

// DescribeDrdsInstances invokes the drds.DescribeDrdsInstances API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsinstances.html
func (client *Client) DescribeDrdsInstances(request *DescribeDrdsInstancesRequest) (response *DescribeDrdsInstancesResponse, err error) {
	response = CreateDescribeDrdsInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsInstancesWithChan invokes the drds.DescribeDrdsInstances API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsInstancesWithChan(request *DescribeDrdsInstancesRequest) (<-chan *DescribeDrdsInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsInstances(request)
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

// DescribeDrdsInstancesWithCallback invokes the drds.DescribeDrdsInstances API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsInstancesWithCallback(request *DescribeDrdsInstancesRequest, callback func(response *DescribeDrdsInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsInstances(request)
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

// DescribeDrdsInstancesRequest is the request struct for api DescribeDrdsInstances
type DescribeDrdsInstancesRequest struct {
	*requests.RpcRequest
	Expired     requests.Boolean            `position:"Query" name:"Expired"`
	PageSize    requests.Integer            `position:"Query" name:"PageSize"`
	Description string                      `position:"Query" name:"Description"`
	Tag         *[]DescribeDrdsInstancesTag `position:"Query" name:"Tag"  type:"Repeated"`
	Type        string                      `position:"Query" name:"Type"`
	PageNumber  requests.Integer            `position:"Query" name:"PageNumber"`
}

// DescribeDrdsInstancesTag is a repeated param struct in DescribeDrdsInstancesRequest
type DescribeDrdsInstancesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDrdsInstancesResponse is the response struct for api DescribeDrdsInstances
type DescribeDrdsInstancesResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	Total      int       `json:"Total" xml:"Total"`
	Instances  Instances `json:"Instances" xml:"Instances"`
}

// CreateDescribeDrdsInstancesRequest creates a request to invoke DescribeDrdsInstances API
func CreateDescribeDrdsInstancesRequest() (request *DescribeDrdsInstancesRequest) {
	request = &DescribeDrdsInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsInstances", "drds", "openAPI")
	return
}

// CreateDescribeDrdsInstancesResponse creates a response to parse from DescribeDrdsInstances response
func CreateDescribeDrdsInstancesResponse() (response *DescribeDrdsInstancesResponse) {
	response = &DescribeDrdsInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
