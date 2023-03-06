package form_deployment

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/form-repository/deployments"
	detailUrl = baseUrl + "/%s"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.FormPrefix)
	AddApi    = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.FormPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.FormPrefix)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.FormPrefix)
)
