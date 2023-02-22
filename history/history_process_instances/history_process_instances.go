package history_process_instances

import (
	"encoding/json"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type HistoryProcessInstances struct {
	Id                     string               `json:"id"`
	BusinessKey            string               `json:"businessKey"`
	ProcessDefinitionId    string               `json:"processDefinitionId"`
	ProcessDefinitionUrl   string               `json:"processDefinitionUrl"`
	StartTime              *time.Time           `json:"startTime"`
	EndTime                *time.Time           `json:"endTime"`
	DurationInMillis       int                  `json:"durationInMillis"`
	StartUserId            string               `json:"startUserId"`
	StartActivityId        string               `json:"startActivityId"`
	EndActivityId          string               `json:"endActivityId"`
	DeleteReason           string               `json:"deleteReason"`
	SuperProcessInstanceId string               `json:"superProcessInstanceId"`
	Url                    string               `json:"url"`
	Variables              []variables.Variable `json:"variables"`
	TenantId               string               `json:"tenantId"`
}

// List 获取所有流程实例历史
func List(req ListRequest) (list []HistoryProcessInstances, count int, err error) {
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
func Detail(processInstanceId string) (resp HistoryProcessInstances, err error) {
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

// ListIdentity 删除单个流程实例历史相关用户
func ListIdentity(processInstanceId string) (resp []IdentityResponse, err error) {
	request := flowablesdk.GetRequest(ListIdentityApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
