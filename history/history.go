package history

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type History struct {
	Id                     string               `json:"id"`
	BusinessKey            string               `json:"businessKey"`
	ProcessDefinitionId    string               `json:"processDefinitionId"`
	ProcessDefinitionUrl   string               `json:"processDefinitionUrl"`
	StartTime              string               `json:"startTime"`
	EndTime                string               `json:"endTime"`
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

type ActivityHistory struct {
	Id                      string     `json:"id"`
	ActivityId              string     `json:"activityId"`
	ActivityName            string     `json:"activityName"`
	ActivityType            string     `json:"activityType"`
	ProcessDefinitionId     string     `json:"processDefinitionId"`
	ProcessDefinitionUrl    string     `json:"processDefinitionUrl"`
	ProcessInstanceId       string     `json:"processInstanceId"`
	ProcessInstanceUrl      string     `json:"processInstanceUrl"`
	ExecutionId             string     `json:"executionId"`
	TaskId                  string     `json:"taskId"`
	CalledProcessInstanceId string     `json:"calledProcessInstanceId"`
	Assignee                string     `json:"assignee"`
	StartTime               *time.Time `json:"startTime"`
	EndTime                 *time.Time `json:"endTime"`
	DurationInMillis        int        `json:"durationInMillis"`
	TenantId                string     `json:"tenantId"`
}

// List 获取所有流程实例历史流转
func (h History) List(req ListRequest) (resp ListResponse, err error) {
	request := flowablesdk.GetRequest(ListApi)

	query := map[string]string{}

	if len(req.ProcessInstanceId) > 0 {
		query["processInstanceId"] = req.ProcessInstanceId
	}

	if len(req.ProcessDefinitionKey) > 0 {
		query["processDefinitionKey"] = req.ProcessDefinitionKey
	}

	if len(req.ProcessDefinitionId) > 0 {
		query["processDefinitionId"] = req.ProcessDefinitionId
	}

	if len(req.BusinessKey) > 0 {
		query["businessKey"] = req.BusinessKey
	}

	if len(req.InvolvedUser) > 0 {
		query["involvedUser"] = req.InvolvedUser
	}

	if req.Finished != nil {
		if *req.Finished {
			query["finished"] = "true"
		} else {
			query["finished"] = "false"
		}
	}

	if len(req.FinishedTime) == 2 {
		query["finishedAfter"] = req.FinishedTime[0].Format(time.RFC3339)
		query["finishedBefore"] = req.FinishedTime[1].Format(time.RFC3339)
	}

	if len(req.StartedTime) == 2 {
		query["startedAfter"] = req.StartedTime[0].Format(time.RFC3339)
		query["startedBefore"] = req.StartedTime[1].Format(time.RFC3339)
	}

	if len(req.StartedBy) > 0 {
		query["startedBy"] = req.StartedBy
	}

	if len(req.TenantId) > 0 {
		query["tenantId"] = req.TenantId
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

	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取单个流程实例历史流转
func (h History) Detail(processInstanceId string) (resp History, err error) {
	request := flowablesdk.GetRequest(DetailApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个流程实例历史流转
func (h History) Delete(processInstanceId string) error {
	request := flowablesdk.GetRequest(DeleteApi, processInstanceId)
	_, err := request.DoHttpRequest()
	return err
}

// ListIdentity 删除单个流程实例历史流转相关用户
func (h History) ListIdentity(processInstanceId string) (resp []IdentityResponse, err error) {
	request := flowablesdk.GetRequest(ListIdentityApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
