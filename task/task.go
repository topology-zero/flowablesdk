package task

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type Tasks struct {
	Id                        string               `json:"id"`
	Url                       string               `json:"url"`
	Owner                     string               `json:"owner"`
	Assignee                  string               `json:"assignee"`
	DelegationState           string               `json:"delegationState"`
	Name                      string               `json:"name"`
	Description               string               `json:"description"`
	CreateTime                *time.Time           `json:"createTime"`
	DueDate                   *time.Time           `json:"dueDate"`
	Priority                  int                  `json:"priority"`
	Suspended                 bool                 `json:"suspended"`
	ClaimTime                 *time.Time           `json:"claimTime"`
	TaskDefinitionKey         string               `json:"taskDefinitionKey"`
	ScopeDefinitionId         string               `json:"scopeDefinitionId"`
	ScopeId                   string               `json:"scopeId"`
	SubScopeId                string               `json:"subScopeId"`
	ScopeType                 string               `json:"scopeType"`
	PropagatedStageInstanceId string               `json:"propagatedStageInstanceId"`
	TenantId                  string               `json:"tenantId"`
	Category                  string               `json:"category"`
	FormKey                   string               `json:"formKey"`
	ParentTaskId              string               `json:"parentTaskId"`
	ParentTaskUrl             string               `json:"parentTaskUrl"`
	ExecutionId               string               `json:"executionId"`
	ExecutionUrl              string               `json:"executionUrl"`
	ProcessInstanceId         string               `json:"processInstanceId"`
	ProcessInstanceUrl        string               `json:"processInstanceUrl"`
	ProcessDefinitionId       string               `json:"processDefinitionId"`
	ProcessDefinitionUrl      string               `json:"processDefinitionUrl"`
	Variables                 []variables.Variable `json:"variables"`
}

// List 获取任务列表
func (t Tasks) List(req ListRequest) (resp ListResponse, err error) {
	query := make(map[string]string)

	if len(req.Name) > 0 {
		query["nameLike"] = req.Name
	}

	if len(req.Description) > 0 {
		query["description"] = req.Description
	}

	if len(req.Priority) == 2 {
		query["minimumPriority"] = strconv.Itoa(req.Priority[0])
		query["maximumPriority"] = strconv.Itoa(req.Priority[1])
	}

	if len(req.Assignee) > 0 {
		query["assignee"] = req.Assignee
	}

	if len(req.Owner) > 0 {
		query["owner"] = req.Owner
	}

	if len(req.CandidateUser) > 0 {
		query["candidateUser"] = req.CandidateUser
	}

	if len(req.CandidateGroups) > 0 {
		query["candidateGroups"] = strings.Join(req.CandidateGroups, ",")
	}

	if len(req.CandidateOrAssigned) > 0 {
		query["candidateOrAssigned"] = req.CandidateOrAssigned
	}

	if len(req.InvolvedUser) > 0 {
		query["involvedUser"] = req.InvolvedUser
	}

	if req.Unassigned {
		query["unassigned"] = "true"
	}

	if len(req.DelegationState) > 0 {
		query["delegationState"] = req.DelegationState
	}

	if len(req.TaskDefinitionKey) > 0 {
		query["taskDefinitionKey"] = req.TaskDefinitionKey
	}

	if len(req.ProcessInstanceId) > 0 {
		query["processInstanceId"] = req.ProcessInstanceId
	}

	if len(req.ProcessInstanceBusinessKey) > 0 {
		query["processInstanceBusinessKey"] = req.ProcessInstanceBusinessKey
	}

	if len(req.ProcessDefinitionId) > 0 {
		query["processDefinitionId"] = req.ProcessDefinitionId
	}

	if len(req.ProcessDefinitionKey) > 0 {
		query["processDefinitionKey"] = req.ProcessDefinitionKey
	}

	if len(req.ProcessDefinitionName) > 0 {
		query["processDefinitionName"] = req.ProcessDefinitionName
	}

	if len(req.CreateTime) == 2 {
		query["createdAfter"] = req.CreateTime[0].Format(time.RFC3339)
		query["createdBefore"] = req.CreateTime[1].Format(time.RFC3339)
	}

	if req.Active != nil {
		if *req.Active {
			query["active"] = "true"
		} else {
			query["active"] = "false"
		}
	}

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.Sort) > 0 {
		query["sort"] = req.Sort
	}

	if len(req.Order) > 0 {
		query["order"] = req.Order
	}

	if req.Start < 0 {
		req.Start = 0
	}
	query["start"] = strconv.Itoa(req.Start)

	if req.Size < 1 {
		req.Size = 10
	}
	query["size"] = strconv.Itoa(req.Size)

	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取任务详情
func (t Tasks) Detail(taskId string) (resp Tasks, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Update 编辑任务
func (t Tasks) Update(taskId string, req UpdateRequest) (resp Tasks, err error) {
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
func (t Tasks) Action(taskId string, req ActionRequest) error {
	request := flowablesdk.GetRequest(ActionApi, taskId)
	request.With(httpclient.WithJson(req))
	_, err := request.DoHttpRequest()
	return err
}

// Delete 删除任务 -- 无法删除整个流程的一部分任务
func (t Tasks) Delete(taskId string, req DeleteRequest) error {
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
