package process_definition

import (
	"time"

	"github.com/topology-zero/flowablesdk/common"
)

type ListRequest struct {
	common.ListCommonRequest // order allow id,key,category,name,version,deploymentId,tenantId
	Version                  int
	Name                     string
	NameLike                 string
	Key                      string
	KeyLike                  string
	ResourceName             string
	ResourceNameLike         string
	Category                 string
	CategoryLike             string
	CategoryNotEquals        string
	DeploymentId             string
	StartableByUser          string
	Latest                   bool
	Suspended                bool
	common.WithTenant
}

type UpdateRequest struct {
	Action                  string     `json:"action,omitempty"` // suspend or activate
	Category                string     `json:"category,omitempty"`
	IncludeProcessInstances bool       `json:"includeProcessInstances,omitempty"`
	Date                    *time.Time `json:"date,omitempty"` //  2013-04-15T00:42:12Z
}

type AddCandidateRequest struct {
	User  string `json:"user,omitempty"`
	Group string `json:"group,omitempty"`
}

type DeleteCandidateRequest struct {
	Family      string // users or groups
	CandidateId string
}
