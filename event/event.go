package event

import "time"

type Event struct {
	Action  string    `json:"action"`
	Id      string    `json:"id"`
	Message []string  `json:"message"`
	TaskUrl string    `json:"taskUrl"`
	Time    time.Time `json:"time"`
	Url     string    `json:"url"`
	UserId  string    `json:"userId"`
}
