package process

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type Process struct{}

// List 获取流程列表
func (p *Process) List(req ListRequest) (resp ListResponse, err error) {
	query := make(map[string]string)

	if req.Version > 0 {
		query["version"] = strconv.Itoa(req.Version)
	}

	if len(req.Name) > 0 {
		query["nameLike"] = req.Name
	}

	if len(req.Key) > 0 {
		query["keyLike"] = req.Key
	}

	if len(req.ResourceName) > 0 {
		query["resourceNameLike"] = req.ResourceName
	}

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.DeploymentId) > 0 {
		query["deploymentId"] = req.DeploymentId
	}

	if len(req.StartableByUser) > 0 {
		query["startableByUser"] = req.StartableByUser
	}

	if req.Latest {
		query["latest"] = "true"
	}

	if req.Suspended {
		query["suspended"] = "true"
	}

	if len(req.Sort) > 0 {
		query["sort"] = req.Sort
	} else {
		query["sort"] = "name"
	}

	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取流程详情
func (p *Process) Detail(deploymentId string) (resp Detail, err error) {
	request := flowablesdk.GetRequest(DetailApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Update 更新流程
func (p *Process) Update(deploymentId string, req UpdateRequest) (resp Detail, err error) {
	request := flowablesdk.GetRequest(UpdateApi, deploymentId)

	query := map[string]any{}

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.Action) > 0 {
		if req.Action == "activate" || req.Action == "suspend" {
			query["action"] = req.Action
		} else {
			err = errors.New("action is not allow")
			return
		}
	}

	if req.IncludeProcessInstances {
		query["IncludeProcessInstances"] = req.IncludeProcessInstances
	}

	if len(req.Date) > 0 {
		query["date"] = req.Date
	}

	if len(query) == 0 {
		err = errors.New("request is empty")
		return
	}

	request.With(httpclient.WithJson(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// ResourceContent 获取流程的xml
func (p *Process) ResourceContent(deploymentId string) (resp string, err error) {
	request := flowablesdk.GetRequest(ResourceContentApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	return string(data), nil
}

// Model 获取流程的 BPMN model
func (p *Process) Model(deploymentId string) (resp map[string]any, err error) {
	request := flowablesdk.GetRequest(ModelApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ProcessCandidate 获取流程所有候选人
func (p *Process) ProcessCandidate(deploymentId string) (resp []Candidate, err error) {
	request := flowablesdk.GetRequest(ProcessCandidateApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ProcessAddCandidate 流程添加候选人 TODO 需要查找为什么 group 不能使用的问题
func (p *Process) ProcessAddCandidate(deploymentId string, req ProcessAddCandidateRequest) (resp Candidate, err error) {
	request := flowablesdk.GetRequest(ProcessAddCandidateApi, deploymentId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ProcessDeleteCandidate 流程删除候选人
func (p *Process) ProcessDeleteCandidate(deploymentId string, req ProcessDeleteCandidateRequest) error {
	if len(req.Family) == 0 {
		return errors.New("family is empty")
	}

	if len(req.IdentityId) == 0 {
		return errors.New("identityId is empty")
	}

	if !(req.Family == "users" || req.Family == "groups") {
		return errors.New("family only allow users or groups")
	}

	request := flowablesdk.GetRequest(ProcessDeleteCandidateApi, deploymentId, req.Family, req.IdentityId)
	_, err := request.DoHttpRequest()
	if err != nil {
		return err
	}
	return nil
}

func (p *Process) CandidateProcess(deploymentId string, req ProcessDeleteCandidateRequest) (resp Candidate, err error) {
	if len(req.Family) == 0 {
		err = errors.New("family is empty")
		return
	}

	if len(req.IdentityId) == 0 {
		err = errors.New("identityId is empty")
		return
	}

	if !(req.Family == "users" || req.Family == "groups") {
		err = errors.New("family only allow users or groups")
		return
	}

	request := flowablesdk.GetRequest(CandidateProcessApi, deploymentId, req.Family, req.IdentityId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
