package form_definition

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/form-repository/form-definitions"
	detailUrl = baseUrl + "/%s"
	modelUrl  = detailUrl + "/models"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.FormPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.FormPrefix)
	ModelApi  = flowablesdk.NewApi(httpclient.GET, modelUrl, flowablesdk.FormPrefix)
)
