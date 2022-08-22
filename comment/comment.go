package comment

type Comment struct {
	Id                 string `json:"id"`
	TaskUrl            string `json:"taskUrl"`
	ProcessInstanceUrl string `json:"processInstanceUrl"`
	Message            string `json:"message"`
	Author             string `json:"author"`
	Time               string `json:"time"`
	TaskId             string `json:"taskId"`
	ProcessInstanceId  string `json:"processInstanceId"`
}
