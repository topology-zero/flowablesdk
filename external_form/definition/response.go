package definition

type ListResponse struct {
	Data  []ExternalFormDefinition `json:"data"`
	Total int                      `json:"total"`
	Start int                      `json:"start"`
	Size  int                      `json:"size"`
}
