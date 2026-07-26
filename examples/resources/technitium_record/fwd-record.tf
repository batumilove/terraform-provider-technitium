# Full example: a primary forwarder with DNSSEC validation over DNS-over-TLS,
# plus a plain fallback.
resource "technitium_zone" "root_forwarder" {
  name = "."
  type = "Forwarder"
}

resource "technitium_record" "quad9_forwarder" {
  zone               = technitium_zone.root_forwarder.name
  name               = "."
  type               = "FWD"
  value              = "dns.quad9.net:853 (9.9.9.9)"
  protocol           = "Tls"
  forwarder_priority = 1
  dnssec_validation  = true
  overwrite          = false
}

# Lower priority is queried first, so this is the fallback. It also differs from
# the record above by value and protocol, which keeps the two independently
# addressable.
#
# When several forwarders share a zone, make sure any two differ by something
# other than dnssec_validation alone — value, protocol or forwarder_priority.
# Records that differ ONLY by dnssec_validation cannot be told apart by the
# Technitium API and one of them can be silently lost. See "DNSSEC validation on
# forwarders" in the resource documentation.
resource "technitium_record" "cloudflare_fallback" {
  zone               = technitium_zone.root_forwarder.name
  name               = "."
  type               = "FWD"
  value              = "1.1.1.1"
  protocol           = "Udp"
  forwarder_priority = 2
  dnssec_validation  = false
  overwrite          = false
}
