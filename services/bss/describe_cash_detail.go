package bss

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

// DescribeCashDetail invokes the bss.DescribeCashDetail API synchronously
// api document: https://help.aliyun.com/api/bss/describecashdetail.html
func (client *Client) DescribeCashDetail(request *DescribeCashDetailRequest) (response *DescribeCashDetailResponse, err error) {
	response = CreateDescribeCashDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCashDetailWithChan invokes the bss.DescribeCashDetail API asynchronously
// api document: https://help.aliyun.com/api/bss/describecashdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCashDetailWithChan(request *DescribeCashDetailRequest) (<-chan *DescribeCashDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeCashDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCashDetail(request)
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

// DescribeCashDetailWithCallback invokes the bss.DescribeCashDetail API asynchronously
// api document: https://help.aliyun.com/api/bss/describecashdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCashDetailWithCallback(request *DescribeCashDetailRequest, callback func(response *DescribeCashDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCashDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeCashDetail(request)
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

// DescribeCashDetailRequest is the request struct for api DescribeCashDetail
type DescribeCashDetailRequest struct {
	*requests.RpcRequest
}

// DescribeCashDetailResponse is the response struct for api DescribeCashDetail
type DescribeCashDetailResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	BalanceAmount        string `json:"BalanceAmount" xml:"BalanceAmount"`
	AmountOwed           string `json:"AmountOwed" xml:"AmountOwed"`
	EnableThresholdAlert string `json:"EnableThresholdAlert" xml:"EnableThresholdAlert"`
	MiniAlertThreshold   int64  `json:"MiniAlertThreshold" xml:"MiniAlertThreshold"`
	FrozenAmount         string `json:"FrozenAmount" xml:"FrozenAmount"`
	CreditCardAmount     string `json:"CreditCardAmount" xml:"CreditCardAmount"`
	RemmitanceAmount     string `json:"RemmitanceAmount" xml:"RemmitanceAmount"`
	CreditLimit          string `json:"CreditLimit" xml:"CreditLimit"`
	AvailableCredit      string `json:"AvailableCredit" xml:"AvailableCredit"`
}

// CreateDescribeCashDetailRequest creates a request to invoke DescribeCashDetail API
func CreateDescribeCashDetailRequest() (request *DescribeCashDetailRequest) {
	request = &DescribeCashDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Bss", "2014-07-14", "DescribeCashDetail", "bss", "openAPI")
	return
}

// CreateDescribeCashDetailResponse creates a response to parse from DescribeCashDetail response
func CreateDescribeCashDetailResponse() (response *DescribeCashDetailResponse) {
	response = &DescribeCashDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
