# Simplest case: forward everything to one upstream resolver.
#
# A Forwarder zone is created empty, then the forwarder itself is a separate
# FWD record. dnssec_validation is optional and can be left out entirely — with
# a single forwarder there is nothing for it to be confused with.
resource "technitium_zone" "forwarder" {
  name = "."
  type = "Forwarder"
}

resource "technitium_record" "upstream" {
  zone      = technitium_zone.forwarder.name
  name      = "."
  type      = "FWD"
  value     = "1.1.1.1"
  protocol  = "Udp"
  overwrite = false
}
