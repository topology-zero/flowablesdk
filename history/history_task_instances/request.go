package history_task_instances

import (
	"time"

	"github.com/MasterJoyHunan/flowablesdk/common"
)

type ListRequest struct {
	common.ListCommonRequest
	TaskId                    string     `json:"taskId,omitempty"`
	ProcessInstanceId         string     `json:"processInstanceId,omitempty"`
	ProcessDefinitionKey      string     `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionId       string     `json:"processDefinitionId,omitempty"`
	ProcessDefinitionName     string     `json:"processDefinitionName,omitempty"`
	ProcessBusinessKey        string     `json:"processBusinessKey,omitempty"`
	ExecutionId               string     `json:"executionId,omitempty"`
	TaskDefinitionKey         string     `json:"taskDefinitionKey,omitempty"`
	TaskName                  string     `json:"taskName,omitempty"`
	TaskDescription           string     `json:"taskDescription,omitempty"`
	TaskCategory              string     `json:"taskCategory,omitempty"`
	TaskDeleteReason          string     `json:"taskDeleteReason,omitempty"`
	TaskAssignee              string     `json:"taskAssignee,omitempty"`
	TaskOwner                 string     `json:"taskOwner,omitempty"`
	TaskInvolvedUser          string     `json:"taskInvolvedUser,omitempty"`
	TaskPriority              string     `json:"taskPriority,omitempty"`
	Finished                  *bool      `json:"finished,omitempty"`
	ProcessFinished           *bool      `json:"processFinished,omitempty"`
	DueDateAfter              *time.Time `json:"dueDateAfter,omitempty"`
	DueDateBefore             *time.Time `json:"dueDateBefore,omitempty"`
	TaskCompletedAfter        *time.Time `json:"taskCompletedAfter,omitempty"`
	TaskCompletedBefore       *time.Time `json:"taskCompletedBefore,omitempty"`
	TaskCreatedBefore         *time.Time `json:"taskCreatedBefore,omitempty"`
	TaskCreatedAfter          *time.Time `json:"taskCreatedAfter,omitempty"`
	IncludeTaskLocalVariables *bool      `json:"includeTaskLocalVariables,omitempty"`
	IncludeProcessVariables   *bool      `json:"includeProcessVariables,omitempty"`
	TenantId                  string     `json:"tenantId,omitempty"`
}
