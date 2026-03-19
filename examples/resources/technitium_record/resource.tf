resource "technitium_record" "web" {
  zone  = "example.com"
  name  = "www.example.com"
  type  = "A"
  value = "192.168.1.100"
  ttl   = 3600
}
