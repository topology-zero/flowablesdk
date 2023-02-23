package task_comment

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/runtime/tasks/%s/comments"
	detailUrl = baseUrl + "/%s"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	AddApi    = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
)
