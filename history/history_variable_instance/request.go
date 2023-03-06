package history_variable_instance

import (
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest // order allow processInstanceId,variableName
	common.WithTenant
	ProcessInstanceId    string                           `json:"processInstanceId,omitempty"`
	TaskId               string                           `json:"taskId,omitempty"`
	ExcludeTaskVariables string                           `json:"excludeTaskVariables,omitempty"`
	VariableName         string                           `json:"variableName,omitempty"`
	VariableNameLike     string                           `json:"variableNameLike,omitempty"`
	Variables            []variable.VariableSearchRequest `json:"variables,omitempty"`
}
