package form_definitions

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/form-repository/form-definitions"
	detailUrl = baseUrl + "/%s"
	modelUrl  = detailUrl + "/model"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.FormPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.FormPrefix)
	ModelApi  = flowablesdk.NewApi(httpclient.GET, modelUrl, flowablesdk.FormPrefix)
)
