// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package validators

import "testing"

func TestBindings_AllRequirementsBound(t *testing.T) {
	// Every DNS-REQ-XXX in DNSSecurityRequirements has at least one binding.
	allBindings := AllBindings()
	boundIDs := make(map[string]bool)
	for _, b := range allBindings {
		boundIDs[b.RequirementID] = true
	}
	for _, req := range DNSSecurityRequirements {
		if !boundIDs[req.ID] {
			t.Errorf("requirement %s (%s) has no binding", req.ID, req.Title)
		}
	}
}

func TestBindings_AllIDsValid(t *testing.T) {
	// Every binding references a valid requirement ID.
	validIDs := make(map[string]bool)
	for _, req := range DNSSecurityRequirements {
		validIDs[req.ID] = true
	}
	for _, b := range AllBindings() {
		if !validIDs[b.RequirementID] {
			t.Errorf("binding references invalid requirement ID: %s", b.RequirementID)
		}
	}
}

func TestBindings_ImplementedHaveValidators(t *testing.T) {
	// Implemented=true bindings must have non-nil StatelessFn or StatefulFn.
	for _, b := range AllBindings() {
		if b.Implemented && b.StatelessFn == nil && b.StatefulFn == nil {
			t.Errorf("binding %s (resource %d) is marked Implemented but has nil validators", b.RequirementID, b.Resource)
		}
	}
}

func TestBindings_UnimplementedAreNil(t *testing.T) {
	// Implemented=false bindings must have nil validators.
	for _, b := range AllBindings() {
		if !b.Implemented && (b.StatelessFn != nil || b.StatefulFn != nil) {
			t.Errorf("binding %s (resource %d) is not implemented but has non-nil validators", b.RequirementID, b.Resource)
		}
	}
}

func TestBindings_CoverageReport(t *testing.T) {
	// Informational — logs implemented vs pending.
	allBindings := AllBindings()
	implemented := 0
	pending := 0
	for _, b := range allBindings {
		if b.Implemented {
			implemented++
		} else {
			pending++
		}
	}
	t.Logf("STIG Validator Coverage: %d implemented, %d pending, %d total bindings", implemented, pending, len(allBindings))
	// List pending.
	for _, b := range allBindings {
		if !b.Implemented {
			t.Logf("  PENDING: %s (resource %d)", b.RequirementID, b.Resource)
		}
	}
}
