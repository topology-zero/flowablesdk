package common

import "encoding/json"

type ListCommonRequest struct {
	Sort  string `json:"sort,omitempty"`
	Order string `json:"order,omitempty"`
	Start int    `json:"start,omitempty"`
	Size  int    `json:"size,omitempty"`
}

type WithTenant struct {
	TenantId        string `json:"tenantId,omitempty"`
	TenantIdLike    string `json:"tenantIdLike,omitempty"`
	WithoutTenantId string `json:"withoutTenantId,omitempty"`
}

type ListCommonResponse struct {
	Data  json.RawMessage `json:"data"`
	Total int             `json:"total"`
	Start int             `json:"start"`
	Size  int             `json:"size"`
}
