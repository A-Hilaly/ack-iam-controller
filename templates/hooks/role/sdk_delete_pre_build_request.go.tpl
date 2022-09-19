	// This causes syncPolicies to delete all associated policies from the role
	rWithNoPolicies := &resource{ko: &svcapitypes.Role{}}
	if err := rm.syncPolicies(ctx, r, rWithNoPolicies); err != nil {
		return nil, err
	}