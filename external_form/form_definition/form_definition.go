package form_definition

import (
	"encoding/json"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/external_form/model"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type FormDefinition struct {
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
func List(req ListRequest) (list []FormDefinition, count int, err error) {
	query := make(map[string]string)

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.CategoryLike) > 0 {
		query["categoryLike"] = req.CategoryLike
	}

	if len(req.CategoryNotEquals) > 0 {
		query["categoryNotEquals"] = req.CategoryNotEquals
	}

	if len(req.Key) > 0 {
		query["key"] = req.Key
	}

	if len(req.KeyLike) > 0 {
		query["keyLike"] = req.KeyLike
	}

	if len(req.Name) > 0 {
		query["name"] = req.Name
	}

	if len(req.NameLike) > 0 {
		query["nameLike"] = req.NameLike
	}

	if len(req.ResourceName) > 0 {
		query["resourceName"] = req.ResourceName
	}

	if len(req.ResourceNameLike) > 0 {
		query["resourceNameLike"] = req.ResourceNameLike
	}

	if req.Version > 0 {
		query["version"] = strconv.Itoa(req.Version)
	}

	if req.VersionGreaterThan > 0 {
		query["versionGreaterThan"] = strconv.Itoa(req.VersionGreaterThan)
	}

	if req.VersionGreaterThanOrEquals > 0 {
		query["versionGreaterThanOrEquals"] = strconv.Itoa(req.VersionGreaterThanOrEquals)
	}

	if req.VersionLowerThan > 0 {
		query["versionLowerThan"] = strconv.Itoa(req.VersionLowerThan)
	}

	if req.VersionLowerThanOrEquals > 0 {
		query["versionLowerThanOrEquals"] = strconv.Itoa(req.VersionLowerThanOrEquals)
	}

	if len(req.DeploymentId) > 0 {
		query["deploymentId"] = req.DeploymentId
	}

	if req.Latest {
		query["Latest"] = "true"
	}

	common.MakeCommonQuery(req.ListCommonRequest, req.WithTenant, query)

	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithQuery(query))
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

// Detail 获取外置表单详情
func Detail(formDefinitionId string) (resp FormDefinition, err error) {
	request := flowablesdk.GetRequest(DetailApi, formDefinitionId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Model 查看表单外置的所有字段
func Model(deploymentId string) (resp model.Model, err error) {
	request := flowablesdk.GetRequest(ModelApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
