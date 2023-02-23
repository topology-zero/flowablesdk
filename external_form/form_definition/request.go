package form_definition

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest
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
