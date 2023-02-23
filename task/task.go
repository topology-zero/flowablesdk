package task

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

type Task struct {
	Id                        string              `json:"id"`
	Url                       string              `json:"url"`
	Owner                     string              `json:"owner"`
	Assignee                  string              `json:"assignee"`
	DelegationState           string              `json:"delegationState"`
	Name                      string              `json:"name"`
	Description               string              `json:"description"`
	CreateTime                *time.Time          `json:"createTime"`
	DueDate                   *time.Time          `json:"dueDate"`
	Priority                  int                 `json:"priority"`
	Suspended                 bool                `json:"suspended"`
	ClaimTime                 *time.Time          `json:"claimTime"`
	TaskDefinitionKey         string              `json:"taskDefinitionKey"`
	ScopeDefinitionId         string              `json:"scopeDefinitionId"`
	ScopeId                   string              `json:"scopeId"`
	SubScopeId                string              `json:"subScopeId"`
	ScopeType                 string              `json:"scopeType"`
	PropagatedStageInstanceId string              `json:"propagatedStageInstanceId"`
	TenantId                  string              `json:"tenantId"`
	Category                  string              `json:"category"`
	FormKey                   string              `json:"formKey"`
	ParentTaskId              string              `json:"parentTaskId"`
	ParentTaskUrl             string              `json:"parentTaskUrl"`
	ExecutionId               string              `json:"executionId"`
	ExecutionUrl              string              `json:"executionUrl"`
	ProcessInstanceId         string              `json:"processInstanceId"`
	ProcessInstanceUrl        string              `json:"processInstanceUrl"`
	ProcessDefinitionId       string              `json:"processDefinitionId"`
	ProcessDefinitionUrl      string              `json:"processDefinitionUrl"`
	Variables                 []variable.Variable `json:"variable"`
}

// List 获取任务列表
func List(req ListRequest) (list []Task, count int, err error) {
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

// Detail 获取任务详情
func Detail(taskId string) (resp Task, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Update 编辑任务
func Update(taskId string, req UpdateRequest) (resp Task, err error) {
	request := flowablesdk.GetRequest(UpdateApi, taskId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Action 对任务进行操作
//
// 完成任务:Complete a task
//{
//  "action" : "complete",
//  "variable" : []
//}
//
// 领取任务:Claim a task
//{
//  "action" : "claim",
//  "assignee" : "userWhoClaims"
//}
//
// 指派任务:Delegate a task
//{
//  "action" : "delegate",
//  "assignee" : "userToDelegateTo"
//}
//
// 解决指派任务:Resolve a task
//{
//  "action" : "resolve"
//}
func Action(taskId string, req ActionRequest) error {
	request := flowablesdk.GetRequest(ActionApi, taskId)
	request.With(httpclient.WithJson(req))
	_, err := request.DoHttpRequest()
	return err
}

// Delete 删除任务 -- 无法删除整个流程的一部分任务
func Delete(taskId string, req DeleteRequest) error {
	request := flowablesdk.GetRequest(DeleteApi, taskId)

	query := map[string]string{}

	if req.CascadeHistory {
		query["cascadeHistory"] = "true"
	} else {
		if len(req.DeleteReason) == 0 {
			return errors.New("deleteReason is empty")
		}
		query["deleteReason"] = req.DeleteReason
	}

	request.With(httpclient.WithQuery(query))
	_, err := request.DoHttpRequest()
	return err
}
