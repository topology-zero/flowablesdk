package deployment

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

type Deployment struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	DeploymentTime string `json:"deploymentTime"`
	Category       string `json:"category"`
	Url            string `json:"url"`
	TenantId       string `json:"tenantId"`
}

// List 获取流程定义列表
func List(req ListRequest) (list []Deployment, count int, err error) {
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

// Detail 获取流程定义列表
func Detail(deploymentId string) (resp Deployment, err error) {
	request := flowablesdk.GetRequest(DetailApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Create 创建流程
func Create(req CreateRequest) (resp Deployment, err error) {
	if !strings.HasSuffix(req.FileName, ".bpmn20.xml") {
		err = errors.New("suffix not .bpmn20.xml")
	}

	request := flowablesdk.GetRequest(CreateApi)
	data, err := request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     strings.NewReader(req.Xml),
	}, nil)).DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除流程
func Delete(deploymentId string) error {
	request := flowablesdk.GetRequest(DeleteApi, deploymentId)
	_, err := request.DoHttpRequest()
	return err
}

// Resource 查看流程资源列表
func Resource(deploymentId string) (resp []ResourceResponse, err error) {
	request := flowablesdk.GetRequest(ResourceApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ResourceDetail 查看流程资源详情
func ResourceDetail(deploymentId, resourceId string) (resp ResourceResponse, err error) {
	request := flowablesdk.GetRequest(ResourceDetailApi, deploymentId, resourceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ResourceContent 获取流程定义的xml
func ResourceContent(deploymentId, resourceId string) (resp string, err error) {
	request := flowablesdk.GetRequest(ResourceContentApi, deploymentId, resourceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	return string(data), nil
}
