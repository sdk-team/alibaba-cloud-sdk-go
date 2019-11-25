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

// DescribeCouponList invokes the bss.DescribeCouponList API synchronously
// api document: https://help.aliyun.com/api/bss/describecouponlist.html
func (client *Client) DescribeCouponList(request *DescribeCouponListRequest) (response *DescribeCouponListResponse, err error) {
	response = CreateDescribeCouponListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCouponListWithChan invokes the bss.DescribeCouponList API asynchronously
// api document: https://help.aliyun.com/api/bss/describecouponlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCouponListWithChan(request *DescribeCouponListRequest) (<-chan *DescribeCouponListResponse, <-chan error) {
	responseChan := make(chan *DescribeCouponListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCouponList(request)
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

// DescribeCouponListWithCallback invokes the bss.DescribeCouponList API asynchronously
// api document: https://help.aliyun.com/api/bss/describecouponlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCouponListWithCallback(request *DescribeCouponListRequest, callback func(response *DescribeCouponListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCouponListResponse
		var err error
		defer close(result)
		response, err = client.DescribeCouponList(request)
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

// DescribeCouponListRequest is the request struct for api DescribeCouponList
type DescribeCouponListRequest struct {
	*requests.RpcRequest
	StartDeliveryTime string           `position:"Query" name:"StartDeliveryTime"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	EndDeliveryTime   string           `position:"Query" name:"EndDeliveryTime"`
	PageNum           requests.Integer `position:"Query" name:"PageNum"`
	Status            string           `position:"Query" name:"Status"`
}

// DescribeCouponListResponse is the response struct for api DescribeCouponList
type DescribeCouponListResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Coupons   Coupons `json:"Coupons" xml:"Coupons"`
}

// CreateDescribeCouponListRequest creates a request to invoke DescribeCouponList API
func CreateDescribeCouponListRequest() (request *DescribeCouponListRequest) {
	request = &DescribeCouponListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Bss", "2014-07-14", "DescribeCouponList", "bss", "openAPI")
	return
}

// CreateDescribeCouponListResponse creates a response to parse from DescribeCouponList response
func CreateDescribeCouponListResponse() (response *DescribeCouponListResponse) {
	response = &DescribeCouponListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
