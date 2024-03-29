package history_variable_instance

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl       = "/history/historic-variable-instances"
	queryUrl      = "/query/historic-variable-instances"
	binaryDataUrl = baseUrl + "/%s/data"
)

var (
	ListApi       = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	BinaryDataApi = flowablesdk.NewApi(httpclient.GET, binaryDataUrl, flowablesdk.ProcessPrefix)
)
