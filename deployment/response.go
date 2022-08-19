package deployment

type ListResponse struct {
	Data  []Deployment `json:"data"`
	Total int          `json:"total"`
	Start int          `json:"start"`
	Size  int          `json:"size"`
}

type Resource struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	DataUrl   string `json:"dataUrl"`
	MediaType string `json:"mediaType"`
	Type      string `json:"type"`
}
