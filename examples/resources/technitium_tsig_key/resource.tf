resource "technitium_tsig_key" "example" {
  key_name  = "transfer.example.com"
  algorithm = "hmac-sha256"
}
