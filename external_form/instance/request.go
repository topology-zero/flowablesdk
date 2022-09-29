package instance

import "github.com/MasterJoyHunan/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest
	Id                  string `json:"id,omitempty"`
	FormDefinitionId    string `json:"formDefinitionId,omitempty"`
	TaskId              string `json:"taskId,omitempty"`
	ProcessInstanceId   string `json:"processInstanceId,omitempty"`
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	ScopeId             string `json:"scopeId,omitempty"`
	ScopeType           string `json:"scopeType,omitempty"`
	ScopeDefinitionId   string `json:"scopeDefinitionId,omitempty"`
	SubmittedBy         string `json:"submittedBy,omitempty"`
	TenantId            string `json:"tenantId,omitempty"`
}

type AddRequest struct {
	FormInstanceId          string         `json:"formInstanceId,omitempty"`
	FormDefinitionId        string         `json:"formDefinitionId,omitempty"`
	FormDefinitionKey       string         `json:"formDefinitionKey,omitempty"`
	TaskId                  string         `json:"taskId,omitempty"`
	ProcessInstanceId       string         `json:"processInstanceId,omitempty"`
	ProcessDefinitionId     string         `json:"processDefinitionId,omitempty"`
	ScopeId                 string         `json:"scopeId,omitempty"`
	ScopeType               string         `json:"scopeType,omitempty"`
	ScopeDefinitionId       string         `json:"scopeDefinitionId,omitempty"`
	TenantId                string         `json:"tenantId,omitempty"`
	ParentDeploymentId      string         `json:"parentDeploymentId,omitempty"`
	Variables               map[string]any `json:"variables,omitempty"`
	FallbackToDefaultTenant *bool          `json:"fallbackToDefaultTenant,omitempty"`
}
