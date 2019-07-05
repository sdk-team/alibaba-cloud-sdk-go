package rds

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

// DescribeDynamicVerificationList invokes the rds.DescribeDynamicVerificationList API synchronously
// api document: https://help.aliyun.com/api/rds/describedynamicverificationlist.html
func (client *Client) DescribeDynamicVerificationList(request *DescribeDynamicVerificationListRequest) (response *DescribeDynamicVerificationListResponse, err error) {
	response = CreateDescribeDynamicVerificationListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDynamicVerificationListWithChan invokes the rds.DescribeDynamicVerificationList API asynchronously
// api document: https://help.aliyun.com/api/rds/describedynamicverificationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDynamicVerificationListWithChan(request *DescribeDynamicVerificationListRequest) (<-chan *DescribeDynamicVerificationListResponse, <-chan error) {
	responseChan := make(chan *DescribeDynamicVerificationListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDynamicVerificationList(request)
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

// DescribeDynamicVerificationListWithCallback invokes the rds.DescribeDynamicVerificationList API asynchronously
// api document: https://help.aliyun.com/api/rds/describedynamicverificationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDynamicVerificationListWithCallback(request *DescribeDynamicVerificationListRequest, callback func(response *DescribeDynamicVerificationListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDynamicVerificationListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDynamicVerificationList(request)
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

// DescribeDynamicVerificationListRequest is the request struct for api DescribeDynamicVerificationList
type DescribeDynamicVerificationListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeDynamicVerificationListResponse is the response struct for api DescribeDynamicVerificationList
type DescribeDynamicVerificationListResponse struct {
	*responses.BaseResponse
	RequestId        string                                       `json:"RequestId" xml:"RequestId"`
	ReplicaId        string                                       `json:"ReplicaId" xml:"ReplicaId"`
	PagNumber        int                                          `json:"PagNumber" xml:"PagNumber"`
	PageRecordCount  int                                          `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                                          `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            []ItemsItemInDescribeDynamicVerificationList `json:"Items" xml:"Items"`
}

// CreateDescribeDynamicVerificationListRequest creates a request to invoke DescribeDynamicVerificationList API
func CreateDescribeDynamicVerificationListRequest() (request *DescribeDynamicVerificationListRequest) {
	request = &DescribeDynamicVerificationListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDynamicVerificationList", "rds", "openAPI")
	return
}

// CreateDescribeDynamicVerificationListResponse creates a response to parse from DescribeDynamicVerificationList response
func CreateDescribeDynamicVerificationListResponse() (response *DescribeDynamicVerificationListResponse) {
	response = &DescribeDynamicVerificationListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
