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

// DescribeServiceOfCloudDBA invokes the rds.DescribeServiceOfCloudDBA API synchronously
// api document: https://help.aliyun.com/api/rds/describeserviceofclouddba.html
func (client *Client) DescribeServiceOfCloudDBA(request *DescribeServiceOfCloudDBARequest) (response *DescribeServiceOfCloudDBAResponse, err error) {
	response = CreateDescribeServiceOfCloudDBAResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceOfCloudDBAWithChan invokes the rds.DescribeServiceOfCloudDBA API asynchronously
// api document: https://help.aliyun.com/api/rds/describeserviceofclouddba.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServiceOfCloudDBAWithChan(request *DescribeServiceOfCloudDBARequest) (<-chan *DescribeServiceOfCloudDBAResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceOfCloudDBAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServiceOfCloudDBA(request)
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

// DescribeServiceOfCloudDBAWithCallback invokes the rds.DescribeServiceOfCloudDBA API asynchronously
// api document: https://help.aliyun.com/api/rds/describeserviceofclouddba.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServiceOfCloudDBAWithCallback(request *DescribeServiceOfCloudDBARequest, callback func(response *DescribeServiceOfCloudDBAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceOfCloudDBAResponse
		var err error
		defer close(result)
		response, err = client.DescribeServiceOfCloudDBA(request)
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

// DescribeServiceOfCloudDBARequest is the request struct for api DescribeServiceOfCloudDBA
type DescribeServiceOfCloudDBARequest struct {
	*requests.RpcRequest
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

// DescribeServiceOfCloudDBAResponse is the response struct for api DescribeServiceOfCloudDBA
type DescribeServiceOfCloudDBAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ListData  string `json:"ListData" xml:"ListData"`
	AttrData  string `json:"AttrData" xml:"AttrData"`
}

// CreateDescribeServiceOfCloudDBARequest creates a request to invoke DescribeServiceOfCloudDBA API
func CreateDescribeServiceOfCloudDBARequest() (request *DescribeServiceOfCloudDBARequest) {
	request = &DescribeServiceOfCloudDBARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeServiceOfCloudDBA", "", "")
	return
}

// CreateDescribeServiceOfCloudDBAResponse creates a response to parse from DescribeServiceOfCloudDBA response
func CreateDescribeServiceOfCloudDBAResponse() (response *DescribeServiceOfCloudDBAResponse) {
	response = &DescribeServiceOfCloudDBAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
