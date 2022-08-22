package task

import (
	"io"

	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type ListRequest struct {
	Name                  string
	Description           string
	Assignee              string   // 分配给该用户的任务
	Owner                 string   // 该用户拥有的任务
	CandidateUser         string   // 任务有该候选者,或有该候选者所在的组
	CandidateGroups       []string // 含有该组的任务
	InvolvedUser          string   // 仅返回用户参与的任务
	TaskDefinitionKey     string
	ProcessDefinitionKey  string
	ProcessDefinitionName string
	StartTime             string
	EndTime               string
	Active                *bool // true 返回未挂起的任务, false 返回挂起的任务
	Category              string
	Sort                  string
	Order                 string
}

type UpdateRequest struct {
	Assignee        string `json:"assignee,omitempty"`        // 转让
	DelegationState string `json:"delegationState,omitempty"` // 状态
	Description     string `json:"description,omitempty"`
	DueDate         string `json:"dueDate,omitempty"` // 过期时间
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
	CascadeHistory bool   // 是否删除历史记录
	DeleteReason   string // CascadeHistory 为 true 时,该值会被忽略
}

type DetailIdentityRequest struct {
	Family     string
	IdentityId string
	Type       string
}

type AddIdentityRequest struct {
	User  string `json:"user,omitempty"`
	Group string `json:"group,omitempty"`
	Type  string `json:"type"`
}

type AddCommentsRequest struct {
	Message               string `json:"message"`
	SaveProcessInstanceId bool   `json:"saveProcessInstanceId,omitempty"`
}

type AddAttachmentsUrlRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	ExternalUrl string `json:"externalUrl,omitempty"`
}

type AddAttachmentsRequest struct {
	Name        string
	Description string
	Type        string
	File        io.Reader
}
