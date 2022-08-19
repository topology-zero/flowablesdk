package process_definition

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl            = "/service/repository/process-definitions"
	detailUrl          = baseUrl + "/%s"
	ResourceContentUrl = detailUrl + "/resourcedata"
	ModelUrl           = detailUrl + "/model"
	CandidateUrl       = detailUrl + "/identitylinks"
	CandidateDetailUrl = CandidateUrl + "/%s/%s"
)

var (
	ListApi            = flowablesdk.NewApi(httpclient.GET, baseUrl)
	DetailApi          = flowablesdk.NewApi(httpclient.GET, detailUrl)
	UpdateApi          = flowablesdk.NewApi(httpclient.PUT, detailUrl)
	ResourceContentApi = flowablesdk.NewApi(httpclient.GET, ResourceContentUrl)
	ModelApi           = flowablesdk.NewApi(httpclient.GET, ModelUrl)
	ListCandidateApi   = flowablesdk.NewApi(httpclient.GET, CandidateUrl)
	AddCandidateApi    = flowablesdk.NewApi(httpclient.POST, CandidateUrl)
	DeleteCandidateApi = flowablesdk.NewApi(httpclient.DELETE, CandidateDetailUrl)
	CandidateDetailApi = flowablesdk.NewApi(httpclient.GET, CandidateDetailUrl)
)
