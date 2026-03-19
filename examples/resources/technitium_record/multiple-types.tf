resource "technitium_record" "ipv6" {
  zone  = "example.com"
  name  = "www.example.com"
  type  = "AAAA"
  value = "2001:db8::1"
}

resource "technitium_record" "alias" {
  zone  = "example.com"
  name  = "app.example.com"
  type  = "CNAME"
  value = "www.example.com"
}

resource "technitium_record" "spf" {
  zone  = "example.com"
  name  = "example.com"
  type  = "TXT"
  value = "v=spf1 mx -all"
}

resource "technitium_record" "ns" {
  zone  = "example.com"
  name  = "sub.example.com"
  type  = "NS"
  value = "ns1.example.com"
}

resource "technitium_record" "ptr" {
  zone  = "1.168.192.in-addr.arpa"
  name  = "100.1.168.192.in-addr.arpa"
  type  = "PTR"
  value = "www.example.com"
}
