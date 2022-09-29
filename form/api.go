package form

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl = "/form/form-data"
)

var (
	DetailFromApi = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	SubmitFromApi = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
)
