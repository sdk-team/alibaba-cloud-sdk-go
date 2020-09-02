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

// QueryDomainChangeRecords invokes the ft.QueryDomainChangeRecords API synchronously
// api document: https://help.aliyun.com/api/ft/querydomainchangerecords.html
func (client *Client) QueryDomainChangeRecords(request *QueryDomainChangeRecordsRequest) (response *QueryDomainChangeRecordsResponse, err error) {
	response = CreateQueryDomainChangeRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainChangeRecordsWithChan invokes the ft.QueryDomainChangeRecords API asynchronously
// api document: https://help.aliyun.com/api/ft/querydomainchangerecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainChangeRecordsWithChan(request *QueryDomainChangeRecordsRequest) (<-chan *QueryDomainChangeRecordsResponse, <-chan error) {
	responseChan := make(chan *QueryDomainChangeRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainChangeRecords(request)
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

// QueryDomainChangeRecordsWithCallback invokes the ft.QueryDomainChangeRecords API asynchronously
// api document: https://help.aliyun.com/api/ft/querydomainchangerecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainChangeRecordsWithCallback(request *QueryDomainChangeRecordsRequest, callback func(response *QueryDomainChangeRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainChangeRecordsResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainChangeRecords(request)
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

// QueryDomainChangeRecordsRequest is the request struct for api QueryDomainChangeRecords
type QueryDomainChangeRecordsRequest struct {
	*requests.RpcRequest
	Product     string           `position:"Query" name:"Product"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	BucUid      requests.Integer `position:"Query" name:"BucUid"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	BucName     string           `position:"Query" name:"BucName"`
	BucEmpId    string           `position:"Query" name:"BucEmpId"`
}

// QueryDomainChangeRecordsResponse is the response struct for api QueryDomainChangeRecords
type QueryDomainChangeRecordsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Data      bool   `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	MessageCN string `json:"MessageCN" xml:"MessageCN"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryDomainChangeRecordsRequest creates a request to invoke QueryDomainChangeRecords API
func CreateQueryDomainChangeRecordsRequest() (request *QueryDomainChangeRecordsRequest) {
	request = &QueryDomainChangeRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "QueryDomainChangeRecords", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDomainChangeRecordsResponse creates a response to parse from QueryDomainChangeRecords response
func CreateQueryDomainChangeRecordsResponse() (response *QueryDomainChangeRecordsResponse) {
	response = &QueryDomainChangeRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
