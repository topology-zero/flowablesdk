package model

import (
	"encoding/json"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type Model struct {
	Name           string     `json:"name"`
	Key            string     `json:"key"`
	Category       string     `json:"category"`
	Version        int        `json:"version"`
	MetaInfo       string     `json:"metaInfo"`
	DeploymentId   string     `json:"deploymentId"`
	Id             string     `json:"id"`
	Url            string     `json:"url"`
	CreateTime     *time.Time `json:"createTime"`
	LastUpdateTime *time.Time `json:"lastUpdateTime"`
	DeploymentUrl  string     `json:"deploymentUrl"`
	TenantId       string     `json:"tenantId"`
}

// List 返回模型列表
func List(req ListRequest) (list []Model, count int, err error) {
	query := make(map[string]string)

	if len(req.Id) > 0 {
		query["id"] = req.Id
	}

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.CategoryLike) > 0 {
		query["categoryLike"] = req.CategoryLike
	}

	if len(req.CategoryNotEquals) > 0 {
		query["categoryNotEquals"] = req.CategoryNotEquals
	}

	if len(req.Name) > 0 {
		query["name"] = req.Name
	}

	if len(req.NameLike) > 0 {
		query["nameLike"] = req.NameLike
	}

	if len(req.Key) > 0 {
		query["key"] = req.Key
	}

	if len(req.DeploymentId) > 0 {
		query["deploymentId"] = req.DeploymentId
	}

	if len(req.Version) > 0 {
		query["version"] = req.Version
	}

	if req.LatestVersion != nil && *req.LatestVersion {
		query["latestVersion"] = "true"
	}

	if req.Deployed != nil {
		if *req.Deployed {
			query["deployed"] = "true"
		} else {
			query["deployed"] = "false"
		}
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

// Detail 获取单个model
func Detail(modelId string) (resp Model, err error) {
	request := flowablesdk.GetRequest(DetailApi, modelId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

func Add(req AddRequest) (resp Model, err error) {
	request := flowablesdk.GetRequest(AddApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

func Update(modelId string, req AddRequest) (resp Model, err error) {
	request := flowablesdk.GetRequest(UpdateApi, modelId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

func Delete(modelId string) (err error) {
	request := flowablesdk.GetRequest(DeleteApi, modelId)
	_, err = request.DoHttpRequest()
	return err
}
