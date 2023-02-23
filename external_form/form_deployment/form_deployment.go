package form_deployment

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type FormDeployment struct {
	Id                 string
	Name               string
	DeploymentTime     time.Time
	Category           string
	Url                string
	ParentDeploymentId string
	TenantId           string
}

// List 获取外置表单列表
func List(req ListRequest) (list []FormDeployment, count int, err error) {
	query := make(map[string]string)

	if len(req.Name) > 0 {
		query["name"] = req.Name
	}

	if len(req.NameLike) > 0 {
		query["nameLike"] = req.NameLike
	}

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.CategoryNotEquals) > 0 {
		query["categoryNotEquals"] = req.CategoryNotEquals
	}

	if len(req.ParentDeploymentId) > 0 {
		query["parentDeploymentId"] = req.ParentDeploymentId
	}

	if len(req.ParentDeploymentIdLike) > 0 {
		query["parentDeploymentIdLike"] = req.ParentDeploymentIdLike
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

// Add 添加外置表单
func Add(req AddRequest) (resp FormDeployment, err error) {
	request := flowablesdk.GetRequest(AddApi)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     strings.NewReader(req.Data),
	}, map[string]string{
		"tenantId": req.TenantId,
		"category": req.Category,
	}))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 部署表单详情
func Detail(deploymentId string) (resp FormDeployment, err error) {
	request := flowablesdk.GetRequest(DetailApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除部署表单
func Delete(deploymentId string) error {
	request := flowablesdk.GetRequest(DeleteApi, deploymentId)
	_, err := request.DoHttpRequest()
	return err
}
