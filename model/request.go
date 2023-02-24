package model

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest
	Id                string
	Category          string
	CategoryLike      string
	CategoryNotEquals string
	Name              string
	NameLike          string
	Key               string
	DeploymentId      string
	Version           string
	LatestVersion     *bool
	Deployed          *bool
	common.WithTenant
}

type AddRequest struct {
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	Category     string `json:"category,omitempty"`
	Version      int    `json:"version,omitempty"`
	MetaInfo     string `json:"metaInfo,omitempty"`
	DeploymentId string `json:"deploymentId,omitempty"`
	TenantId     string `json:"tenantId,omitempty"`
}
