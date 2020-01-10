package cloudwf

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

// ReportZoneMinute invokes the cloudwf.ReportZoneMinute API synchronously
// api document: https://help.aliyun.com/api/cloudwf/reportzoneminute.html
func (client *Client) ReportZoneMinute(request *ReportZoneMinuteRequest) (response *ReportZoneMinuteResponse, err error) {
	response = CreateReportZoneMinuteResponse()
	err = client.DoAction(request, response)
	return
}

// ReportZoneMinuteWithChan invokes the cloudwf.ReportZoneMinute API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/reportzoneminute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportZoneMinuteWithChan(request *ReportZoneMinuteRequest) (<-chan *ReportZoneMinuteResponse, <-chan error) {
	responseChan := make(chan *ReportZoneMinuteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportZoneMinute(request)
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

// ReportZoneMinuteWithCallback invokes the cloudwf.ReportZoneMinute API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/reportzoneminute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportZoneMinuteWithCallback(request *ReportZoneMinuteRequest, callback func(response *ReportZoneMinuteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportZoneMinuteResponse
		var err error
		defer close(result)
		response, err = client.ReportZoneMinute(request)
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

// ReportZoneMinuteRequest is the request struct for api ReportZoneMinute
type ReportZoneMinuteRequest struct {
	*requests.RpcRequest
	BeginDate string           `position:"Query" name:"BeginDate"`
	EndDate   string           `position:"Query" name:"EndDate"`
	Agsid     requests.Integer `position:"Query" name:"Agsid"`
}

// ReportZoneMinuteResponse is the response struct for api ReportZoneMinute
type ReportZoneMinuteResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateReportZoneMinuteRequest creates a request to invoke ReportZoneMinute API
func CreateReportZoneMinuteRequest() (request *ReportZoneMinuteRequest) {
	request = &ReportZoneMinuteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneMinute", "cloudwf", "openAPI")
	return
}

// CreateReportZoneMinuteResponse creates a response to parse from ReportZoneMinute response
func CreateReportZoneMinuteResponse() (response *ReportZoneMinuteResponse) {
	response = &ReportZoneMinuteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}