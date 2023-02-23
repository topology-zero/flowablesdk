package history_task_instances

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
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
