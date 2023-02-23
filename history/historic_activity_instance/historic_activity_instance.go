package historic_activity_instance

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type HistoricActivityInstance struct {
	Id                      string `json:"id"`
	ActivityId              string `json:"activityId"`
	ActivityName            string `json:"activityName"`
	ActivityType            string `json:"activityType"`
	ProcessDefinitionId     string `json:"processDefinitionId"`
	ProcessDefinitionUrl    string `json:"processDefinitionUrl"`
	ProcessInstanceId       string `json:"processInstanceId"`
	ProcessInstanceUrl      string `json:"processInstanceUrl"`
	ExecutionId             string `json:"executionId"`
	TaskId                  string `json:"taskId"`
	CalledProcessInstanceId string `json:"calledProcessInstanceId"`
	Assignee                string `json:"assignee"`
	StartTime               string `json:"startTime"`
	EndTime                 string `json:"endTime"`
	DurationInMillis        int    `json:"durationInMillis"`
	TenantId                string `json:"tenantId"`
}

// List 获取流程历史活动
func List(req ListRequest) (list []HistoricActivityInstance, count int, err error) {
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
