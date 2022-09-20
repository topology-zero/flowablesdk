package history

import "github.com/MasterJoyHunan/flowablesdk/variables"

type ListResponse struct {
	Data  []History `json:"data"`
	Total int       `json:"total"`
	Start int       `json:"start"`
	Size  int       `json:"size"`
}

type ListActivityResponse struct {
	Data  []ActivityHistory `json:"data"`
	Total int               `json:"total"`
	Start int               `json:"start"`
	Size  int               `json:"size"`
}

type ListVariableResponse struct {
	Data  []ListVariable `json:"data"`
	Total int            `json:"total"`
	Start int            `json:"start"`
	Size  int            `json:"size"`
}

type ListVariable struct {
	Id                 string                       `json:"id"`
	ProcessInstanceId  string                       `json:"processInstanceId"`
	ProcessInstanceUrl string                       `json:"processInstanceUrl"`
	TaskId             string                       `json:"taskId"`
	Variable           []variables.VariableResponse `json:"variable"`
}

type IdentityResponse struct {
	Type               string `json:"type"`
	UserId             string `json:"userId"`
	GroupId            string `json:"groupId"`
	TaskId             string `json:"taskId"`
	TaskUrl            string `json:"taskUrl"`
	ProcessInstanceId  string `json:"processInstanceId"`
	ProcessInstanceUrl string `json:"processInstanceUrl"`
}
