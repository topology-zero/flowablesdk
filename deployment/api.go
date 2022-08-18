package deployment

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl            = "/service/repository/deployments"
	detailUrl          = baseUrl + "/%s"
	ResourceURL        = detailUrl + "/resources"
	ResourceDetailURL  = ResourceURL + "/%s"
	ResourceContentURL = detailUrl + "/resourcedata/%s"
)

var (
	ListApi            = flowablesdk.NewApi(httpclient.GET, baseUrl)
	DetailApi          = flowablesdk.NewApi(httpclient.GET, detailUrl)
	CreateApi          = flowablesdk.NewApi(httpclient.POST, baseUrl)
	DeleteApi          = flowablesdk.NewApi(httpclient.DELETE, detailUrl)
	ResourceApi        = flowablesdk.NewApi(httpclient.GET, ResourceURL)
	ResourceDetailApi  = flowablesdk.NewApi(httpclient.GET, ResourceDetailURL)
	ResourceContentApi = flowablesdk.NewApi(httpclient.GET, ResourceContentURL)
)
