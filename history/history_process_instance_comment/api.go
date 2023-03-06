package history_process_instance_comment

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/history/historic-process-instances/%s/comments"
	detailUrl = baseUrl + "/%s"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	AddApi    = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
)
