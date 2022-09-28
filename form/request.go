package form

type GetFromRequest struct {
	TaskId              string
	ProcessDefinitionId string
}

type SubmitFromRequest struct {
	TaskId              string       `json:"taskId,omitempty"`
	ProcessDefinitionId string       `json:"processDefinitionId,omitempty"`
	Properties          []Properties `json:"properties"`
}

type Properties struct {
	Id    string `json:"id"`
	Value any    `json:"value"`
}
