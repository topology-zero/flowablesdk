package task_events

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/runtime/tasks/%s/events"
	detailUrl = baseUrl + "/%s"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteEventsApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix) // 源码中提供了删除接口,文档中没有提供,暂不实现该接口
)
