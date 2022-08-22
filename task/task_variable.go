package task

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

// ListVariables 获取任务的所有变量
func (t Tasks) ListVariables(taskId string, scope ...string) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(ListVariablesApi, taskId)
	query := map[string]string{}

	if len(scope) != 0 {
		query["scope"] = scope[0]
	}

	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DetailVariables 获取任务的所有变量
func (t Tasks) DetailVariables(taskId, variableName string, scope ...string) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(DetailVariableApi, taskId, variableName)
	query := map[string]string{}

	if len(scope) != 0 {
		query["scope"] = scope[0]
	}

	request.With(httpclient.WithQuery(query))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// AddVariables 为任务添加变量 (多个)
func (t Tasks) AddVariables(taskId string, variables []variables.VariableRequest) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(AddVariablesApi, taskId)
	request.With(httpclient.WithJson(variables))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// UpdateVariable 为任务修改变量 (单个)
func (t Tasks) UpdateVariable(taskId string, variable variables.VariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariableApi, taskId, variable.Name)
	request.With(httpclient.WithJson(variable))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DeleteVariables 为任务删除所有 scope = local 的变量
func (t Tasks) DeleteVariables(taskId string) error {
	request := flowablesdk.GetRequest(DeleteVariablesApi, taskId)
	_, err := request.DoHttpRequest()
	return err
}

// DeleteVariable 为任务删除变量 (单个)
func (t Tasks) DeleteVariable(taskId, variable string) error {
	request := flowablesdk.GetRequest(DeleteVariableApi, taskId, variable)
	_, err := request.DoHttpRequest()
	return err
}
