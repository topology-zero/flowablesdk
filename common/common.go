package common

import "encoding/json"

type ListCommonRequest struct {
	Sort  string `json:"sort,omitempty"`
	Order string `json:"order,omitempty"`
	Start int    `json:"start,omitempty"`
	Size  int    `json:"size,omitempty"`
}

type ListCommonResponse struct {
	Data  json.RawMessage `json:"data"`
	Total int             `json:"total"`
	Start int             `json:"start"`
	Size  int             `json:"size"`
}
