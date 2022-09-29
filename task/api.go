package task

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl               = "/runtime/tasks"
	detailUrl             = baseUrl + "/%s"
	variablesUrl          = detailUrl + "/variables"
	detailVariablesUrl    = variablesUrl + "/%s"
	binaryVariablesUrl    = detailVariablesUrl + "/data"
	identityUrl           = detailUrl + "/identitylinks"
	usersIdentityUrl      = identityUrl + "/users"
	groupsIdentityUrl     = identityUrl + "/groups"
	detailIdentityUrl     = identityUrl + "/%s/%s/%s"
	commentsUrl           = detailUrl + "/comments"
	detailCommentsUrl     = commentsUrl + "/%s"
	eventsUrl             = detailUrl + "/events"
	detailEventsUrl       = eventsUrl + "/%s"
	attachmentsUrl        = detailUrl + "/attachments"
	detailAttachmentsUrl  = attachmentsUrl + "/%s"
	contentAttachmentsUrl = detailAttachmentsUrl + "/content"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	UpdateApi = flowablesdk.NewApi(httpclient.PUT, detailUrl, flowablesdk.ProcessPrefix)
	ActionApi = flowablesdk.NewApi(httpclient.POST, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)

	ListVariablesApi   = flowablesdk.NewApi(httpclient.GET, variablesUrl, flowablesdk.ProcessPrefix)
	DetailVariableApi  = flowablesdk.NewApi(httpclient.GET, detailVariablesUrl, flowablesdk.ProcessPrefix)
	AddVariablesApi    = flowablesdk.NewApi(httpclient.POST, variablesUrl, flowablesdk.ProcessPrefix)
	UpdateVariableApi  = flowablesdk.NewApi(httpclient.PUT, detailVariablesUrl, flowablesdk.ProcessPrefix)
	DeleteVariablesApi = flowablesdk.NewApi(httpclient.DELETE, variablesUrl, flowablesdk.ProcessPrefix)
	DeleteVariableApi  = flowablesdk.NewApi(httpclient.DELETE, detailVariablesUrl, flowablesdk.ProcessPrefix)

	DetailBinaryVariableApi = flowablesdk.NewApi(httpclient.GET, binaryVariablesUrl, flowablesdk.ProcessPrefix)
	AddBinaryVariableApi    = flowablesdk.NewApi(httpclient.POST, variablesUrl, flowablesdk.ProcessPrefix)
	UpdateBinaryVariableApi = flowablesdk.NewApi(httpclient.PUT, detailVariablesUrl, flowablesdk.ProcessPrefix)

	ListIdentityApi       = flowablesdk.NewApi(httpclient.GET, identityUrl, flowablesdk.ProcessPrefix)
	ListUsersIdentityApi  = flowablesdk.NewApi(httpclient.GET, usersIdentityUrl, flowablesdk.ProcessPrefix)
	ListGroupsIdentityApi = flowablesdk.NewApi(httpclient.GET, groupsIdentityUrl, flowablesdk.ProcessPrefix)
	DetailIdentityApi     = flowablesdk.NewApi(httpclient.GET, detailIdentityUrl, flowablesdk.ProcessPrefix)
	AddIdentityApi        = flowablesdk.NewApi(httpclient.POST, identityUrl, flowablesdk.ProcessPrefix)
	DeleteIdentityApi     = flowablesdk.NewApi(httpclient.DELETE, detailIdentityUrl, flowablesdk.ProcessPrefix)

	ListCommentsApi   = flowablesdk.NewApi(httpclient.GET, commentsUrl, flowablesdk.ProcessPrefix)
	DetailCommentsApi = flowablesdk.NewApi(httpclient.GET, detailCommentsUrl, flowablesdk.ProcessPrefix)
	AddCommentsApi    = flowablesdk.NewApi(httpclient.POST, commentsUrl, flowablesdk.ProcessPrefix)
	DeleteCommentsApi = flowablesdk.NewApi(httpclient.DELETE, detailCommentsUrl, flowablesdk.ProcessPrefix)

	ListEventsApi   = flowablesdk.NewApi(httpclient.GET, eventsUrl, flowablesdk.ProcessPrefix)
	DetailEventsApi = flowablesdk.NewApi(httpclient.GET, detailEventsUrl, flowablesdk.ProcessPrefix)
	DeleteEventsApi = flowablesdk.NewApi(httpclient.GET, detailEventsUrl, flowablesdk.ProcessPrefix) // 源码中提供了删除接口,文档中没有提供,暂不实现该接口

	ListAttachmentsApi    = flowablesdk.NewApi(httpclient.GET, attachmentsUrl, flowablesdk.ProcessPrefix)
	AddAttachmentsApi     = flowablesdk.NewApi(httpclient.POST, attachmentsUrl, flowablesdk.ProcessPrefix)
	DeleteAttachmentsApi  = flowablesdk.NewApi(httpclient.DELETE, detailAttachmentsUrl, flowablesdk.ProcessPrefix) // 源码中提供了删除接口,文档中没有提供,已实现该接口
	ContentAttachmentsApi = flowablesdk.NewApi(httpclient.GET, contentAttachmentsUrl, flowablesdk.ProcessPrefix)   // 源码中提供了获取接口,文档中没有提供,已实现该接口
)
