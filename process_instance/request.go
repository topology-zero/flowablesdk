package process_instance

import (
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/variable"
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
	ProcessDefinitionId        string                     `json:"processDefinitionId,omitempty"`
	ProcessDefinitionKey       string                     `json:"processDefinitionKey,omitempty"`
	Message                    string                     `json:"message,omitempty"`
	Name                       string                     `json:"name,omitempty"`
	BusinessKey                string                     `json:"businessKey,omitempty"`
	Variables                  []variable.VariableRequest `json:"variables,omitempty"`
	TransientVariables         []variable.VariableRequest `json:"transientVariables,omitempty"`
	StartFormVariables         []variable.VariableRequest `json:"startFormVariables,omitempty"`
	TenantId                   string                     `json:"tenantId,omitempty"`
	Outcome                    string                     `json:"outcome,omitempty"`
	OverrideDefinitionTenantId string                     `json:"overrideDefinitionTenantId,omitempty"`
}
