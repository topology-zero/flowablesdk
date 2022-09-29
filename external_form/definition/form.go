package definition

import (
	"encoding/json"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/external_form/model"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type ExternalFormDefinition struct {
	Id           string
	Url          string
	Category     string
	Name         string
	Key          string
	Description  string
	Version      int
	ResourceName string
	DeploymentId string
	TenantId     string
}

// List 获取外置表单列表
func (f *ExternalFormDefinition) List(req ListRequest) (resp ListResponse, err error) {
	query := make(map[string]string)

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.Key) > 0 {
		query["key"] = req.Key
	}

	if len(req.Name) > 0 {
		query["name"] = req.Name
	}

	if len(req.ResourceName) > 0 {
		query["resourceName"] = req.ResourceName
	}

	if len(req.Version) > 0 {
		query["version"] = req.Version
	}

	if len(req.DeploymentId) > 0 {
		query["deploymentId"] = req.DeploymentId
	}

	if len(req.TenantId) > 0 {
		query["tenantId"] = req.TenantId
	}

	if req.Latest {
		query["Latest"] = "true"
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

// Detail 获取外置表单详情
func (f *ExternalFormDefinition) Detail(formDefinitionId string) (resp ExternalFormDefinition, err error) {
	request := flowablesdk.GetRequest(DetailApi, formDefinitionId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Model 查看表单外置的所有字段
func (f *ExternalFormDefinition) Model(deploymentId string) (resp model.Model, err error) {
	request := flowablesdk.GetRequest(ModelApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
