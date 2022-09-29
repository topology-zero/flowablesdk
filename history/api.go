package history

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl          = "/history/historic-process-instances"
	detailUrl        = baseUrl + "/%s"
	identityUrl      = detailUrl + "/identitylinks"
	commentUrl       = detailUrl + "/comments"
	detailCommentUrl = commentUrl + "/%s"

	activityUrl = "/query/historic-activity-instances"

	variableUrl = "/query/historic-variable-instances"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi       = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
	ListIdentityApi = flowablesdk.NewApi(httpclient.GET, identityUrl, flowablesdk.ProcessPrefix)

	ListCommentApi   = flowablesdk.NewApi(httpclient.GET, commentUrl, flowablesdk.ProcessPrefix)
	AddCommentApi    = flowablesdk.NewApi(httpclient.POST, commentUrl, flowablesdk.ProcessPrefix)
	DetailCommentApi = flowablesdk.NewApi(httpclient.GET, detailCommentUrl, flowablesdk.ProcessPrefix)
	DeleteCommentApi = flowablesdk.NewApi(httpclient.DELETE, detailCommentUrl, flowablesdk.ProcessPrefix)

	ListActivityApi = flowablesdk.NewApi(httpclient.POST, activityUrl, flowablesdk.ProcessPrefix)

	ListVariableApi = flowablesdk.NewApi(httpclient.POST, variableUrl, flowablesdk.ProcessPrefix)
)
