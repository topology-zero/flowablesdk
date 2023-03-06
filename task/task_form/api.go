package task_form

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const baseUrl = "/runtime/tasks/%s/form"

var ListApi = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
