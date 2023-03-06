package form_definition

import "github.com/topology-zero/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest   // order allow id,key,category,name,version,deploymentId,tenantId
	Category                   string
	CategoryLike               string
	CategoryNotEquals          string
	Key                        string
	KeyLike                    string
	Name                       string
	NameLike                   string
	ResourceName               string
	ResourceNameLike           string
	Version                    int
	VersionGreaterThan         int
	VersionGreaterThanOrEquals int
	VersionLowerThan           int
	VersionLowerThanOrEquals   int
	DeploymentId               string
	Latest                     bool
	common.WithTenant
}
