package v1

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/oapi"
	"go_chi_template/src/repository/single"
	"strconv"
)

func TenantUpdateAppService(
	ctx context.Context,
	app *config.App,
	req oapi.PatchApiV1TenantTenantIdRequestObject,
) (oapi.PatchApiV1TenantTenantIdResponseObject, error) {
	tenantId, _ := strconv.ParseInt(req.Body.TenantId, 10, 64)

	tenant, _ := single.GetTenantById(app, tenantId)

	tenantDto := oapi.Tenant{
		Id:        strconv.FormatInt(tenant.ID, 10),
		Name:      tenant.Name,
		ShortCode: tenant.ShortCode,
	}

	respDto := oapi.PatchApiV1TenantTenantId200JSONResponse{
		Tenant: &tenantDto,
	}

	return &respDto, nil
}
