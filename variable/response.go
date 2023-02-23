package variable

type VariableResponse struct {
	Name          string `json:"name"`
	VariableScope string `json:"variableScope"`
	Value         any    `json:"value"`
}
