resource "technitium_allowed_zones" "corporate" {
  domains = [
    "internal.example.com",
    "vpn.example.com",
    "mail.example.com",
  ]
}
