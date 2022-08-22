package events

type Events struct {
	Action  string   `json:"action"`
	Id      string   `json:"id"`
	Message []string `json:"message"`
	TaskUrl string   `json:"taskUrl"`
	Time    string   `json:"time"`
	Url     string   `json:"url"`
	UserId  string   `json:"userId"`
}
