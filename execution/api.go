package execution

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/runtime/executions"
	queryUrl  = "/query/executions"
	detailUrl = baseUrl + "/%s"
	activeUrl = detailUrl + "/activities"
)

var (
	ListApi    = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	DetailApi  = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	ActiveApi  = flowablesdk.NewApi(httpclient.GET, activeUrl, flowablesdk.ProcessPrefix)
	ExecuteApi = flowablesdk.NewApi(httpclient.PUT, detailUrl, flowablesdk.ProcessPrefix)
)
