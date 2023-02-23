package task_comment

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/comment"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// List 获取任务全部备注
func List(taskId string) (resp []comment.Comment, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取任务单个备注
func Detail(taskId, commentsId string) (resp comment.Comment, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId, commentsId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Add 为任务添加备注
func Add(taskId string, req comment.AddComment) (resp comment.Comment, err error) {
	request := flowablesdk.GetRequest(AddApi, taskId)
	request.With(
		httpclient.WithJson(req),
		httpclient.WithHeader(httpclient.FlowableUserId, req.UserId),
	)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 为任务删除备注
func Delete(taskId, commentsId string) error {
	request := flowablesdk.GetRequest(DeleteApi, taskId, commentsId)
	_, err := request.DoHttpRequest()
	return err
}
