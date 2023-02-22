package history_task_instances

type IdentityResponse struct {
	Type               string `json:"type"`
	UserId             string `json:"userId"`
	GroupId            string `json:"groupId"`
	TaskId             string `json:"taskId"`
	TaskUrl            string `json:"taskUrl"`
	ProcessInstanceId  string `json:"processInstanceId"`
	ProcessInstanceUrl string `json:"processInstanceUrl"`
}
