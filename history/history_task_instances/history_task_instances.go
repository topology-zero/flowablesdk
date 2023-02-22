package history_task_instances

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type HistoryProcessInstancesTask struct {
	Id                   string                       `json:"id"`
	ProcessDefinitionId  string                       `json:"processDefinitionId"`
	ProcessDefinitionUrl string                       `json:"processDefinitionUrl"`
	ProcessInstanceId    string                       `json:"processInstanceId"`
	ProcessInstanceUrl   string                       `json:"processInstanceUrl"`
	ExecutionId          string                       `json:"executionId"`
	Name                 string                       `json:"name"`
	Description          string                       `json:"description"`
	DeleteReason         string                       `json:"deleteReason"`
	Owner                string                       `json:"owner"`
	Assignee             string                       `json:"assignee"`
	StartTime            string                       `json:"startTime"`
	EndTime              string                       `json:"endTime"`
	DurationInMillis     int                          `json:"durationInMillis"`
	WorkTimeInMillis     int                          `json:"workTimeInMillis"`
	ClaimTime            string                       `json:"claimTime"`
	TaskDefinitionKey    string                       `json:"taskDefinitionKey"`
	FormKey              string                       `json:"formKey"`
	Priority             int                          `json:"priority"`
	DueDate              string                       `json:"dueDate"`
	ParentTaskId         string                       `json:"parentTaskId"`
	Url                  string                       `json:"url"`
	TaskVariables        []variables.VariableResponse `json:"taskVariables"`
	ProcessVariables     []variables.VariableResponse `json:"processVariables"`
	TenantId             string                       `json:"tenantId"`
}

// List 获取流程实例历史任务
func List(req ListRequest) (list []HistoryProcessInstancesTask, count int, err error) {
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

// Detail 获取单个流程实例历史任务
func Detail(taskId string) (resp HistoryProcessInstancesTask, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个流程实例历史任务
func Delete(taskId string) error {
	request := flowablesdk.GetRequest(DeleteApi, taskId)
	_, err := request.DoHttpRequest()
	return err
}
