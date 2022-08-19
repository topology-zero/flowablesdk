package process_definition

type ListResponse struct {
	Data  []Process `json:"data"`
	Total int       `json:"total"`
	Start int       `json:"start"`
	Size  int       `json:"size"`
}
