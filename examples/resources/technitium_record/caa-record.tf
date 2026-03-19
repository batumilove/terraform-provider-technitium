resource "technitium_record" "caa" {
  zone      = "example.com"
  name      = "example.com"
  type      = "CAA"
  value     = "letsencrypt.org"
  caa_flags = 0
  caa_tag   = "issue"
  ttl       = 3600
}
