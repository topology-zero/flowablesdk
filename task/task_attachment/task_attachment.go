package task_attachment

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/attachment"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// List 获取任务全部附件
func List(taskId string) (resp []attachment.Attachment, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// AddUrl 任务添加网络附件
func AddUrl(taskId string, req attachment.AddAttachment) (resp attachment.Attachment, err error) {
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

// Add 任务添加附件
func Add(taskId string, req attachment.AddAttachment) (resp attachment.Attachment, err error) {
	request := flowablesdk.GetRequest(AddApi, taskId)
	maps := map[string]string{
		"name": req.Name,
	}
	if len(req.Type) > 0 {
		maps["type"] = req.Type
	}
	if len(req.Description) > 0 {
		maps["description"] = req.Description
	}

	request.With(
		httpclient.WithMultipartFrom(&httpclient.UploadFile{
			Field:    "file",
			FileName: req.Name,
			File:     req.File,
		}, maps),
		httpclient.WithHeader(httpclient.FlowableUserId, req.UserId),
	)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 任务删除附件
func Delete(taskId, attachmentId string, userId ...string) error {
	request := flowablesdk.GetRequest(DeleteApi, taskId, attachmentId)
	if len(userId) > 0 {
		request.With(httpclient.WithHeader(httpclient.FlowableUserId, userId[0]))
	}
	_, err := request.DoHttpRequest()
	return err
}

// Detail 查看任务附件二进制
func Detail(taskId, attachmentId string) (data []byte, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId, attachmentId)
	return request.DoHttpRequest()
}
