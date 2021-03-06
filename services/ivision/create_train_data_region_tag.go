package ivision

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

// CreateTrainDataRegionTag invokes the ivision.CreateTrainDataRegionTag API synchronously
// api document: https://help.aliyun.com/api/ivision/createtraindataregiontag.html
func (client *Client) CreateTrainDataRegionTag(request *CreateTrainDataRegionTagRequest) (response *CreateTrainDataRegionTagResponse, err error) {
	response = CreateCreateTrainDataRegionTagResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrainDataRegionTagWithChan invokes the ivision.CreateTrainDataRegionTag API asynchronously
// api document: https://help.aliyun.com/api/ivision/createtraindataregiontag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrainDataRegionTagWithChan(request *CreateTrainDataRegionTagRequest) (<-chan *CreateTrainDataRegionTagResponse, <-chan error) {
	responseChan := make(chan *CreateTrainDataRegionTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrainDataRegionTag(request)
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

// CreateTrainDataRegionTagWithCallback invokes the ivision.CreateTrainDataRegionTag API asynchronously
// api document: https://help.aliyun.com/api/ivision/createtraindataregiontag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrainDataRegionTagWithCallback(request *CreateTrainDataRegionTagRequest, callback func(response *CreateTrainDataRegionTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrainDataRegionTagResponse
		var err error
		defer close(result)
		response, err = client.CreateTrainDataRegionTag(request)
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

// CreateTrainDataRegionTagRequest is the request struct for api CreateTrainDataRegionTag
type CreateTrainDataRegionTagRequest struct {
	*requests.RpcRequest
	ProjectId string           `position:"Query" name:"ProjectId"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	TagItems  string           `position:"Query" name:"TagItems"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	DataId    string           `position:"Query" name:"DataId"`
}

// CreateTrainDataRegionTagResponse is the response struct for api CreateTrainDataRegionTag
type CreateTrainDataRegionTagResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	TrainData TrainData `json:"TrainData" xml:"TrainData"`
}

// CreateCreateTrainDataRegionTagRequest creates a request to invoke CreateTrainDataRegionTag API
func CreateCreateTrainDataRegionTagRequest() (request *CreateTrainDataRegionTagRequest) {
	request = &CreateTrainDataRegionTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "CreateTrainDataRegionTag", "ivision", "openAPI")
	return
}

// CreateCreateTrainDataRegionTagResponse creates a response to parse from CreateTrainDataRegionTag response
func CreateCreateTrainDataRegionTagResponse() (response *CreateTrainDataRegionTagResponse) {
	response = &CreateTrainDataRegionTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
