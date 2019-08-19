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

// InnerOpsDescribeAccessControlListAttribute invokes the slb.InnerOpsDescribeAccessControlListAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/inneropsdescribeaccesscontrollistattribute.html
func (client *Client) InnerOpsDescribeAccessControlListAttribute(request *InnerOpsDescribeAccessControlListAttributeRequest) (response *InnerOpsDescribeAccessControlListAttributeResponse, err error) {
	response = CreateInnerOpsDescribeAccessControlListAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// InnerOpsDescribeAccessControlListAttributeWithChan invokes the slb.InnerOpsDescribeAccessControlListAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/inneropsdescribeaccesscontrollistattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InnerOpsDescribeAccessControlListAttributeWithChan(request *InnerOpsDescribeAccessControlListAttributeRequest) (<-chan *InnerOpsDescribeAccessControlListAttributeResponse, <-chan error) {
	responseChan := make(chan *InnerOpsDescribeAccessControlListAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InnerOpsDescribeAccessControlListAttribute(request)
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

// InnerOpsDescribeAccessControlListAttributeWithCallback invokes the slb.InnerOpsDescribeAccessControlListAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/inneropsdescribeaccesscontrollistattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InnerOpsDescribeAccessControlListAttributeWithCallback(request *InnerOpsDescribeAccessControlListAttributeRequest, callback func(response *InnerOpsDescribeAccessControlListAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InnerOpsDescribeAccessControlListAttributeResponse
		var err error
		defer close(result)
		response, err = client.InnerOpsDescribeAccessControlListAttribute(request)
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

// InnerOpsDescribeAccessControlListAttributeRequest is the request struct for api InnerOpsDescribeAccessControlListAttribute
type InnerOpsDescribeAccessControlListAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AclEntryComment      string           `position:"Query" name:"AclEntryComment"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// InnerOpsDescribeAccessControlListAttributeResponse is the response struct for api InnerOpsDescribeAccessControlListAttribute
type InnerOpsDescribeAccessControlListAttributeResponse struct {
	*responses.BaseResponse
	ResourceOwnerId  int64                                                        `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	RequestId        string                                                       `json:"RequestId" xml:"RequestId"`
	AclId            string                                                       `json:"AclId" xml:"AclId"`
	AclName          string                                                       `json:"AclName" xml:"AclName"`
	AclEntrys        AclEntrysInInnerOpsDescribeAccessControlListAttribute        `json:"AclEntrys" xml:"AclEntrys"`
	RelatedListeners RelatedListenersInInnerOpsDescribeAccessControlListAttribute `json:"RelatedListeners" xml:"RelatedListeners"`
}

// CreateInnerOpsDescribeAccessControlListAttributeRequest creates a request to invoke InnerOpsDescribeAccessControlListAttribute API
func CreateInnerOpsDescribeAccessControlListAttributeRequest() (request *InnerOpsDescribeAccessControlListAttributeRequest) {
	request = &InnerOpsDescribeAccessControlListAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "InnerOpsDescribeAccessControlListAttribute", "slb", "openAPI")
	return
}

// CreateInnerOpsDescribeAccessControlListAttributeResponse creates a response to parse from InnerOpsDescribeAccessControlListAttribute response
func CreateInnerOpsDescribeAccessControlListAttributeResponse() (response *InnerOpsDescribeAccessControlListAttributeResponse) {
	response = &InnerOpsDescribeAccessControlListAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
