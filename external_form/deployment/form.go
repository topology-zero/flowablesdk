package deployment

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type ExternalFormDeployment struct {
	Id                 string
	Name               string
	DeploymentTime     time.Time
	Category           string
	Url                string
	ParentDeploymentId string
	TenantId           string
}

// List 获取外置表单列表
func (f *ExternalFormDeployment) List(req ListRequest) (resp ListResponse, err error) {
	query := make(map[string]string)

	if len(req.Category) > 0 {
		query["category"] = req.Category
	}

	if len(req.Name) > 0 {
		query["name"] = req.Name
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

	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Add 添加外置表单
func (f *ExternalFormDeployment) Add(req AddRequest) (resp ExternalFormDeployment, err error) {
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
func (f *ExternalFormDeployment) Detail(deploymentId string) (resp ExternalFormDeployment, err error) {
	request := flowablesdk.GetRequest(DetailApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除部署表单
func (f *ExternalFormDeployment) Delete(deploymentId string) error {
	request := flowablesdk.GetRequest(DeleteApi, deploymentId)
	_, err := request.DoHttpRequest()
	return err
}
