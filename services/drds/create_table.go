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

// CreateTable invokes the drds.CreateTable API synchronously
// api document: https://help.aliyun.com/api/drds/createtable.html
func (client *Client) CreateTable(request *CreateTableRequest) (response *CreateTableResponse, err error) {
	response = CreateCreateTableResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTableWithChan invokes the drds.CreateTable API asynchronously
// api document: https://help.aliyun.com/api/drds/createtable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTableWithChan(request *CreateTableRequest) (<-chan *CreateTableResponse, <-chan error) {
	responseChan := make(chan *CreateTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTable(request)
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

// CreateTableWithCallback invokes the drds.CreateTable API asynchronously
// api document: https://help.aliyun.com/api/drds/createtable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTableWithCallback(request *CreateTableRequest, callback func(response *CreateTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTableResponse
		var err error
		defer close(result)
		response, err = client.CreateTable(request)
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

// CreateTableRequest is the request struct for api CreateTable
type CreateTableRequest struct {
	*requests.RpcRequest
	DbName             string `position:"Query" name:"DbName"`
	DdlSql             string `position:"Query" name:"DdlSql"`
	AllowFullTableScan string `position:"Query" name:"AllowFullTableScan"`
	ShardKey           string `position:"Query" name:"ShardKey"`
	ShardType          string `position:"Query" name:"ShardType"`
	DrdsInstanceId     string `position:"Query" name:"DrdsInstanceId"`
}

// CreateTableResponse is the response struct for api CreateTable
type CreateTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateTableRequest creates a request to invoke CreateTable API
func CreateCreateTableRequest() (request *CreateTableRequest) {
	request = &CreateTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2015-04-13", "CreateTable", "drds", "openAPI")
	return
}

// CreateCreateTableResponse creates a response to parse from CreateTable response
func CreateCreateTableResponse() (response *CreateTableResponse) {
	response = &CreateTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
