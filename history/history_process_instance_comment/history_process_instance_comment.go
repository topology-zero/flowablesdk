package history_process_instance_comment

import (
	"encoding/json"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/comment"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

// List 单个流程实例历史相关备注
func List(processInstanceId string) (resp []comment.Comment, err error) {
	request := flowablesdk.GetRequest(ListApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Add 单个流程实例历史相关备注
func Add(processInstanceId string, req comment.AddComment) (resp comment.Comment, err error) {
	request := flowablesdk.GetRequest(AddApi, processInstanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 单个流程实例历史单个备注
func Detail(processInstanceId, commentId string) (resp comment.Comment, err error) {
	request := flowablesdk.GetRequest(DetailApi, processInstanceId, commentId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个流程实例历史备注
func Delete(processInstanceId, commentId string) (err error) {
	request := flowablesdk.GetRequest(DeleteApi, processInstanceId, commentId)
	_, err = request.DoHttpRequest()
	return
}
