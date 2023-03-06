package history_activity_instance

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	queryUrl = "/query/historic-activity-instances"
)

var (
	ListApi = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
)
