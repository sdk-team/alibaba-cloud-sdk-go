package slb

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

// RemoveDomainWafBlackList invokes the slb.RemoveDomainWafBlackList API synchronously
// api document: https://help.aliyun.com/api/slb/removedomainwafblacklist.html
func (client *Client) RemoveDomainWafBlackList(request *RemoveDomainWafBlackListRequest) (response *RemoveDomainWafBlackListResponse, err error) {
	response = CreateRemoveDomainWafBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDomainWafBlackListWithChan invokes the slb.RemoveDomainWafBlackList API asynchronously
// api document: https://help.aliyun.com/api/slb/removedomainwafblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveDomainWafBlackListWithChan(request *RemoveDomainWafBlackListRequest) (<-chan *RemoveDomainWafBlackListResponse, <-chan error) {
	responseChan := make(chan *RemoveDomainWafBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDomainWafBlackList(request)
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

// RemoveDomainWafBlackListWithCallback invokes the slb.RemoveDomainWafBlackList API asynchronously
// api document: https://help.aliyun.com/api/slb/removedomainwafblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveDomainWafBlackListWithCallback(request *RemoveDomainWafBlackListRequest, callback func(response *RemoveDomainWafBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDomainWafBlackListResponse
		var err error
		defer close(result)
		response, err = client.RemoveDomainWafBlackList(request)
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

// RemoveDomainWafBlackListRequest is the request struct for api RemoveDomainWafBlackList
type RemoveDomainWafBlackListRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DomainName           string           `position:"Query" name:"DomainName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	List                 string           `position:"Query" name:"List"`
}

// RemoveDomainWafBlackListResponse is the response struct for api RemoveDomainWafBlackList
type RemoveDomainWafBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveDomainWafBlackListRequest creates a request to invoke RemoveDomainWafBlackList API
func CreateRemoveDomainWafBlackListRequest() (request *RemoveDomainWafBlackListRequest) {
	request = &RemoveDomainWafBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "RemoveDomainWafBlackList", "slb", "openAPI")
	return
}

// CreateRemoveDomainWafBlackListResponse creates a response to parse from RemoveDomainWafBlackList response
func CreateRemoveDomainWafBlackListResponse() (response *RemoveDomainWafBlackListResponse) {
	response = &RemoveDomainWafBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
