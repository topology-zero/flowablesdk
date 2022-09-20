package task

import (
	"io"
	"time"

	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

type ListRequest struct {
	common.ListCommonRequest
	Name                       string      // 任务名称
	Description                string      // 任务描述
	Priority                   []int       // 优先级区间
	Assignee                   string      // 分配给该用户的任务
	Owner                      string      // 该用户拥有的任务
	CandidateUser              string      // 任务有该候选者,或有该候选者所在的组
	CandidateGroups            []string    // 含有该组的任务
	CandidateOrAssigned        string      // candidate user or groups
	InvolvedUser               string      // 仅返回用户参与的任务
	Unassigned                 bool        // 为 true 返回没有指定候选人的任务
	DelegationState            string      // pending and resolved ...
	TaskDefinitionKey          string      // 任务定义ID
	ProcessInstanceId          string      // 流程实例ID
	ProcessInstanceBusinessKey string      // 流程实例KEY
	ProcessDefinitionId        string      // 流程定义ID
	ProcessDefinitionKey       string      // 流程定义KEY
	ProcessDefinitionName      string      // 流程定义名称
	CreateTime                 []time.Time // 任务创建时间区间
	Active                     *bool       // true 返回未挂起的任务, false 返回挂起的任务
	Category                   string
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
	UserId                string `json:"-"`
	Message               string `json:"message"`
	SaveProcessInstanceId bool   `json:"saveProcessInstanceId,omitempty"`
}

type AddAttachmentsUrlRequest struct {
	UserId      string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	ExternalUrl string `json:"externalUrl,omitempty"`
}

type AddAttachmentsRequest struct {
	UserId      string
	Name        string
	Description string
	Type        string
	File        io.Reader
}
