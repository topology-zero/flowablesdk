package process_instance

import "io"

type ListRequest struct {
	Id                   string
	ProcessDefinitionKey string
	ProcessDefinitionId  string
	BusinessKey          string
	InvolvedUser         string
	Suspended            *bool
	TenantId             string
	Sort                 string
}

type UpdateRequest struct {
	Action string `json:"action"` // suspend | activate
}

type StartProcessRequest struct {
	ProcessDefinitionId  string     `json:"processDefinitionId,omitempty"` // ProcessDefinitionId | ProcessDefinitionKey | Message 三选一
	ProcessDefinitionKey string     `json:"processDefinitionKey,omitempty"`
	Message              string     `json:"message,omitempty"`
	TenantId             string     `json:"tenantId,omitempty"`
	BusinessKey          string     `json:"businessKey"`
	ReturnVariables      bool       `json:"returnVariables"`
	Variables            []Variable `json:"variables"`
}

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Candidate struct {
	User string `json:"user"`
	Type string `json:"type"`
}

type VariableRequest struct {
	Name  string `json:"name"`
	Type  string `json:"type"` // java的基础类型 integer | string | ......
	Value any    `json:"value"`
}

type FileVariableRequest struct {
	Name     string
	FileName string
	Value    io.Reader
}
