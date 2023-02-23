package task_identity

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/runtime/tasks/%s/identitylinks"
	usersUrl  = baseUrl + "/users"
	groupsUrl = baseUrl + "/groups"
	detailUrl = baseUrl + "/%s/%s/%s"
)

var (
	ListApi       = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	ListUsersApi  = flowablesdk.NewApi(httpclient.GET, usersUrl, flowablesdk.ProcessPrefix)
	ListGroupsApi = flowablesdk.NewApi(httpclient.GET, groupsUrl, flowablesdk.ProcessPrefix)
	DetailApi     = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	AddApi        = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	DeleteApi     = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
)
