package history_process_instances

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl     = "/history/historic-process-instances"
	queryUrl    = "/query/historic-process-instances"
	detailUrl   = baseUrl + "/%s"
	identityUrl = detailUrl + "/identitylinks"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi       = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
	ListIdentityApi = flowablesdk.NewApi(httpclient.GET, identityUrl, flowablesdk.ProcessPrefix)
)
