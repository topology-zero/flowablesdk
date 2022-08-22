package task

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

// AddBinaryVariable 添加二进制变量 (单个)
func (t Tasks) AddBinaryVariable(taskId string, req variables.FileVariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(AddBinaryVariableApi, taskId)

	maps := map[string]string{
		"name": req.Name,
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

// UpdateBinaryVariable 编辑二进制变量 (单个)
func (t Tasks) UpdateBinaryVariable(taskId string, req variables.FileVariableRequest) (resp variables.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateBinaryVariableApi, taskId, req.Name)

	maps := map[string]string{
		"name": req.Name,
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

// DetailBinaryVariable 查看二进制变量 (单个)
func (t Tasks) DetailBinaryVariable(taskId, variableName string, scope ...string) (resp []byte, err error) {
	request := flowablesdk.GetRequest(DetailBinaryVariableApi, taskId, variableName)
	query := map[string]string{}

	if len(scope) != 0 {
		query["scope"] = scope[0]
	}

	request.With(httpclient.WithQuery(query))
	resp, err = request.DoHttpRequest()
	return
}
