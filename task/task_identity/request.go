package task_identity

type DetailRequest struct {
	Family     string // groups or users
	IdentityId string
	Type       string
}

type AddRequest struct {
	Group string `json:"group,omitempty"`
	User  string `json:"user,omitempty"`
	Type  string `json:"type"`
}
