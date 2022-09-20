package history

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl          = "/service/history/historic-process-instances"
	detailUrl        = baseUrl + "/%s"
	identityUrl      = detailUrl + "/identitylinks"
	commentUrl       = detailUrl + "/comments"
	detailCommentUrl = commentUrl + "/%s"

	activityUrl = "/service/query/historic-activity-instances"

	variableUrl = "/service/query/historic-variable-instances"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.GET, baseUrl)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl)
	DeleteApi       = flowablesdk.NewApi(httpclient.DELETE, detailUrl)
	ListIdentityApi = flowablesdk.NewApi(httpclient.GET, identityUrl)

	ListCommentApi   = flowablesdk.NewApi(httpclient.GET, commentUrl)
	AddCommentApi    = flowablesdk.NewApi(httpclient.POST, commentUrl)
	DetailCommentApi = flowablesdk.NewApi(httpclient.GET, detailCommentUrl)
	DeleteCommentApi = flowablesdk.NewApi(httpclient.DELETE, detailCommentUrl)

	ListActivityApi = flowablesdk.NewApi(httpclient.POST, activityUrl)

	ListVariableApi = flowablesdk.NewApi(httpclient.POST, variableUrl)
)
