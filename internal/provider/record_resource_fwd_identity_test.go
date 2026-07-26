// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// fwdModel builds the minimum FWD state needed to exercise the param builders.
func fwdModel(forwarder, protocol string, priority int64, dnssec *bool) *RecordResourceModel {
	m := &RecordResourceModel{
		Zone:              types.StringValue("fwd.example.com"),
		Name:              types.StringValue("fwd.example.com"),
		Type:              types.StringValue("FWD"),
		Value:             types.StringValue(forwarder),
		Protocol:          types.StringValue(protocol),
		ForwarderPriority: types.Int64Value(priority),
		DNSSECValidation:  types.BoolNull(),
	}
	if dnssec != nil {
		m.DNSSECValidation = types.BoolValue(*dnssec)
	}
	return m
}

// TestBuildDeleteParams_FWD_DNSSECIsIdentifying is the regression test for the
// destroy half of the FWD identity change.
//
// buildRecordID treats dnssecValidation as part of a FWD record's identity, so
// the delete request states it too. This asserts the REQUEST SHAPE only.
//
// It deliberately does not claim the server disambiguates. Measured against live
// Technitium 15.4, delete IGNORES dnssecValidation when matching: with two
// records differing only by this field, all four combinations of creation order
// and parameter value removed the FIRST-CREATED record. A colliding pair
// therefore cannot be individually destroyed through the 15.4 API. Sending the
// parameter makes the request express the caller's intent and starts working if
// the server ever honours it; the provider contains the problem by marking
// dnssec_validation RequiresReplace so it never attempts an in-place update,
// which is the path that silently collapses two records into one.
func TestBuildDeleteParams_FWD_DNSSECIsIdentifying(t *testing.T) {
	r := &RecordResource{}
	tr, fa := true, false

	paramsTrue := r.buildDeleteParams(fwdModel("1.1.1.1", "Udp", 1, &tr))
	paramsFalse := r.buildDeleteParams(fwdModel("1.1.1.1", "Udp", 1, &fa))

	if got := paramsTrue["dnssecValidation"]; got != "true" {
		t.Errorf("dnssecValidation = %q, want \"true\"", got)
	}
	if got := paramsFalse["dnssecValidation"]; got != "false" {
		t.Errorf("dnssecValidation = %q, want \"false\"", got)
	}

	// The point of the test: two records identical in every other respect must
	// produce DIFFERENT delete requests, or the API cannot tell them apart.
	if paramsTrue["dnssecValidation"] == paramsFalse["dnssecValidation"] {
		t.Fatal("delete params for the true/false siblings are identical — " +
			"the API would delete whichever was created first")
	}
	for _, k := range []string{"forwarder", "protocol", "forwarderPriority"} {
		if paramsTrue[k] != paramsFalse[k] {
			t.Errorf("%s differs between siblings (%q vs %q); the collision case "+
				"requires these to match so dnssecValidation is the only discriminator",
				k, paramsTrue[k], paramsFalse[k])
		}
	}
}

// TestBuildDeleteParams_FWD_DNSSECOmittedWhenNull covers the legacy path: a
// record imported with a 3-field ID has dnssec_validation null, and the delete
// request must then omit the parameter rather than assert a value the user
// never supplied.
func TestBuildDeleteParams_FWD_DNSSECOmittedWhenNull(t *testing.T) {
	r := &RecordResource{}
	params := r.buildDeleteParams(fwdModel("1.1.1.1", "Udp", 1, nil))
	if v, present := params["dnssecValidation"]; present {
		t.Errorf("dnssecValidation = %q, want the key to be absent when state is null", v)
	}
	if params["forwarder"] != "1.1.1.1" || params["protocol"] != "Udp" || params["forwarderPriority"] != "1" {
		t.Errorf("legacy identity params malformed: %v", params)
	}
}
