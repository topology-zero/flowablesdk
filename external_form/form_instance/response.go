package form_instance

import (
	"github.com/topology-zero/flowablesdk/external_form/model"
	"github.com/topology-zero/flowablesdk/pkg/timefmt"
)

type FormInstanceModelResponse struct {
	Id                  string        `json:"id"`
	Name                string        `json:"name"`
	Description         string        `json:"description"`
	Key                 string        `json:"key"`
	Version             int           `json:"version"`
	FormInstanceId      string        `json:"formInstanceId"`
	SubmittedBy         string        `json:"submittedBy"`
	SubmittedDate       *timefmt.Time `json:"submittedDate"`
	SelectedOutcome     string        `json:"selectedOutcome"`
	TaskId              string        `json:"taskId"`
	ProcessInstanceId   string        `json:"processInstanceId"`
	ProcessDefinitionId string        `json:"processDefinitionId"`
	TenantId            string        `json:"tenantId"`
	Url                 string        `json:"url"`
	model.Model
}
