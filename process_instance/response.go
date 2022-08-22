package process_instance

type ListResponse struct {
	Data  []Instance `json:"data"`
	Total int        `json:"total"`
	Start int        `json:"start"`
	Size  int        `json:"size"`
}
