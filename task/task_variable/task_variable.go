package task_variable

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

// List 获取任务的所有变量
func List(taskId string, scope ...string) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
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

// Detail 获取任务的单个变量
func Detail(taskId, variableName string, scope ...string) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId, variableName)
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

// DetailBinary 获取任务的单个文件
func DetailBinary(taskId, variableName string, scope ...string) (data []byte, err error) {
	request := flowablesdk.GetRequest(DetailBinaryApi, taskId, variableName)
	query := map[string]string{}

	if len(scope) != 0 {
		query["scope"] = scope[0]
	}

	request.With(httpclient.WithQuery(query))
	return request.DoHttpRequest()
}

// Add 为任务添加变量 (多个)
func Add(taskId string, variables []variables.VariableRequest) (resp []variables.Variable, err error) {
	request := flowablesdk.GetRequest(AddApi, taskId)
	request.With(httpclient.WithJson(variables))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// AddBinary 为任务添加文件 (单个)
func AddBinary(taskId string, req variables.FileVariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(AddApi, taskId)
	maps := map[string]string{
		"name": req.VariableName,
		"type": "binary",
	}
	if len(req.Scope) > 0 {
		maps["scope"] = req.Scope
	}

	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, maps))

	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Update 为任务修改变量 (单个)
func Update(taskId string, variable variables.VariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateApi, taskId, variable.Name)
	request.With(httpclient.WithJson(variable))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// UpdateBinary 为任务修改文件 (单个)
func UpdateBinary(taskId string, req variables.FileVariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateApi, taskId, req.VariableName)
	maps := map[string]string{
		"name": req.VariableName,
		"type": "binary",
	}
	if len(req.Scope) > 0 {
		maps["scope"] = req.Scope
	}

	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, maps))

	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DeleteAllLocal 为任务删除所有 scope = local 的变量
func DeleteAllLocal(taskId string) error {
	request := flowablesdk.GetRequest(DeleteAllApi, taskId)
	_, err := request.DoHttpRequest()
	return err
}

// Delete 为任务删除变量 (单个)
func Delete(taskId, variable string) error {
	request := flowablesdk.GetRequest(DeleteApi, taskId, variable)
	_, err := request.DoHttpRequest()
	return err
}
