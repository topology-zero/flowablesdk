package form

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

type Form struct {
	FormKey              string           `json:"formKey"`
	DeploymentId         string           `json:"deploymentId"`
	ProcessDefinitionId  string           `json:"processDefinitionId"`
	ProcessDefinitionUrl string           `json:"processDefinitionUrl"`
	TaskId               string           `json:"taskId"`
	TaskUrl              string           `json:"taskUrl"`
	FormProperties       []FormProperties `json:"formProperties"`
}

type FormProperties struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Value       any          `json:"value"`
	Readable    bool         `json:"readable"`
	Writable    bool         `json:"writable"`
	Required    bool         `json:"required"`
	DatePattern string       `json:"datePattern"`
	EnumValues  []EnumValues `json:"enumValues"`
}

type EnumValues struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetFrom 获取 task/processDefinition 的表格
func (f *Form) GetFrom(req GetFromRequest) (resp Form, err error) {
	request := flowablesdk.GetRequest(DetailFromApi)

	query := map[string]string{}

	if len(req.TaskId) > 0 {
		query["taskId"] = req.TaskId
	} else {
		query["processDefinitionId"] = req.ProcessDefinitionId
	}

	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// SubmitForm 提交 task/processDefinition 的表格
func (f *Form) SubmitForm(req SubmitFromRequest) (resp SubmitFromResponse, err error) {
	request := flowablesdk.GetRequest(SubmitFromApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
