package task

type ListResponse struct {
	Data  []Tasks `json:"data"`
	Total int     `json:"total"`
	Start int     `json:"start"`
	Size  int     `json:"size"`
}
