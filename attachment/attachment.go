package attachment

type Attachment struct {
	Id                 string `json:"id"`
	Url                string `json:"url"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Type               string `json:"type"`
	TaskUrl            string `json:"taskUrl"`
	ProcessInstanceUrl string `json:"processInstanceUrl"`
	ExternalUrl        string `json:"externalUrl"`
	ContentUrl         string `json:"contentUrl"`
}
