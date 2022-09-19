// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateRoleInput) {
	res := &svcsdk.UpdateRoleInput{}
{{ GoCodeSetUpdateInput .CRD "r.ko" "res" 1 }}
	return res
}
