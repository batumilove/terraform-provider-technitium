// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package provider

// nssCompliantTSIGAlgorithms lists HMAC algorithms that meet FIPS 140-3 / CNSSI 1253
// requirements for National Security Systems. Used by both technitium_tsig_key and
// technitium_zone resources for cross-resource compliance enforcement.
var nssCompliantTSIGAlgorithms = map[string]bool{
	"hmac-sha256": true,
	"hmac-sha384": true,
	"hmac-sha512": true,
}

// isNSSCompliantTSIGAlgorithm checks if a TSIG algorithm meets FIPS 140-3 requirements.
func isNSSCompliantTSIGAlgorithm(algo string) bool {
	return nssCompliantTSIGAlgorithms[algo]
}
