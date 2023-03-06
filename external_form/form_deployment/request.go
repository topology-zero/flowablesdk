package form_deployment

import (
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/external_form/model"
)

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
	Data     model.Model
	TenantId string
	Category string
}
