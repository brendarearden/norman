package authorization

import (
	"net/http"

	"github.com/rancher/norman/pkg/httperror"

	"github.com/rancher/norman/pkg/types"
	"github.com/rancher/norman/pkg/types/slice"
)

type AllAccess struct {
}

func (*AllAccess) CanCreate(apiOp *types.APIOperation, schema *types.Schema) error {
	if slice.ContainsString(schema.CollectionMethods, http.MethodPost) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not create "+schema.ID)
}

func (*AllAccess) CanGet(apiOp *types.APIOperation, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, http.MethodGet) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not get "+schema.ID)
}

func (*AllAccess) CanList(apiOp *types.APIOperation, schema *types.Schema) error {
	if slice.ContainsString(schema.CollectionMethods, http.MethodGet) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not list "+schema.ID)
}

func (*AllAccess) CanUpdate(apiOp *types.APIOperation, obj map[string]interface{}, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, http.MethodPut) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not update "+schema.ID)
}

func (*AllAccess) CanDelete(apiOp *types.APIOperation, obj map[string]interface{}, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, http.MethodDelete) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not delete "+schema.ID)
}

func (*AllAccess) CanDo(apiGroup, resource, verb string, apiOp *types.APIOperation, obj map[string]interface{}, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, verb) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not perform "+verb+" "+schema.ID)
}

func (*AllAccess) Filter(apiOp *types.APIOperation, schema *types.Schema, obj map[string]interface{}) map[string]interface{} {
	return obj
}

func (*AllAccess) FilterList(apiOp *types.APIOperation, schema *types.Schema, obj []map[string]interface{}) []map[string]interface{} {
	return obj
}