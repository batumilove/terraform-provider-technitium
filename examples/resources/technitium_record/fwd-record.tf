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

resource "technitium_record" "cloudflare_fallback" {
  zone               = technitium_zone.root_forwarder.name
  name               = "."
  type               = "FWD"
  value              = "1.1.1.1"
  protocol           = "Udp"
  forwarder_priority = 2
  overwrite          = false
}
