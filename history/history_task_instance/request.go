package history_task_instance

import (
	"time"

	"github.com/MasterJoyHunan/flowablesdk/common"
)

type ListRequest struct {
	common.ListCommonRequest             // order allow deleteReason,duration,endTime,executionId,taskInstanceId,processDefinitionId,processInstanceId,start,assignee,taskDefinitionKey,description,dueDate,name,owner,priority,tenantId,startTime
	TaskId                    string     `json:"taskId,omitempty"`
	ProcessInstanceId         string     `json:"processInstanceId,omitempty"`
	ProcessDefinitionKey      string     `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionKeyLike  string     `json:"processDefinitionKeyLike,omitempty"`
	ProcessDefinitionId       string     `json:"processDefinitionId,omitempty"`
	ProcessDefinitionName     string     `json:"processDefinitionName,omitempty"`
	ProcessDefinitionNameLike string     `json:"processDefinitionNameLike,omitempty"`
	ProcessBusinessKey        string     `json:"processBusinessKey,omitempty"`
	ProcessBusinessKeyLike    string     `json:"processBusinessKeyLike,omitempty"`
	ExecutionId               string     `json:"executionId,omitempty"`
	TaskDefinitionKey         string     `json:"taskDefinitionKey,omitempty"`
	TaskName                  string     `json:"taskName,omitempty"`
	TaskNameLike              string     `json:"taskNameLike,omitempty"`
	TaskDescription           string     `json:"taskDescription,omitempty"`
	TaskDescriptionLike       string     `json:"taskDescriptionLike,omitempty"`
	TaskCategory              string     `json:"taskCategory,omitempty"`
	TaskDeleteReason          string     `json:"taskDeleteReason,omitempty"`
	TaskDeleteReasonLike      string     `json:"taskDeleteReasonLike,omitempty"`
	TaskAssignee              string     `json:"taskAssignee,omitempty"`
	TaskAssigneeLike          string     `json:"taskAssigneeLike,omitempty"`
	TaskOwner                 string     `json:"taskOwner,omitempty"`
	TaskOwnerLike             string     `json:"taskOwnerLike,omitempty"`
	TaskInvolvedUser          string     `json:"taskInvolvedUser,omitempty"`
	TaskPriority              string     `json:"taskPriority,omitempty"`
	Finished                  *bool      `json:"finished,omitempty"`
	ProcessFinished           *bool      `json:"processFinished,omitempty"`
	ParentTaskId              string     `json:"parentTaskId,omitempty"`
	DueDate                   *time.Time `json:"dueDate,omitempty"`
	DueDateAfter              *time.Time `json:"dueDateAfter,omitempty"`
	DueDateBefore             *time.Time `json:"dueDateBefore,omitempty"`
	WithoutDueDate            *bool      `json:"withoutDueDate,omitempty"`
	TaskCompletedOn           *time.Time `json:"taskCompletedOn,omitempty"`
	TaskCompletedAfter        *time.Time `json:"taskCompletedAfter,omitempty"`
	TaskCompletedBefore       *time.Time `json:"taskCompletedBefore,omitempty"`
	TaskCreatedOn             *time.Time `json:"taskCreatedOn,omitempty"`
	TaskCreatedBefore         *time.Time `json:"taskCreatedBefore,omitempty"`
	TaskCreatedAfter          *time.Time `json:"taskCreatedAfter,omitempty"`
	IncludeTaskLocalVariables *bool      `json:"includeTaskLocalVariables,omitempty"`
	IncludeProcessVariables   *bool      `json:"includeProcessVariables,omitempty"`
	common.WithTenant
}
