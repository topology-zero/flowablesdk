package execution

import (
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest // order allow processDefinitionId,processDefinitionKey,processInstanceId,tenantId
	common.WithTenant
	Id                           string                           `json:"id,omitempty"`
	ActivityId                   string                           `json:"activityId,omitempty"`
	ProcessDefinitionKey         string                           `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionId          string                           `json:"processDefinitionId,omitempty"`
	ProcessInstanceId            string                           `json:"processInstanceId,omitempty"`
	MessageEventSubscriptionName string                           `json:"messageEventSubscriptionName,omitempty"`
	SignalEventSubscriptionName  string                           `json:"signalEventSubscriptionName,omitempty"`
	ParentId                     string                           `json:"parentId,omitempty"`
	Variables                    []variable.VariableSearchRequest `json:"variable,omitempty"`
	ProcessInstanceVariables     []variable.VariableSearchRequest `json:"processInstanceVariables,omitempty"`
}

type ExecuteRequest struct {
	Action      string                    `json:"action,omitempty"` // signal,messageEventReceived
	MessageName string                    `json:"messageName,omitempty"`
	Variables   []variable.SimpleVariable `json:"variables,omitempty"`
}
