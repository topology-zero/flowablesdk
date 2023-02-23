package task

import (
	"time"

	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type ListRequest struct {
	common.ListCommonRequest
	Name                           string                      `json:"name,omitempty"`
	NameLike                       string                      `json:"nameLike,omitempty"`
	Description                    string                      `json:"description,omitempty"`
	Priority                       int                         `json:"priority,omitempty"`
	MinimumPriority                int                         `json:"minimumPriority,omitempty"`
	MaximumPriority                int                         `json:"maximumPriority,omitempty"`
	Assignee                       string                      `json:"assignee,omitempty"`
	AssigneeLike                   string                      `json:"assigneeLike,omitempty"`
	Owner                          string                      `json:"owner,omitempty"`
	OwnerLike                      string                      `json:"ownerLike,omitempty"`
	Unassigned                     *bool                       `json:"unassigned,omitempty"`
	DelegationState                string                      `json:"delegationState,omitempty"`
	CandidateUser                  string                      `json:"candidateUser,omitempty"`
	CandidateGroup                 string                      `json:"candidateGroup,omitempty"`
	CandidateGroups                string                      `json:"candidateGroups,omitempty"`
	InvolvedUser                   string                      `json:"involvedUser,omitempty"`
	TaskDefinitionKey              string                      `json:"taskDefinitionKey,omitempty"`
	TaskDefinitionKeyLike          string                      `json:"taskDefinitionKeyLike,omitempty"`
	ProcessInstanceId              string                      `json:"processInstanceId,omitempty"`
	ProcessInstanceBusinessKey     string                      `json:"processInstanceBusinessKey,omitempty"`
	ProcessInstanceBusinessKeyLike string                      `json:"processInstanceBusinessKeyLike,omitempty"`
	ProcessDefinitionId            string                      `json:"processDefinitionId,omitempty"`
	ProcessDefinitionKey           string                      `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionKeyLike       string                      `json:"processDefinitionKeyLike,omitempty"`
	ProcessDefinitionName          string                      `json:"processDefinitionName,omitempty"`
	ProcessDefinitionNameLike      string                      `json:"processDefinitionNameLike,omitempty"`
	ExecutionId                    string                      `json:"executionId,omitempty"`
	CreatedOn                      *time.Time                  `json:"createdOn,omitempty"`
	CreatedBefore                  *time.Time                  `json:"createdBefore,omitempty"`
	CreatedAfter                   *time.Time                  `json:"createdAfter,omitempty"`
	DueOn                          *time.Time                  `json:"dueOn,omitempty"`
	DueBefore                      *time.Time                  `json:"dueBefore,omitempty"`
	DueAfter                       *time.Time                  `json:"dueAfter,omitempty"`
	WithoutDueDate                 *time.Time                  `json:"withoutDueDate,omitempty"`
	ExcludeSubTasks                *bool                       `json:"excludeSubTasks,omitempty"`
	Active                         *bool                       `json:"active,omitempty"`
	IncludeTaskLocalVariables      *bool                       `json:"includeTaskLocalVariables,omitempty"`
	IncludeProcessVariables        *bool                       `json:"includeProcessVariables,omitempty"`
	CandidateOrAssigned            string                      `json:"candidateOrAssigned,omitempty"`
	Category                       string                      `json:"category,omitempty"`
	TaskVariables                  []variables.VariableRequest `json:"taskVariables,omitempty"`
	ProcessInstanceVariables       []variables.VariableRequest `json:"processInstanceVariables,omitempty"`
	common.WithTenant
}

type UpdateRequest struct {
	Assignee        string `json:"assignee,omitempty"`
	DelegationState string `json:"delegationState,omitempty"`
	Description     string `json:"description,omitempty"`
	DueDate         string `json:"dueDate,omitempty"`
	Name            string `json:"name,omitempty"`
	Owner           string `json:"owner,omitempty"`
	ParentTaskId    string `json:"parentTaskId,omitempty"`
	Priority        int    `json:"priority,omitempty"`
}

type ActionRequest struct {
	Action    string                      `json:"action"`
	Assignee  string                      `json:"assignee,omitempty"`
	Variables []variables.VariableRequest `json:"variables,omitempty"`
}

type DeleteRequest struct {
	CascadeHistory bool
	DeleteReason   string
}
