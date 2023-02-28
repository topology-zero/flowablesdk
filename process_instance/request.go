package process_instance

import (
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest        // order allow id,processDefinitionId,tenantId,processDefinitionKey
	Id                       string `json:"id,omitempty"`
	ProcessDefinitionKey     string `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionId      string `json:"processDefinitionId,omitempty"`
	BusinessKey              string `json:"businessKey,omitempty"`
	InvolvedUser             string `json:"involvedUser,omitempty"`
	Suspended                *bool  `json:"suspended,omitempty"`
	SuperProcessInstanceId   string `json:"superProcessInstanceId,omitempty"`
	SubProcessInstanceId     string `json:"subProcessInstanceId,omitempty"`
	ExcludeSubprocesses      *bool  `json:"excludeSubprocesses,omitempty"`
	IncludeProcessVariables  *bool  `json:"includeProcessVariables,omitempty"`
	common.WithTenant
}

type UpdateRequest struct {
	Action string `json:"action"` // suspend | activate
}

type StartProcessRequest struct {
	ProcessDefinitionId  string                     `json:"processDefinitionId,omitempty"` // ProcessDefinitionId | ProcessDefinitionKey | Message 三选一
	ProcessDefinitionKey string                     `json:"processDefinitionKey,omitempty"`
	Message              string                     `json:"message,omitempty"`
	TenantId             string                     `json:"tenantId,omitempty"`
	BusinessKey          string                     `json:"businessKey"`
	ReturnVariables      bool                       `json:"returnVariables"`
	Variables            []variable.VariableRequest `json:"variable,omitempty"`
}
