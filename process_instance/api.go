package process_instance

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl           = "/service/runtime/process-instances"
	detailUrl         = baseUrl + "/%s"
	identityUrl       = detailUrl + "/identitylinks"
	identityDetailUrl = identityUrl + "/users/%s/%s"
	variablesUrl      = detailUrl + "/variables"
	variableDetailUrl = variablesUrl + "/%s"
)

var (
	ListApi            = flowablesdk.NewApi(httpclient.GET, baseUrl)
	DetailApi          = flowablesdk.NewApi(httpclient.GET, detailUrl)
	UpdateApi          = flowablesdk.NewApi(httpclient.GET, detailUrl)
	DeleteApi          = flowablesdk.NewApi(httpclient.GET, detailUrl)
	StartApi           = flowablesdk.NewApi(httpclient.POST, baseUrl)
	ListCandidateApi   = flowablesdk.NewApi(httpclient.GET, identityUrl)
	AddCandidateApi    = flowablesdk.NewApi(httpclient.POST, identityUrl)
	DeleteCandidateApi = flowablesdk.NewApi(httpclient.DELETE, identityDetailUrl)
	ListVariablesApi   = flowablesdk.NewApi(httpclient.GET, variablesUrl)
	AddVariablesApi    = flowablesdk.NewApi(httpclient.POST, variablesUrl)
	EditVariablesApi   = flowablesdk.NewApi(httpclient.PUT, variablesUrl)
	EditVariableApi    = flowablesdk.NewApi(httpclient.PUT, variableDetailUrl)
	VariableDetailApi  = flowablesdk.NewApi(httpclient.GET, variableDetailUrl)
)
