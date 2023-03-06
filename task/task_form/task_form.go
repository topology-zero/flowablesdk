package task_form

import (
	"encoding/base64"
	"encoding/json"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/external_form/model"
)

// GetForm 获取 task 的外部表单
func GetForm(taskId string) (resp model.Model, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	decode, _ := base64.StdEncoding.DecodeString(string(data))
	err = json.Unmarshal(decode, &resp)
	return
}
