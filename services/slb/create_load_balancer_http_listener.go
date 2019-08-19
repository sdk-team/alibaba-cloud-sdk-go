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

// CreateLoadBalancerHTTPListener invokes the slb.CreateLoadBalancerHTTPListener API synchronously
// api document: https://help.aliyun.com/api/slb/createloadbalancerhttplistener.html
func (client *Client) CreateLoadBalancerHTTPListener(request *CreateLoadBalancerHTTPListenerRequest) (response *CreateLoadBalancerHTTPListenerResponse, err error) {
	response = CreateCreateLoadBalancerHTTPListenerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerHTTPListenerWithChan invokes the slb.CreateLoadBalancerHTTPListener API asynchronously
// api document: https://help.aliyun.com/api/slb/createloadbalancerhttplistener.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLoadBalancerHTTPListenerWithChan(request *CreateLoadBalancerHTTPListenerRequest) (<-chan *CreateLoadBalancerHTTPListenerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerHTTPListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancerHTTPListener(request)
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

// CreateLoadBalancerHTTPListenerWithCallback invokes the slb.CreateLoadBalancerHTTPListener API asynchronously
// api document: https://help.aliyun.com/api/slb/createloadbalancerhttplistener.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLoadBalancerHTTPListenerWithCallback(request *CreateLoadBalancerHTTPListenerRequest, callback func(response *CreateLoadBalancerHTTPListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerHTTPListenerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancerHTTPListener(request)
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

// CreateLoadBalancerHTTPListenerRequest is the request struct for api CreateLoadBalancerHTTPListener
type CreateLoadBalancerHTTPListenerRequest struct {
	*requests.RpcRequest
	AccessKeyId            string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	ListenerForward        string           `position:"Query" name:"ListenerForward"`
	XForwardedFor          string           `position:"Query" name:"XForwardedFor"`
	HealthCheckURI         string           `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold     requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       requests.Integer `position:"Query" name:"HealthyThreshold"`
	AclStatus              string           `position:"Query" name:"AclStatus"`
	Scheduler              string           `position:"Query" name:"Scheduler"`
	AclType                string           `position:"Query" name:"AclType"`
	HealthCheck            string           `position:"Query" name:"HealthCheck"`
	ForwardPort            string           `position:"Query" name:"ForwardPort"`
	MaxConnection          requests.Integer `position:"Query" name:"MaxConnection"`
	EnableHttp2            string           `position:"Query" name:"EnableHttp2"`
	CookieTimeout          requests.Integer `position:"Query" name:"CookieTimeout"`
	StickySessionType      string           `position:"Query" name:"StickySessionType"`
	VpcIds                 string           `position:"Query" name:"VpcIds"`
	VServerGroupId         string           `position:"Query" name:"VServerGroupId"`
	AclId                  string           `position:"Query" name:"AclId"`
	ListenerPort           requests.Integer `position:"Query" name:"ListenerPort"`
	Cookie                 string           `position:"Query" name:"Cookie"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth              requests.Integer `position:"Query" name:"Bandwidth"`
	StickySession          string           `position:"Query" name:"StickySession"`
	HealthCheckMethod      string           `position:"Query" name:"HealthCheckMethod"`
	HealthCheckDomain      string           `position:"Query" name:"HealthCheckDomain"`
	RequestTimeout         requests.Integer `position:"Query" name:"RequestTimeout"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	Gzip                   string           `position:"Query" name:"Gzip"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Tags                   string           `position:"Query" name:"Tags"`
	IdleTimeout            requests.Integer `position:"Query" name:"IdleTimeout"`
	LoadBalancerId         string           `position:"Query" name:"LoadBalancerId"`
	XForwardedForSLBIP     string           `position:"Query" name:"XForwardedFor_SLBIP"`
	BackendServerPort      requests.Integer `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval    requests.Integer `position:"Query" name:"HealthCheckInterval"`
	XForwardedForProto     string           `position:"Query" name:"XForwardedFor_proto"`
	XForwardedForSLBID     string           `position:"Query" name:"XForwardedFor_SLBID"`
	HealthCheckConnectPort requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string           `position:"Query" name:"HealthCheckHttpCode"`
}

// CreateLoadBalancerHTTPListenerResponse is the response struct for api CreateLoadBalancerHTTPListener
type CreateLoadBalancerHTTPListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLoadBalancerHTTPListenerRequest creates a request to invoke CreateLoadBalancerHTTPListener API
func CreateCreateLoadBalancerHTTPListenerRequest() (request *CreateLoadBalancerHTTPListenerRequest) {
	request = &CreateLoadBalancerHTTPListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerHTTPListener", "slb", "openAPI")
	return
}

// CreateCreateLoadBalancerHTTPListenerResponse creates a response to parse from CreateLoadBalancerHTTPListener response
func CreateCreateLoadBalancerHTTPListenerResponse() (response *CreateLoadBalancerHTTPListenerResponse) {
	response = &CreateLoadBalancerHTTPListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
