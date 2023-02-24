package common

import "strconv"

func MakeCommonQuery(req ListCommonRequest, tenant WithTenant, query map[string]string) {
	if len(req.Sort) > 0 {
		query["sort"] = req.Sort
	}

	if len(req.Order) > 0 {
		query["order"] = req.Order
	}

	if req.Start < 0 {
		req.Start = 0
	}
	query["start"] = strconv.Itoa(req.Start)

	if req.Size < 1 {
		req.Size = 10
	}
	query["size"] = strconv.Itoa(req.Size)

	if len(tenant.TenantId) > 0 {
		query["tenantId"] = tenant.TenantId
	}

	if len(tenant.TenantIdLike) > 0 {
		query["tenantIdLike"] = tenant.TenantIdLike
	}

	if tenant.WithoutTenantId != nil && *tenant.WithoutTenantId {
		query["withoutTenantId"] = "true"
	}
}
