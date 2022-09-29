package instance

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/external_form/model"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/pkg/timefmt"
)

type ExternalFormInstance struct {
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
func (f *ExternalFormInstance) List(req ListRequest) (resp ListResponse, err error) {
	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Add 添加一个 form 实例
func (f *ExternalFormInstance) Add(req AddRequest) error {
	request := flowablesdk.GetRequest(AddApi)
	request.With(httpclient.WithJson(req))
	_, err := request.DoHttpRequest()
	return err
}

// Detail 查看 form 实例详情
func (f *ExternalFormInstance) Detail(formInstanceId string) (resp ExternalFormInstance, err error) {
	request := flowablesdk.GetRequest(DetailApi, formInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Model 查看 form 实例字段信息
func (f *ExternalFormInstance) Model(req AddRequest) (resp model.Model, err error) {
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
func (f *ExternalFormInstance) DetailAndModel(req AddRequest) (resp FormInstanceModelResponse, err error) {
	request := flowablesdk.GetRequest(DetailAndModelApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
