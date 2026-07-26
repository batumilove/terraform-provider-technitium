# Conditional forwarding: send one internal namespace to the resolvers that are
# authoritative for it, while everything else follows the server's normal path.
#
# The Forwarder zone is named for the domain being forwarded rather than "." —
# only queries under that domain are forwarded.
resource "technitium_zone" "corp_internal" {
  name = "corp.example.net"
  type = "Forwarder"
}

# Two upstreams for redundancy. They differ by value, so they remain
# individually addressable; the priorities set which is tried first.
resource "technitium_record" "corp_dns_primary" {
  zone               = technitium_zone.corp_internal.name
  name               = technitium_zone.corp_internal.name
  type               = "FWD"
  value              = "10.10.0.53"
  protocol           = "Udp"
  forwarder_priority = 1
  overwrite          = false
}

resource "technitium_record" "corp_dns_secondary" {
  zone               = technitium_zone.corp_internal.name
  name               = technitium_zone.corp_internal.name
  type               = "FWD"
  value              = "10.20.0.53"
  protocol           = "Udp"
  forwarder_priority = 2
  overwrite          = false
}

# Internal resolvers usually serve names that do not validate publicly, so
# DNSSEC validation is left off here. Set dnssec_validation = true only for
# upstreams you expect to return signed, publicly-validatable answers.
