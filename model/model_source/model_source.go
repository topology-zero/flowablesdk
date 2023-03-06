package model_source

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

func DetailSource(modelId string) (data []byte, err error) {
	request := flowablesdk.GetRequest(DetailSourceApi, modelId)
	return request.DoHttpRequest()
}

func UpdateSource(modelId string, req SourceRequest) error {
	request := flowablesdk.GetRequest(UpdateSourceApi, modelId)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, nil))
	_, err := request.DoHttpRequest()
	return err
}

func DetailExtraSource(modelId string) (data []byte, err error) {
	request := flowablesdk.GetRequest(DetailExtraSourceApi, modelId)
	return request.DoHttpRequest()
}

func UpdateExtraSource(modelId string, req SourceRequest) error {
	request := flowablesdk.GetRequest(UpdateExtraSourceApi, modelId)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, nil))
	_, err := request.DoHttpRequest()
	return err
}
