package common

type ListCommonRequest struct {
	Sort  string `json:"sort,omitempty"`
	Order string `json:"order,omitempty"`
	Start int    `json:"start,omitempty"`
	Size  int    `json:"size,omitempty"`
}
