package task

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/comment"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// ListComments 获取任务全部备注
func (t Tasks) ListComments(taskId string) (resp []comment.Comment, err error) {
	request := flowablesdk.GetRequest(ListCommentsApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DetailComments 获取任务单个备注
func (t Tasks) DetailComments(taskId, commentsId string) (resp comment.Comment, err error) {
	request := flowablesdk.GetRequest(DetailCommentsApi, taskId, commentsId)
	data, err := request.DoHttpRequest()

	err = json.Unmarshal(data, &resp)
	return
}

// AddComments 为任务添加备注
func (t Tasks) AddComments(taskId string, req AddCommentsRequest) (resp comment.Comment, err error) {
	request := flowablesdk.GetRequest(AddCommentsApi, taskId)
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

// DeleteComments 为任务删除备注
func (t Tasks) DeleteComments(taskId, commentsId string) error {
	request := flowablesdk.GetRequest(DeleteCommentsApi, taskId, commentsId)
	_, err := request.DoHttpRequest()
	return err
}
