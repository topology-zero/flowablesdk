package model

import (
	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/repository/models"
	detailUrl = baseUrl + "/%s"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	AddApi    = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
	UpdateApi = flowablesdk.NewApi(httpclient.PUT, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
)
