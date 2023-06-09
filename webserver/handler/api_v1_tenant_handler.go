package handler

import (
	"context"
	"go_chi_template/oapi"
	"go_chi_template/src/app_service/shared"
	v1 "go_chi_template/src/app_service/v1"
)

func (h *Handler) GetApiV1Tenant(
	ctx context.Context,
	request oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	resp, err := v1.TenantListAppService(ctx, h.app, request)
	return resp, err
}

func (h *Handler) PostApiV1Tenant(
	ctx context.Context,
	request oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	errResp, err := shared.ValidateTenantCreateAppService(ctx, h.app, request.Body)

	if errResp != nil {
		return errResp, err
	}

	resp, err := v1.TenantCreateAppService(ctx, h.app, request)

	return resp, err
}

func (h *Handler) GetApiV1TenantTenantId(
	ctx context.Context,
	request oapi.GetApiV1TenantTenantIdRequestObject,
) (oapi.GetApiV1TenantTenantIdResponseObject, error) {
	resp, err := v1.TenantDetailAppService(ctx, h.app, request)
	return resp, err
}

func (h *Handler) PatchApiV1TenantTenantId(
	ctx context.Context,
	request oapi.PatchApiV1TenantTenantIdRequestObject,
) (oapi.PatchApiV1TenantTenantIdResponseObject, error) {
	resp, err := v1.TenantUpdateAppService(ctx, h.app, request)
	return resp, err
}
