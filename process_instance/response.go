package process_instance

type ListResponse struct {
	Data  []Instance `json:"data"`
	Total int        `json:"total"`
	Start int        `json:"start"`
	Size  int        `json:"size"`
}

type VariableResponse struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    any    `json:"value"`
	Scope    string `json:"scope"`
	ValueUrl string `json:"valueUrl"`
}
