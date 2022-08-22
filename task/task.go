package task

import (
	"encoding/json"
	"errors"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type Tasks struct {
	Assignee          string               `json:"assignee"`
	CreateTime        string               `json:"createTime"`
	DelegationState   string               `json:"delegationState"`
	Description       string               `json:"description"`
	DueDate           string               `json:"dueDate"`
	Execution         string               `json:"execution"`
	Id                string               `json:"id"`
	Name              string               `json:"name"`
	Owner             string               `json:"owner"`
	ParentTask        string               `json:"parentTask"`
	Priority          int                  `json:"priority"`
	ProcessDefinition string               `json:"processDefinition"`
	ProcessInstance   string               `json:"processInstance"`
	Suspended         bool                 `json:"suspended"`
	TaskDefinitionKey string               `json:"taskDefinitionKey"`
	Url               string               `json:"url"`
	TenantId          string               `json:"tenantId"`
	Variables         []variables.Variable `json:"variables"`
}

// List 获取任务列表
func (t Tasks) List(req ListRequest) (resp ListResponse, err error) {
	query := make(map[string]string)

	if len(req.Name) > 0 {
		query["nameLike"] = req.Name
	}

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.Sort) > 0 {
		query["sort"] = req.Sort
	} else {
		query["sort"] = "id"
	}

	// TODO 好多参数未完成

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
