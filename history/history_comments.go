package history

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/comment"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// ListCommentApi
// AddCommentApi
// DetailCommentApi
// DeleteCommentApi

// ListComment 单个流程实例历史流转相关备注
func (h History) ListComment(processInstanceId string) (resp []comment.Comment, err error) {
	request := flowablesdk.GetRequest(ListCommentApi, processInstanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// AddComment 单个流程实例历史流转相关备注
func (h History) AddComment(processInstanceId string, req AddCommentRequest) (resp []comment.Comment, err error) {
	request := flowablesdk.GetRequest(AddCommentApi, processInstanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
