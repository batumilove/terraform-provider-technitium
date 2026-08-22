// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestRecordResourceModifyPlan_UpdateInvalidatesComputedState(t *testing.T) {
	ctx := context.Background()
	r := &RecordResource{}
	s := recordResourceSchemaForPlanTest(t, ctx, r)

	state := recordPlanTestValue(t, ctx, s, "192.168.10.217",
		"home.arpa::herd-sandbox.home.arpa::A::192.168.10.217",
		"2026-06-07T18:58:07.59Z")
	plan := recordPlanTestValue(t, ctx, s, "192.168.10.250",
		"home.arpa::herd-sandbox.home.arpa::A::192.168.10.217",
		"2026-06-07T18:58:07.59Z")

	resp := resource.ModifyPlanResponse{Plan: tfsdk.Plan{Schema: s, Raw: plan}}
	r.ModifyPlan(ctx, resource.ModifyPlanRequest{
		State: tfsdk.State{Schema: s, Raw: state},
		Plan:  tfsdk.Plan{Schema: s, Raw: plan},
	}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ModifyPlan returned errors: %v", resp.Diagnostics)
	}

	var got RecordResourceModel
	if diags := resp.Plan.Get(ctx, &got); diags.HasError() {
		t.Fatalf("reading modified plan: %v", diags)
	}
	if !got.ID.IsUnknown() {
		t.Errorf("id = %q, want unknown so Update may return the value-derived ID", got.ID.ValueString())
	}
	if !got.LastModified.IsUnknown() {
		t.Errorf("last_modified = %q, want unknown so Update may return the server timestamp", got.LastModified.ValueString())
	}
}

func TestRecordResourceModifyPlan_NoOpPreservesComputedState(t *testing.T) {
	ctx := context.Background()
	r := &RecordResource{}
	s := recordResourceSchemaForPlanTest(t, ctx, r)

	raw := recordPlanTestValue(t, ctx, s, "192.168.10.250",
		"home.arpa::herd-sandbox.home.arpa::A::192.168.10.250",
		"2026-08-21T20:26:36.4066303Z")
	resp := resource.ModifyPlanResponse{Plan: tfsdk.Plan{Schema: s, Raw: raw}}
	r.ModifyPlan(ctx, resource.ModifyPlanRequest{
		State: tfsdk.State{Schema: s, Raw: raw},
		Plan:  tfsdk.Plan{Schema: s, Raw: raw},
	}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ModifyPlan returned errors: %v", resp.Diagnostics)
	}

	var got RecordResourceModel
	if diags := resp.Plan.Get(ctx, &got); diags.HasError() {
		t.Fatalf("reading modified plan: %v", diags)
	}
	if got.ID.IsUnknown() || got.ID.ValueString() != "home.arpa::herd-sandbox.home.arpa::A::192.168.10.250" {
		t.Errorf("id = %v, want unchanged known ID", got.ID)
	}
	if got.LastModified.IsUnknown() || got.LastModified.ValueString() != "2026-08-21T20:26:36.4066303Z" {
		t.Errorf("last_modified = %v, want unchanged known timestamp", got.LastModified)
	}
}

func recordResourceSchemaForPlanTest(t *testing.T, ctx context.Context, r *RecordResource) schema.Schema {
	t.Helper()
	resp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("record schema returned errors: %v", resp.Diagnostics)
	}
	return resp.Schema
}

func recordPlanTestValue(t *testing.T, ctx context.Context, s schema.Schema, value, id, lastModified string) tftypes.Value {
	t.Helper()
	typ := s.Type().TerraformType(ctx)
	obj, ok := typ.(tftypes.Object)
	if !ok {
		t.Fatalf("record schema type = %T, want tftypes.Object", typ)
	}
	attrs := make(map[string]tftypes.Value, len(obj.AttributeTypes))
	for name, attrType := range obj.AttributeTypes {
		var v interface{}
		switch name {
		case "id":
			v = id
		case "zone":
			v = "home.arpa"
		case "name":
			v = "herd-sandbox.home.arpa"
		case "type":
			v = "A"
		case "ttl":
			v = int64(900)
		case "value":
			v = value
		case "overwrite":
			v = false
		case "last_modified":
			v = lastModified
		default:
			v = nil
		}
		attrs[name] = tftypes.NewValue(attrType, v)
	}
	return tftypes.NewValue(obj, attrs)
}
