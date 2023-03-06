package model_source

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl        = "/repository/models"
	detailUrl      = baseUrl + "/%s"
	sourceUrl      = detailUrl + "/source"
	extraSourceUrl = detailUrl + "/source-extra"
)

var (
	DetailSourceApi      = flowablesdk.NewApi(httpclient.GET, sourceUrl, flowablesdk.ProcessPrefix)
	UpdateSourceApi      = flowablesdk.NewApi(httpclient.PUT, sourceUrl, flowablesdk.ProcessPrefix)
	DetailExtraSourceApi = flowablesdk.NewApi(httpclient.GET, extraSourceUrl, flowablesdk.ProcessPrefix)
	UpdateExtraSourceApi = flowablesdk.NewApi(httpclient.PUT, extraSourceUrl, flowablesdk.ProcessPrefix)
)
