package deployment

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest
	Name     string
	Category string
}

type CreateRequest struct {
	Category string
	FileName string // must end by .bpmn20.xml
	Xml      string
}
