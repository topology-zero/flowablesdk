package deployment

type ListResponse struct {
	Data  []Detail `json:"data"`
	Total int      `json:"total"`
	Start int      `json:"start"`
	Size  int      `json:"size"`
}

type Detail struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	DeploymentTime string `json:"deploymentTime"`
	Category       string `json:"category"`
	Url            string `json:"url"`
	TenantId       string `json:"tenantId"`
}

type Resource struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	DataUrl   string `json:"dataUrl"`
	MediaType string `json:"mediaType"`
	Type      string `json:"type"`
}
