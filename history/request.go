package history

import (
	"time"

	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type ListRequest struct {
	common.ListCommonRequest
	ProcessInstanceId    string
	ProcessDefinitionKey string
	ProcessDefinitionId  string
	BusinessKey          string
	InvolvedUser         string
	Finished             *bool
	FinishedTime         []time.Time
	StartedTime          []time.Time
	StartedBy            string
	TenantId             string
}

type AddCommentRequest struct {
	Message               string `json:"message"`
	SaveProcessInstanceId bool   `json:"saveProcessInstanceId,omitempty"`
}

type ListActivityRequest struct {
	common.ListCommonRequest
	ActivityId          string `json:"activityId,omitempty"`
	ActivityInstanceId  string `json:"activityInstanceId,omitempty"`
	ActivityName        string `json:"activityName,omitempty"`
	ActivityType        string `json:"activityType,omitempty"`
	ExecutionId         string `json:"executionId,omitempty"`
	Finished            *bool  `json:"finished,omitempty"`
	TaskAssignee        string `json:"taskAssignee,omitempty"`
	ProcessInstanceId   string `json:"processInstanceId,omitempty"`
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	TenantId            string `json:"tenantId,omitempty"`
}

type ListVariableRequest struct {
	common.ListCommonRequest
	ProcessInstanceId    string                            `json:"processInstanceId,omitempty"`
	TaskId               string                            `json:"taskId,omitempty"`
	ExcludeTaskVariables bool                              `json:"excludeTaskVariables,omitempty"`
	VariableName         string                            `json:"variableName,omitempty"`
	VariableNameLike     string                            `json:"variableNameLike,omitempty"`
	Variables            []variables.VariableSearchRequest `json:"variables,omitempty"`
}
