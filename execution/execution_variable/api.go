package execution_variable

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl       = "/runtime/executions/%s/variables"
	detailUrl     = baseUrl + "/%s"
	binaryDataUrl = detailUrl + "/data"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	AddApi          = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	UpdateApi       = flowablesdk.NewApi(httpclient.PUT, baseUrl, flowablesdk.ProcessPrefix)
	DeleteApi       = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
	DeleteAllApi    = flowablesdk.NewApi(httpclient.DELETE, baseUrl, flowablesdk.ProcessPrefix)
	AddBinaryApi    = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	UpdateBinaryApi = flowablesdk.NewApi(httpclient.PUT, detailUrl, flowablesdk.ProcessPrefix)
	BinaryDataApi   = flowablesdk.NewApi(httpclient.GET, binaryDataUrl, flowablesdk.ProcessPrefix)
)
