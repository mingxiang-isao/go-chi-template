package dto

import (
	"go_chi_template/db/go_chi_template/public/model"
)

type TenantWithUserCountDto struct {
	model.Tenant
	UserCount int `alias:"tenant.user_count"`
}
