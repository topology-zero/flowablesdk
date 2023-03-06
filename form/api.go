package form

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl = "/form/form-data"
)

var (
	DetailFromApi = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	SubmitFromApi = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
)
