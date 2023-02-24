package history_variable_instance

import (
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest
	common.WithTenant
	ProcessInstanceId    string                           `json:"processInstanceId,omitempty"`
	TaskId               string                           `json:"taskId,omitempty"`
	ExcludeTaskVariables string                           `json:"excludeTaskVariables,omitempty"`
	VariableName         string                           `json:"variableName,omitempty"`
	VariableNameLike     string                           `json:"variableNameLike,omitempty"`
	Variables            []variable.VariableSearchRequest `json:"variables,omitempty"`
}
