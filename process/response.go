package process

type ListResponse struct {
	Data  []Detail `json:"data"`
	Total int      `json:"total"`
	Start int      `json:"start"`
	Size  int      `json:"size"`
}

type Detail struct {
	Id                       string `json:"id"`
	Url                      string `json:"url"`
	Version                  int    `json:"version"`
	Key                      string `json:"key"`
	Category                 string `json:"category"`
	Suspended                bool   `json:"suspended"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	DeploymentId             string `json:"deploymentId"`
	DeploymentUrl            string `json:"deploymentUrl"`
	GraphicalNotationDefined bool   `json:"graphicalNotationDefined"`
	Resource                 string `json:"resource"`
	DiagramResource          string `json:"diagramResource"`
	StartFormDefined         bool   `json:"startFormDefined"`
}

type Candidate struct {
	Url   string
	User  string
	Group string
	Type  string
}
