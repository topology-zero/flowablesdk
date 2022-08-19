package process_definition

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/candidate"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type Process struct {
	Id                       string `json:"id"`
	Url                      string `json:"url"`
	Version                  int    `json:"version"`
	Key                      string `json:"key"`
	Category                 string `json:"category"`
	Suspended                bool   `json:"suspended"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	DeploymentId             string `json:"deploymentId"`
	DeploymentUrl            string `json:"deploymentUrl"`
	GraphicalNotationDefined bool   `json:"graphicalNotationDefined"`
	Resource                 string `json:"resource"`
	DiagramResource          string `json:"diagramResource"`
	StartFormDefined         bool   `json:"startFormDefined"`
}

// List 获取流程列表
func (p Process) List(req ListRequest) (resp ListResponse, err error) {
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
func (p Process) Detail(deploymentId string) (resp Process, err error) {
	request := flowablesdk.GetRequest(DetailApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Update 更新流程
func (p Process) Update(deploymentId string, req UpdateRequest) (resp Process, err error) {
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
func (p Process) ResourceContent(deploymentId string) (resp string, err error) {
	request := flowablesdk.GetRequest(ResourceContentApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	return string(data), nil
}

// Model 获取流程的 BPMN model
func (p Process) Model(deploymentId string) (resp map[string]any, err error) {
	request := flowablesdk.GetRequest(ModelApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ListCandidate 获取流程所有候选人
func (p Process) ListCandidate(deploymentId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListCandidateApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddCandidate 流程添加候选人 TODO 需要查找为什么 group 不能使用的问题
func (p Process) AddCandidate(deploymentId string, req AddCandidateRequest) (resp candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(AddCandidateApi, deploymentId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// DeleteCandidate 流程删除候选人
func (p Process) DeleteCandidate(deploymentId string, req DeleteCandidateRequest) error {
	if len(req.Family) == 0 {
		return errors.New("family is empty")
	}

	if len(req.CandidateId) == 0 {
		return errors.New("identityId is empty")
	}

	if !(req.Family == "users" || req.Family == "groups") {
		return errors.New("family only allow users or groups")
	}

	request := flowablesdk.GetRequest(DeleteCandidateApi, deploymentId, req.Family, req.CandidateId)
	_, err := request.DoHttpRequest()
	if err != nil {
		return err
	}
	return nil
}

// CandidateDetail 查看流程指定候选人
func (p Process) CandidateDetail(deploymentId string, req DeleteCandidateRequest) (resp candidate.Candidate, err error) {
	if len(req.Family) == 0 {
		err = errors.New("family is empty")
		return
	}

	if len(req.CandidateId) == 0 {
		err = errors.New("identityId is empty")
		return
	}

	if !(req.Family == "users" || req.Family == "groups") {
		err = errors.New("family only allow users or groups")
		return
	}

	request := flowablesdk.GetRequest(CandidateDetailApi, deploymentId, req.Family, req.CandidateId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
