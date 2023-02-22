package history_process_instances

type AddRequest struct {
	Message               string `json:"message"`
	SaveProcessInstanceId *bool  `json:"saveProcessInstanceId,omitempty"`
}
