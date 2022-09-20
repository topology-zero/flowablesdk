package attachment

import "time"

type Attachment struct {
	Id                 string    `json:"id"`
	Url                string    `json:"url"`
	Name               string    `json:"name"`
	UserId             string    `json:"userId"`
	Description        string    `json:"description"`
	Type               string    `json:"type"`
	TaskUrl            string    `json:"taskUrl"`
	ProcessInstanceUrl string    `json:"processInstanceUrl"`
	ExternalUrl        string    `json:"externalUrl"`
	ContentUrl         string    `json:"contentUrl"`
	Time               time.Time `json:"time"`
}
