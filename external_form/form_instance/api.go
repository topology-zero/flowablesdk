package form_instance

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl           = "/form/form-instances"
	queryUrl          = "/query/form-instances"
	detailUrl         = "/form/form-instance/%s"
	modelUrl          = "/form/models"
	detailAndModelUrl = "/form/form-instance-models"
)

var (
	ListApi           = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.FormPrefix)
	AddApi            = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.FormPrefix)
	DetailApi         = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.FormPrefix)
	ModelApi          = flowablesdk.NewApi(httpclient.POST, modelUrl, flowablesdk.FormPrefix)
	DetailAndModelApi = flowablesdk.NewApi(httpclient.POST, detailAndModelUrl, flowablesdk.FormPrefix)
)
