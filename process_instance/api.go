package process_instance

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl           = "/runtime/process-instances"
	queryUrl          = "/query/process-instances"
	detailUrl         = baseUrl + "/%s"
	identityUrl       = detailUrl + "/identitylinks"
	identityDetailUrl = identityUrl + "/users/%s/%s"
	variablesUrl      = detailUrl + "/variable"
	variableDetailUrl = variablesUrl + "/%s"
)

var (
	ListApi            = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	DetailApi          = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	UpdateApi          = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi          = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	StartApi           = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	ListCandidateApi   = flowablesdk.NewApi(httpclient.GET, identityUrl, flowablesdk.ProcessPrefix)
	AddCandidateApi    = flowablesdk.NewApi(httpclient.POST, identityUrl, flowablesdk.ProcessPrefix)
	DeleteCandidateApi = flowablesdk.NewApi(httpclient.DELETE, identityDetailUrl, flowablesdk.ProcessPrefix)
	ListVariablesApi   = flowablesdk.NewApi(httpclient.GET, variablesUrl, flowablesdk.ProcessPrefix)
	AddVariablesApi    = flowablesdk.NewApi(httpclient.POST, variablesUrl, flowablesdk.ProcessPrefix)
	UpdateVariablesApi = flowablesdk.NewApi(httpclient.PUT, variablesUrl, flowablesdk.ProcessPrefix)
	UpdateVariableApi  = flowablesdk.NewApi(httpclient.PUT, variableDetailUrl, flowablesdk.ProcessPrefix)
	VariableDetailApi  = flowablesdk.NewApi(httpclient.GET, variableDetailUrl, flowablesdk.ProcessPrefix)
)
