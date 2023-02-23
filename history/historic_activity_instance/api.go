package historic_activity_instance

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	queryUrl = "/query/historic-task-instances"
)

var (
	ListApi = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
)
