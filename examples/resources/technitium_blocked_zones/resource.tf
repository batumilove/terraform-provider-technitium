resource "technitium_blocked_zones" "malware" {
  domains = [
    "malware.example.com",
    "phishing.example.com",
    "tracking.example.com",
  ]
}
