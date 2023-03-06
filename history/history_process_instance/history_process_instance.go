package history_process_instance

import (
	"encoding/json"
	"time"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/identity"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
	"github.com/topology-zero/flowablesdk/variable"
)

type HistoryProcessInstance struct {
	Id                     string              `json:"id"`
	BusinessKey            string              `json:"businessKey"`
	ProcessDefinitionId    string              `json:"processDefinitionId"`
	ProcessDefinitionUrl   string              `json:"processDefinitionUrl"`
	StartTime              *time.Time          `json:"startTime"`
	EndTime                *time.Time          `json:"endTime"`
	DurationInMillis       int                 `json:"durationInMillis"`
	StartUserId            string              `json:"startUserId"`
	StartActivityId        string              `json:"startActivityId"`
	EndActivityId          string              `json:"endActivityId"`
	DeleteReason           string              `json:"deleteReason"`
	SuperProcessInstanceId string              `json:"superProcessInstanceId"`
	Url                    string              `json:"url"`
	Variables              []variable.Variable `json:"variable"`
	TenantId               string              `json:"tenantId"`
}

// List 获取所有流程实例历史
func List(req ListRequest) (list []HistoryProcessInstance, count int, err error) {
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

// Detail 获取单个流程实例历史
func Detail(processInstanceId string) (resp HistoryProcessInstance, err error) {
	request := flowablesdk.GetRequest(DetailApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个流程实例历史
func Delete(processInstanceId string) error {
	request := flowablesdk.GetRequest(DeleteApi, processInstanceId)
	_, err := request.DoHttpRequest()
	return err
}

// ListIdentity 单个流程实例历史相关用户
func ListIdentity(processInstanceId string) (resp []identity.Identity, err error) {
	request := flowablesdk.GetRequest(ListIdentityApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// BinaryData 获取单个流程实例历史二进制文件
func BinaryData(taskId, variableName string) (data []byte, err error) {
	request := flowablesdk.GetRequest(BinaryDataApi, taskId, variableName)
	data, err = request.DoHttpRequest()
	return
}
