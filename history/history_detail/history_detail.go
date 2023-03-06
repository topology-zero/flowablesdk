package history_detail

import (
	"encoding/json"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
	"github.com/topology-zero/flowablesdk/variable"
)

type HistoryDetail struct {
	Id                 string            `json:"id"`
	ProcessInstanceId  string            `json:"processInstanceId"`
	ProcessInstanceUrl string            `json:"processInstanceUrl"`
	ExecutionId        string            `json:"executionId"`
	ActivityInstanceId string            `json:"activityInstanceId"`
	TaskId             string            `json:"taskId"`
	TaskUrl            string            `json:"taskUrl"`
	Time               string            `json:"time"`
	DetailType         string            `json:"detailType"`
	Revision           string            `json:"revision"`
	PropertyId         string            `json:"propertyId"`
	PropertyValue      string            `json:"propertyValue"`
	Variable           variable.Variable `json:"variable"`
}

// List 获取流程历史变量
func List(req ListRequest) (list []HistoryDetail, count int, err error) {
	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	var commonData common.ListCommonResponse
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return
	}

	count = commonData.Total
	err = json.Unmarshal(commonData.Data, &list)
	return
}

// BinaryData 获取单个流程历史变量二进制文件
func BinaryData(varInstanceId string) (data []byte, err error) {
	request := flowablesdk.GetRequest(BinaryDataApi, varInstanceId)
	data, err = request.DoHttpRequest()
	return
}
