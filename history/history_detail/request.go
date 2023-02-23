package history_detail

import (
	"github.com/MasterJoyHunan/flowablesdk/common"
)

type ListRequest struct {
	common.ListCommonRequest
	Id                        string `json:"id,omitempty"`
	ProcessInstanceId         string `json:"processInstanceId,omitempty"`
	ExecutionId               string `json:"executionId,omitempty"`
	ActivityInstanceId        string `json:"activityInstanceId,omitempty"`
	TaskId                    string `json:"taskId,omitempty"`
	SelectOnlyFormProperties  *bool  `json:"selectOnlyFormProperties,omitempty"`
	SelectOnlyVariableUpdates *bool  `json:"selectOnlyVariableUpdates,omitempty"`
}
