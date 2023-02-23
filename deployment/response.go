package deployment

type ResourceResponse struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	DataUrl   string `json:"dataUrl"`
	MediaType string `json:"mediaType"`
	Type      string `json:"type"`
}
