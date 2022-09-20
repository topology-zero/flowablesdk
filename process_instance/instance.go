package process_instance

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/candidate"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type Instance struct {
	Id                   string `json:"id"`
	Url                  string `json:"url"`
	BusinessKey          string `json:"businessKey"`
	Suspended            bool   `json:"suspended"`
	ProcessDefinitionUrl string `json:"processDefinitionUrl"`
	ActivityId           string `json:"activityId"`
	TenantId             string `json:"tenantId"`
}

// List 返回所有流程实例
func (i Instance) List(req ListRequest) (resp ListResponse, err error) {
	request := flowablesdk.GetRequest(ListApi)

	query := map[string]string{}

	if len(req.Id) > 0 {
		query["id"] = req.Id
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

	if req.Suspended != nil {
		if *req.Suspended {
			query["involvedUser"] = "true"
		} else {
			query["involvedUser"] = "false"
		}
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

// Detail 返回单个流程实例
func (i Instance) Detail(instanceId string) (resp Instance, err error) {
	request := flowablesdk.GetRequest(DetailApi, instanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Update 更新一个流程实例
func (i Instance) Update(instanceId string, req UpdateRequest) (resp Instance, err error) {
	request := flowablesdk.GetRequest(UpdateApi, instanceId)

	if len(req.Action) > 0 {
		if !(req.Action == "activate" || req.Action == "suspend") {
			err = errors.New("action is not allow")
			return
		}
	}

	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个流程实例
func (i Instance) Delete(instanceId string) error {
	request := flowablesdk.GetRequest(DeleteApi, instanceId)
	_, err := request.DoHttpRequest()
	return err
}

// Start 启动一个流程实例
func (i Instance) Start(req StartProcessRequest) (resp Instance, err error) {
	request := flowablesdk.GetRequest(StartApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ListCandidate 获取单个流程实例所有相关人员
func (i Instance) ListCandidate(instanceId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListCandidateApi, instanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddCandidate 单个流程实例添加参与用户
func (i Instance) AddCandidate(instanceId string, req candidate.Candidate) (resp candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(AddCandidateApi, instanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// DeleteCandidate 单个流程实例删除参与用户
func (i Instance) DeleteCandidate(instanceId string, req candidate.Candidate) error {
	request := flowablesdk.GetRequest(DeleteCandidateApi, instanceId, req.User, req.Type)
	_, err := request.DoHttpRequest()
	return err
}

// ListVariables 获取单个流程实例所有变量
func (i Instance) ListVariables(instanceId string) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(ListVariablesApi, instanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddVariables 单个流程实例添加变量
func (i Instance) AddVariables(instanceId string, req []variables.VariableRequest) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(AddVariablesApi, instanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// UpdateVariables 单个流程实例编辑变量(多个)
func (i Instance) UpdateVariables(instanceId string, req []variables.VariableRequest) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariablesApi, instanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// UpdateVariable 单个流程实例编辑变量(单个)
func (i Instance) UpdateVariable(instanceId, variable string, req variables.VariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariableApi, instanceId, variable)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// VariableDetail 获取单个流程实例所有变量
func (i Instance) VariableDetail(instanceId, variable string) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(VariableDetailApi, instanceId, variable)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddFileVariable 单个流程实例添加文件变量
func (i Instance) AddFileVariable(instanceId string, req variables.FileVariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(AddVariablesApi, instanceId)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, map[string]string{
		"name": req.Name,
		"type": "binary",
	}))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// UpdateFileVariable 单个流程实例编辑文件变量
func (i Instance) UpdateFileVariable(instanceId string, req variables.FileVariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariableApi, instanceId, req.Name)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, map[string]string{
		"name": req.Name,
		"type": "binary",
	}))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}
