package deployment

type ListResponse struct {
	Data  []ExternalFormDeployment `json:"data"`
	Total int                      `json:"total"`
	Start int                      `json:"start"`
	Size  int                      `json:"size"`
}
