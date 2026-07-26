// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccZoneResource_Forwarder is the end-to-end regression test for issue #75
// ("Forwarder zone - cannot be created").
//
// Technitium refuses to create a Forwarder zone unless the request either
// supplies an initial `forwarder` parameter or opts out with
// `initializeForwarder=false`. The provider does the latter so that FWD records
// can be managed declaratively as separate technitium_record resources instead
// of being baked into zone creation.
//
// That behaviour was previously covered only by a httptest mock
// (TestZoneCreate_ForwarderCreatesEmptyZone), which asserts the query string the
// client *sends* and therefore cannot detect a server that rejects it. Verified
// against live Technitium 15.4:
//
//	POST /api/zones/create?zone=X&type=Forwarder
//	  -> {"status":"error","errorMessage":"Parameter 'forwarder' missing."}
//	POST /api/zones/create?zone=X&type=Forwarder&initializeForwarder=false
//	  -> {"status":"ok"}
//
// The reporter used a NAMED zone; the mock covers the root zone ".". Both are
// exercised here because "works at the apex" would not have implied the other.
func TestAccZoneResource_Forwarder(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccZoneResourceForwarderConfig("acc-forwarder.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_zone.forwarder", "name", "acc-forwarder.example.com"),
					resource.TestCheckResourceAttr("technitium_zone.forwarder", "type", "Forwarder"),
					resource.TestCheckResourceAttr("technitium_zone.forwarder", "status", "enabled"),
				),
			},
			{
				ResourceName:            "technitium_zone.forwarder",
				ImportState:             true,
				ImportStateId:           "acc-forwarder.example.com",
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"soa_serial_date_scheme", "dnssec"},
			},
		},
	})
}

// TestAccZoneResource_ForwarderWithRecord covers the workflow the provider docs
// actually recommend for issue #75: create the Forwarder zone empty, then attach
// FWD records as independent resources. This is the combination that broke —
// zone creation failing meant the documented example could never be applied.
func TestAccZoneResource_ForwarderWithRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccZoneForwarderWithRecordConfig("acc-fwd-rec.example.com", "1.1.1.1", "Udp", 1, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_zone.fwd_zone", "type", "Forwarder"),
					resource.TestCheckResourceAttr("technitium_record.fwd", "type", "FWD"),
					resource.TestCheckResourceAttr("technitium_record.fwd", "value", "1.1.1.1"),
					resource.TestCheckResourceAttr("technitium_record.fwd", "protocol", "Udp"),
					resource.TestCheckResourceAttr("technitium_record.fwd", "forwarder_priority", "1"),
					resource.TestCheckResourceAttr("technitium_record.fwd", "dnssec_validation", "true"),
				),
			},
		},
	})
}

func testAccZoneResourceForwarderConfig(name string) string {
	return testAccProviderHCL() + fmt.Sprintf(`
resource "technitium_zone" "forwarder" {
  name = %q
  type = "Forwarder"
}
`, name)
}

func testAccZoneForwarderWithRecordConfig(zone, forwarder, protocol string, priority int, dnssec bool) string {
	return testAccProviderHCL() + fmt.Sprintf(`
resource "technitium_zone" "fwd_zone" {
  name = %q
  type = "Forwarder"
}

resource "technitium_record" "fwd" {
  zone               = technitium_zone.fwd_zone.name
  name               = %q
  type               = "FWD"
  value              = %q
  protocol           = %q
  forwarder_priority = %d
  dnssec_validation  = %t
  overwrite          = false
}
`, zone, zone, forwarder, protocol, priority, dnssec)
}
