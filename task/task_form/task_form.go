package task_form

import (
	"github.com/MasterJoyHunan/flowablesdk"
)

// GetForm 获取 task 的外部表单
func GetForm(taskId string) ([]byte, error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	return request.DoHttpRequest()
}
