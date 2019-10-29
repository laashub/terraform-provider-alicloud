package emr

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

// GetQueueOutputStatisticInfo invokes the emr.GetQueueOutputStatisticInfo API synchronously
// api document: https://help.aliyun.com/api/emr/getqueueoutputstatisticinfo.html
func (client *Client) GetQueueOutputStatisticInfo(request *GetQueueOutputStatisticInfoRequest) (response *GetQueueOutputStatisticInfoResponse, err error) {
	response = CreateGetQueueOutputStatisticInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetQueueOutputStatisticInfoWithChan invokes the emr.GetQueueOutputStatisticInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/getqueueoutputstatisticinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetQueueOutputStatisticInfoWithChan(request *GetQueueOutputStatisticInfoRequest) (<-chan *GetQueueOutputStatisticInfoResponse, <-chan error) {
	responseChan := make(chan *GetQueueOutputStatisticInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQueueOutputStatisticInfo(request)
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

// GetQueueOutputStatisticInfoWithCallback invokes the emr.GetQueueOutputStatisticInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/getqueueoutputstatisticinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetQueueOutputStatisticInfoWithCallback(request *GetQueueOutputStatisticInfoRequest, callback func(response *GetQueueOutputStatisticInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQueueOutputStatisticInfoResponse
		var err error
		defer close(result)
		response, err = client.GetQueueOutputStatisticInfo(request)
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

// GetQueueOutputStatisticInfoRequest is the request struct for api GetQueueOutputStatisticInfo
type GetQueueOutputStatisticInfoRequest struct {
	*requests.RpcRequest
	FromDatetime    string           `position:"Query" name:"FromDatetime"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ToDatetime      string           `position:"Query" name:"ToDatetime"`
}

// GetQueueOutputStatisticInfoResponse is the response struct for api GetQueueOutputStatisticInfo
type GetQueueOutputStatisticInfoResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	QueueOutputList QueueOutputList `json:"QueueOutputList" xml:"QueueOutputList"`
}

// CreateGetQueueOutputStatisticInfoRequest creates a request to invoke GetQueueOutputStatisticInfo API
func CreateGetQueueOutputStatisticInfoRequest() (request *GetQueueOutputStatisticInfoRequest) {
	request = &GetQueueOutputStatisticInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "GetQueueOutputStatisticInfo", "emr", "openAPI")
	return
}

// CreateGetQueueOutputStatisticInfoResponse creates a response to parse from GetQueueOutputStatisticInfo response
func CreateGetQueueOutputStatisticInfoResponse() (response *GetQueueOutputStatisticInfoResponse) {
	response = &GetQueueOutputStatisticInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}