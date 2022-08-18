package process

type ListRequest struct {
	Version         int
	Name            string
	Key             string
	ResourceName    string
	Category        string
	DeploymentId    string
	StartableByUser string
	Latest          bool
	Suspended       bool
	Sort            string
}

type UpdateRequest struct {
	Category                string
	Action                  string // suspend or activate
	IncludeProcessInstances bool
	Date                    string //  2013-04-15T00:42:12Z
}

type ProcessAddCandidateRequest struct {
	User    string `json:"user,omitempty"`
	GroupId string `json:"groupId,omitempty"`
}

type ProcessDeleteCandidateRequest struct {
	Family     string // users or groups
	IdentityId string
}
