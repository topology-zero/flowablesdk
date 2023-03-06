package history_process_instance

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl       = "/history/historic-process-instances"
	queryUrl      = "/query/historic-process-instances"
	detailUrl     = baseUrl + "/%s"
	identityUrl   = detailUrl + "/identitylinks"
	binaryDataUrl = detailUrl + "/variable/%s/data"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi       = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
	ListIdentityApi = flowablesdk.NewApi(httpclient.GET, identityUrl, flowablesdk.ProcessPrefix)
	BinaryDataApi   = flowablesdk.NewApi(httpclient.GET, binaryDataUrl, flowablesdk.ProcessPrefix)
)
