package deployment

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest
	Name     string
	Category string
	TenantId string
}

type AddRequest struct {
	FileName string // must end by .form
	Data     string
	TenantId string
	Category string
}
