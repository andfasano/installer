package cms

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

// DescribeHostAvailabilityList invokes the cms.DescribeHostAvailabilityList API synchronously
func (client *Client) DescribeHostAvailabilityList(request *DescribeHostAvailabilityListRequest) (response *DescribeHostAvailabilityListResponse, err error) {
	response = CreateDescribeHostAvailabilityListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHostAvailabilityListWithChan invokes the cms.DescribeHostAvailabilityList API asynchronously
func (client *Client) DescribeHostAvailabilityListWithChan(request *DescribeHostAvailabilityListRequest) (<-chan *DescribeHostAvailabilityListResponse, <-chan error) {
	responseChan := make(chan *DescribeHostAvailabilityListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHostAvailabilityList(request)
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

// DescribeHostAvailabilityListWithCallback invokes the cms.DescribeHostAvailabilityList API asynchronously
func (client *Client) DescribeHostAvailabilityListWithCallback(request *DescribeHostAvailabilityListRequest, callback func(response *DescribeHostAvailabilityListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHostAvailabilityListResponse
		var err error
		defer close(result)
		response, err = client.DescribeHostAvailabilityList(request)
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

// DescribeHostAvailabilityListRequest is the request struct for api DescribeHostAvailabilityList
type DescribeHostAvailabilityListRequest struct {
	*requests.RpcRequest
	GroupId    requests.Integer `position:"Query" name:"GroupId"`
	TaskName   string           `position:"Query" name:"TaskName"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Id         requests.Integer `position:"Query" name:"Id"`
}

// DescribeHostAvailabilityListResponse is the response struct for api DescribeHostAvailabilityList
type DescribeHostAvailabilityListResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	Success   bool     `json:"Success" xml:"Success"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Total     int      `json:"Total" xml:"Total"`
	TaskList  TaskList `json:"TaskList" xml:"TaskList"`
}

// CreateDescribeHostAvailabilityListRequest creates a request to invoke DescribeHostAvailabilityList API
func CreateDescribeHostAvailabilityListRequest() (request *DescribeHostAvailabilityListRequest) {
	request = &DescribeHostAvailabilityListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeHostAvailabilityList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHostAvailabilityListResponse creates a response to parse from DescribeHostAvailabilityList response
func CreateDescribeHostAvailabilityListResponse() (response *DescribeHostAvailabilityListResponse) {
	response = &DescribeHostAvailabilityListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}