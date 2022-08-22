package task

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/attachment"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// ListAttachments 获取任务全部附件
func (t Tasks) ListAttachments(taskId string) (resp []attachment.Attachment, err error) {
	request := flowablesdk.GetRequest(ListAttachmentsApi, taskId)
	data, err := request.DoHttpRequest()

	err = json.Unmarshal(data, &resp)
	return
}

// AddAttachmentsUrl 任务添加网络附件
func (t Tasks) AddAttachmentsUrl(taskId string, req AddAttachmentsUrlRequest) (resp attachment.Attachment, err error) {
	request := flowablesdk.GetRequest(AddAttachmentsApi, taskId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()

	err = json.Unmarshal(data, &resp)
	return
}

// AddAttachments 任务添加附件
func (t Tasks) AddAttachments(taskId string, req AddAttachmentsRequest) (resp attachment.Attachment, err error) {
	request := flowablesdk.GetRequest(AddAttachmentsApi, taskId)
	maps := map[string]string{
		"name": req.Name,
	}
	if len(req.Type) > 0 {
		maps["type"] = req.Type
	}
	if len(req.Description) > 0 {
		maps["description"] = req.Description
	}

	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.Name,
		File:     req.File,
	}, maps))
	data, err := request.DoHttpRequest()

	err = json.Unmarshal(data, &resp)
	return
}

// DeleteAttachments 任务删除附件
func (t Tasks) DeleteAttachments(taskId, attachmentId string) error {
	request := flowablesdk.GetRequest(DeleteAttachmentsApi, taskId, attachmentId)
	_, err := request.DoHttpRequest()
	return err
}

// ContentAttachments 查看任务附件二进制
func (t Tasks) ContentAttachments(taskId, attachmentId string) (data []byte, err error) {
	request := flowablesdk.GetRequest(ContentAttachmentsApi, taskId, attachmentId)
	data, err = request.DoHttpRequest()
	return
}
