package process_instance

import (
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type ListRequest struct {
	common.ListCommonRequest
	Id                   string
	ProcessDefinitionKey string
	ProcessDefinitionId  string
	BusinessKey          string
	InvolvedUser         string
	Suspended            *bool
	TenantId             string
}

type UpdateRequest struct {
	Action string `json:"action"` // suspend | activate
}

type StartProcessRequest struct {
	ProcessDefinitionId  string                      `json:"processDefinitionId,omitempty"` // ProcessDefinitionId | ProcessDefinitionKey | Message 三选一
	ProcessDefinitionKey string                      `json:"processDefinitionKey,omitempty"`
	Message              string                      `json:"message,omitempty"`
	TenantId             string                      `json:"tenantId,omitempty"`
	BusinessKey          string                      `json:"businessKey"`
	ReturnVariables      bool                        `json:"returnVariables"`
	Variables            []variables.VariableRequest `json:"variables,omitempty"`
}
