resource "technitium_record" "sip" {
  zone     = "example.com"
  name     = "_sip._tcp.example.com"
  type     = "SRV"
  value    = "sip.example.com"
  priority = 10
  weight   = 60
  port     = 5060
  ttl      = 3600
}
