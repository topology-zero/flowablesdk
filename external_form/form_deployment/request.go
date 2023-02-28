package form_deployment

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest // order allow id,name,deployTime,tenantId
	Name                     string
	NameLike                 string
	Category                 string
	CategoryNotEquals        string
	ParentDeploymentId       string
	ParentDeploymentIdLike   string
	common.WithTenant
}

type AddRequest struct {
	FileName string // must end by .form
	Data     string
	TenantId string
	Category string
}
