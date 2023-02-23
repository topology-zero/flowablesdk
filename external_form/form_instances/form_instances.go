package form_instances

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/external_form/model"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/pkg/timefmt"
)

type FormInstances struct {
	Id                  string
	FormDefinitionId    string
	TaskId              string
	ProcessInstanceId   string
	ProcessDefinitionId string
	ScopeId             string
	ScopeType           string
	ScopeDefinitionId   string
	SubmittedDate       *timefmt.Time
	SubmittedBy         string
	FormValuesId        string
	TenantId            string
	Url                 string
}

// List 查看 form 实例列表
func List(req ListRequest) (list []FormInstances, count int, err error) {
	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithJson(req))
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

// Add 添加一个 form 实例
func Add(req AddRequest) error {
	request := flowablesdk.GetRequest(AddApi)
	request.With(httpclient.WithJson(req))
	_, err := request.DoHttpRequest()
	return err
}

// Detail 查看 form 实例详情
func Detail(formInstanceId string) (resp FormInstances, err error) {
	request := flowablesdk.GetRequest(DetailApi, formInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Model 查看 form 实例字段信息
func Model(req AddRequest) (resp model.Model, err error) {
	request := flowablesdk.GetRequest(ModelApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DetailAndModel 查看 form 实例详情 + 字段信息
func DetailAndModel(req AddRequest) (resp FormInstanceModelResponse, err error) {
	request := flowablesdk.GetRequest(DetailAndModelApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
