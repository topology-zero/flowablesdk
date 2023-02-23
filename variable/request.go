package variable

import "io"

type VariableRequest struct {
	Name  string `json:"name"`
	Type  string `json:"type"`            // java的基础类型 integer | string | ......
	Scope string `json:"scope,omitempty"` // task 使用
	Value any    `json:"value,omitempty"`
}

type FileVariableRequest struct {
	VariableName string
	FileName     string
	Scope        string // task 使用
	Value        io.Reader
}

type VariableSearchRequest struct {
	Name      string `json:"name"`
	Value     any    `json:"value"`
	Operation string `json:"operation"`
	Type      string `json:"type"`
}

type SimpleVariable struct {
	Name  string `json:"name,omitempty"`
	Value any    `json:"value,omitempty"`
}
