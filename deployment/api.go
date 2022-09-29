package deployment

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl            = "/repository/deployments"
	detailUrl          = baseUrl + "/%s"
	ResourceURL        = detailUrl + "/resources"
	ResourceDetailURL  = ResourceURL + "/%s"
	ResourceContentURL = detailUrl + "/resourcedata/%s"
)

var (
	ListApi            = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi          = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	CreateApi          = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	DeleteApi          = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
	ResourceApi        = flowablesdk.NewApi(httpclient.GET, ResourceURL, flowablesdk.ProcessPrefix)
	ResourceDetailApi  = flowablesdk.NewApi(httpclient.GET, ResourceDetailURL, flowablesdk.ProcessPrefix)
	ResourceContentApi = flowablesdk.NewApi(httpclient.GET, ResourceContentURL, flowablesdk.ProcessPrefix)
)
