package history_process_instances

import (
	"time"

	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type ListRequest struct {
	common.ListCommonRequest
	ProcessInstanceId       string                             `json:"processInstanceId,omitempty"`
	ProcessDefinitionKey    string                             `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionId     string                             `json:"processDefinitionId,omitempty"`
	BusinessKey             string                             `json:"businessKey,omitempty"`
	InvolvedUser            string                             `json:"involvedUser,omitempty"`
	Finished                *bool                              `json:"finished,omitempty"`
	FinishedAfter           *time.Time                         `json:"finishedAfter,omitempty"`
	FinishedBefore          *time.Time                         `json:"finishedBefore,omitempty"`
	StartedAfter            *time.Time                         `json:"startedAfter,omitempty"`
	StartedBefore           *time.Time                         `json:"startedBefore,omitempty"`
	StartedBy               string                             `json:"startedBy,omitempty"`
	IncludeProcessVariables *bool                              `json:"includeProcessVariables,omitempty"`
	TenantId                string                             `json:"tenantId,omitempty"`
	Variables               *[]variables.VariableSearchRequest `json:"variables,omitempty"`
}
