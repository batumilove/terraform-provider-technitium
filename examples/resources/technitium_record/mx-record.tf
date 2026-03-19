resource "technitium_record" "mail" {
  zone     = "example.com"
  name     = "example.com"
  type     = "MX"
  value    = "mail.example.com"
  priority = 10
  ttl      = 3600
}
