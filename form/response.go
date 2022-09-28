package form

type SubmitFromResponse struct {
	Id                   string `json:"id"`
	Url                  string `json:"url"`
	BusinessKey          string `json:"businessKey"`
	Suspended            bool   `json:"suspended"`
	ProcessDefinitionId  string `json:"processDefinitionId"`
	ProcessDefinitionUrl string `json:"processDefinitionUrl"`
	ActivityId           string `json:"activityId"`
}
