package definition

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest
	Category     string
	Key          string
	Name         string
	ResourceName string
	Version      string
	DeploymentId string
	TenantId     string
	Latest       bool
}
