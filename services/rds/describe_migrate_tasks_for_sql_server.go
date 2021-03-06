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

// DescribeMigrateTasksForSQLServer invokes the rds.DescribeMigrateTasksForSQLServer API synchronously
// api document: https://help.aliyun.com/api/rds/describemigratetasksforsqlserver.html
func (client *Client) DescribeMigrateTasksForSQLServer(request *DescribeMigrateTasksForSQLServerRequest) (response *DescribeMigrateTasksForSQLServerResponse, err error) {
	response = CreateDescribeMigrateTasksForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMigrateTasksForSQLServerWithChan invokes the rds.DescribeMigrateTasksForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/describemigratetasksforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMigrateTasksForSQLServerWithChan(request *DescribeMigrateTasksForSQLServerRequest) (<-chan *DescribeMigrateTasksForSQLServerResponse, <-chan error) {
	responseChan := make(chan *DescribeMigrateTasksForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMigrateTasksForSQLServer(request)
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

// DescribeMigrateTasksForSQLServerWithCallback invokes the rds.DescribeMigrateTasksForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/describemigratetasksforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMigrateTasksForSQLServerWithCallback(request *DescribeMigrateTasksForSQLServerRequest, callback func(response *DescribeMigrateTasksForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMigrateTasksForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.DescribeMigrateTasksForSQLServer(request)
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

// DescribeMigrateTasksForSQLServerRequest is the request struct for api DescribeMigrateTasksForSQLServer
type DescribeMigrateTasksForSQLServerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	EndTime              string           `position:"Query" name:"EndTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeMigrateTasksForSQLServerResponse is the response struct for api DescribeMigrateTasksForSQLServer
type DescribeMigrateTasksForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId        string                                  `json:"RequestId" xml:"RequestId"`
	DBInstanceID     string                                  `json:"DBInstanceID" xml:"DBInstanceID"`
	DBInstanceName   string                                  `json:"DBInstanceName" xml:"DBInstanceName"`
	StartTime        string                                  `json:"StartTime" xml:"StartTime"`
	EndTime          string                                  `json:"EndTime" xml:"EndTime"`
	TotalRecordCount int                                     `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                                     `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                                     `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeMigrateTasksForSQLServer `json:"Items" xml:"Items"`
}

// CreateDescribeMigrateTasksForSQLServerRequest creates a request to invoke DescribeMigrateTasksForSQLServer API
func CreateDescribeMigrateTasksForSQLServerRequest() (request *DescribeMigrateTasksForSQLServerRequest) {
	request = &DescribeMigrateTasksForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeMigrateTasksForSQLServer", "", "")
	return
}

// CreateDescribeMigrateTasksForSQLServerResponse creates a response to parse from DescribeMigrateTasksForSQLServer response
func CreateDescribeMigrateTasksForSQLServerResponse() (response *DescribeMigrateTasksForSQLServerResponse) {
	response = &DescribeMigrateTasksForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
