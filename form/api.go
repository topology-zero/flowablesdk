package form

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl = "/service/form/form-data"
)

var (
	DetailFromApi = flowablesdk.NewApi(httpclient.GET, baseUrl)
	SubmitFromApi = flowablesdk.NewApi(httpclient.POST, baseUrl)
)
