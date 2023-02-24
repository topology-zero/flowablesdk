package execution_variable

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

// List 获取流程执行变量
func List(executionId string, scope ...string) (list []variable.Variable, err error) {
	request := flowablesdk.GetRequest(ListApi, executionId)
	if len(scope) > 0 {
		request.With(httpclient.WithQuery(map[string]string{
			"scope": scope[0],
		}))
	}
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &list)
	return
}

// Detail 获取单个流程执行变量
func Detail(executionId, variableName string, scope ...string) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(DetailApi, executionId, variableName)
	if len(scope) > 0 {
		request.With(httpclient.WithQuery(map[string]string{
			"scope": scope[0],
		}))
	}
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Add 添加流程执行变量
func Add(executionId string, req []variable.VariableRequest) (resp []variable.Variable, err error) {
	request := flowablesdk.GetRequest(AddApi, executionId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	err = json.Unmarshal(data, &resp)
	return
}

// Update 编辑流程执行变量
func Update(executionId string, req []variable.VariableRequest) (resp []variable.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateApi, executionId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	err = json.Unmarshal(data, &resp)
	return
}

// AddBinary 添加流程执行二进制变量
func AddBinary(executionId string, req variable.FileVariableRequest) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(AddBinaryApi, executionId)
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

// UpdateBinary 编辑流程执行二进制变量
func UpdateBinary(executionId, variableName string, req variable.FileVariableRequest) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateBinaryApi, executionId, variableName)
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

// Delete 删除流程执行变量
func Delete(executionId, variableName string) (err error) {
	request := flowablesdk.GetRequest(DeleteApi, executionId, variableName)
	_, err = request.DoHttpRequest()
	return
}

// DeleteAllLocal 删除所有流程执行 scope = local 的变量
func DeleteAllLocal(executionId string) (err error) {
	request := flowablesdk.GetRequest(DeleteAllApi, executionId)
	_, err = request.DoHttpRequest()
	return
}

// BinaryData 获取流程执行二进制文件
func BinaryData(executionId, variableName string) (data []byte, err error) {
	request := flowablesdk.GetRequest(BinaryDataApi, executionId, variableName)
	return request.DoHttpRequest()
}
