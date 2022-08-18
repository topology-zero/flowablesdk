package deployment

import (
	"archive/zip"
	"bytes"
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type Deployment struct{}

// List 获取流程定义列表
func (d *Deployment) List(req ListRequest) (resp ListResponse, err error) {
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
		query["sort"] = "deployTime"
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

// Detail 获取流程定义列表
func (d *Deployment) Detail(deploymentId string) (resp Detail, err error) {
	request := flowablesdk.GetRequest(DetailApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Create 创建流程
func (d *Deployment) Create(req CreateRequest) (resp Detail, err error) {
	request := flowablesdk.GetRequest(CreateApi)
	var bf bytes.Buffer
	zipWriter := zip.NewWriter(&bf)

	zipFileWrite, err := zipWriter.Create(req.FileName)
	if err != nil {
		return
	}
	_, err = zipFileWrite.Write([]byte(req.Xml))
	if err != nil {
		return
	}

	_ = zipWriter.Close()

	data, err := request.With(httpclient.WithFile("file", req.FileName+".zip", &bf)).DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除流程
func (d *Deployment) Delete(deploymentId string) error {
	request := flowablesdk.GetRequest(DeleteApi, deploymentId)
	_, err := request.DoHttpRequest()
	if err != nil {
		return err
	}
	return nil
}

// Resource 查看流程资源列表
func (d *Deployment) Resource(deploymentId string) (resp []Resource, err error) {
	request := flowablesdk.GetRequest(ResourceApi, deploymentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ResourceDetail 查看流程资源详情
func (d *Deployment) ResourceDetail(deploymentId, resourceId string) (resp Resource, err error) {
	request := flowablesdk.GetRequest(ResourceDetailApi, deploymentId, resourceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ResourceContent 获取流程定义的xml
func (d *Deployment) ResourceContent(deploymentId, resourceId string) (resp string, err error) {
	request := flowablesdk.GetRequest(ResourceContentApi, deploymentId, resourceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	return string(data), nil
}
