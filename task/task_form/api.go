package task_form

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const baseUrl = "/runtime/tasks/%s/form"

var ListApi = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
