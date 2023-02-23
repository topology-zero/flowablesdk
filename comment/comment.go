package comment

import "time"

type Comment struct {
	Id                 string    `json:"id"`
	TaskUrl            string    `json:"taskUrl"`
	ProcessInstanceUrl string    `json:"processInstanceUrl"`
	Message            string    `json:"message"`
	Author             string    `json:"author"`
	Time               time.Time `json:"time"`
	TaskId             string    `json:"taskId"`
	ProcessInstanceId  string    `json:"processInstanceId"`
}

type AddComment struct {
	Message               string `json:"message"`
	SaveProcessInstanceId bool   `json:"saveProcessInstanceId"`
	UserId                string `json:"-"`
}
