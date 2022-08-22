package task

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl               = "/service/runtime/tasks"
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
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl)
	UpdateApi = flowablesdk.NewApi(httpclient.PUT, detailUrl)
	ActionApi = flowablesdk.NewApi(httpclient.POST, detailUrl)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl)

	ListVariablesApi   = flowablesdk.NewApi(httpclient.GET, variablesUrl)
	DetailVariableApi  = flowablesdk.NewApi(httpclient.GET, detailVariablesUrl)
	AddVariablesApi    = flowablesdk.NewApi(httpclient.POST, variablesUrl)
	UpdateVariableApi  = flowablesdk.NewApi(httpclient.PUT, detailVariablesUrl)
	DeleteVariablesApi = flowablesdk.NewApi(httpclient.DELETE, variablesUrl)
	DeleteVariableApi  = flowablesdk.NewApi(httpclient.DELETE, detailVariablesUrl)

	DetailBinaryVariableApi = flowablesdk.NewApi(httpclient.GET, binaryVariablesUrl)
	AddBinaryVariableApi    = flowablesdk.NewApi(httpclient.POST, variablesUrl)
	UpdateBinaryVariableApi = flowablesdk.NewApi(httpclient.PUT, detailVariablesUrl)

	ListIdentityApi       = flowablesdk.NewApi(httpclient.GET, identityUrl)
	ListUsersIdentityApi  = flowablesdk.NewApi(httpclient.GET, usersIdentityUrl)
	ListGroupsIdentityApi = flowablesdk.NewApi(httpclient.GET, groupsIdentityUrl)
	DetailIdentityApi     = flowablesdk.NewApi(httpclient.GET, detailIdentityUrl)
	AddIdentityApi        = flowablesdk.NewApi(httpclient.POST, identityUrl)
	DeleteIdentityApi     = flowablesdk.NewApi(httpclient.DELETE, detailIdentityUrl)

	ListCommentsApi   = flowablesdk.NewApi(httpclient.GET, commentsUrl)
	DetailCommentsApi = flowablesdk.NewApi(httpclient.GET, detailCommentsUrl)
	AddCommentsApi    = flowablesdk.NewApi(httpclient.POST, commentsUrl)
	DeleteCommentsApi = flowablesdk.NewApi(httpclient.DELETE, detailCommentsUrl)

	ListEventsApi   = flowablesdk.NewApi(httpclient.GET, eventsUrl)
	DetailEventsApi = flowablesdk.NewApi(httpclient.GET, detailEventsUrl)
	DeleteEventsApi = flowablesdk.NewApi(httpclient.GET, detailEventsUrl) // 源码中提供了删除接口,文档中没有提供,暂不实现该接口

	ListAttachmentsApi    = flowablesdk.NewApi(httpclient.GET, attachmentsUrl)
	AddAttachmentsApi     = flowablesdk.NewApi(httpclient.POST, attachmentsUrl)
	DeleteAttachmentsApi  = flowablesdk.NewApi(httpclient.DELETE, detailAttachmentsUrl) // 源码中提供了删除接口,文档中没有提供,已实现该接口
	ContentAttachmentsApi = flowablesdk.NewApi(httpclient.GET, contentAttachmentsUrl)   // 源码中提供了删除接口,文档中没有提供,已实现该接口
)
