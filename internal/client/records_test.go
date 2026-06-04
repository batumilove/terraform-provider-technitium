// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package client

import "testing"

func TestRecordValueParam_FWD(t *testing.T) {
	if got := RecordValueParam("FWD"); got != "forwarder" {
		t.Fatalf("RecordValueParam(FWD) = %q, want forwarder", got)
	}
}

func TestRecordValueFromRData_FWD(t *testing.T) {
	rdata := map[string]interface{}{"forwarder": "dns.quad9.net:853 (9.9.9.9)"}
	if got := RecordValueFromRData("FWD", rdata); got != "dns.quad9.net:853 (9.9.9.9)" {
		t.Fatalf("RecordValueFromRData(FWD) = %q", got)
	}
}
