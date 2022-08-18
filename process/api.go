package process

import (
	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

const (
	baseUrl             = "/service/repository/process-definitions"
	detailUrl           = baseUrl + "/%s"
	ResourceContentUrl  = detailUrl + "/resourcedata"
	ModelUrl            = detailUrl + "/model"
	ProcessCandidateUrl = detailUrl + "/identitylinks"
	CandidateProcessUrl = detailUrl + "/identitylinks/%s/%s"
)

var (
	ListApi                   = flowablesdk.NewApi(httpclient.GET, baseUrl)
	DetailApi                 = flowablesdk.NewApi(httpclient.GET, detailUrl)
	UpdateApi                 = flowablesdk.NewApi(httpclient.PUT, detailUrl)
	ResourceContentApi        = flowablesdk.NewApi(httpclient.GET, ResourceContentUrl)
	ModelApi                  = flowablesdk.NewApi(httpclient.GET, ModelUrl)
	ProcessCandidateApi       = flowablesdk.NewApi(httpclient.GET, ProcessCandidateUrl)
	ProcessAddCandidateApi    = flowablesdk.NewApi(httpclient.POST, ProcessCandidateUrl)
	ProcessDeleteCandidateApi = flowablesdk.NewApi(httpclient.DELETE, CandidateProcessUrl)
	CandidateProcessApi       = flowablesdk.NewApi(httpclient.GET, CandidateProcessUrl)
)
