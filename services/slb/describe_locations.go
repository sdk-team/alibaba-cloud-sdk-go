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

// DescribeLocations invokes the slb.DescribeLocations API synchronously
// api document: https://help.aliyun.com/api/slb/describelocations.html
func (client *Client) DescribeLocations(request *DescribeLocationsRequest) (response *DescribeLocationsResponse, err error) {
	response = CreateDescribeLocationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLocationsWithChan invokes the slb.DescribeLocations API asynchronously
// api document: https://help.aliyun.com/api/slb/describelocations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLocationsWithChan(request *DescribeLocationsRequest) (<-chan *DescribeLocationsResponse, <-chan error) {
	responseChan := make(chan *DescribeLocationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLocations(request)
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

// DescribeLocationsWithCallback invokes the slb.DescribeLocations API asynchronously
// api document: https://help.aliyun.com/api/slb/describelocations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLocationsWithCallback(request *DescribeLocationsRequest, callback func(response *DescribeLocationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLocationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLocations(request)
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

// DescribeLocationsRequest is the request struct for api DescribeLocations
type DescribeLocationsRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Namespace            string           `position:"Query" name:"Namespace"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NamespaceUid         string           `position:"Query" name:"NamespaceUid"`
}

// DescribeLocationsResponse is the response struct for api DescribeLocations
type DescribeLocationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeLocationsRequest creates a request to invoke DescribeLocations API
func CreateDescribeLocationsRequest() (request *DescribeLocationsRequest) {
	request = &DescribeLocationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLocations", "slb", "openAPI")
	return
}

// CreateDescribeLocationsResponse creates a response to parse from DescribeLocations response
func CreateDescribeLocationsResponse() (response *DescribeLocationsResponse) {
	response = &DescribeLocationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
