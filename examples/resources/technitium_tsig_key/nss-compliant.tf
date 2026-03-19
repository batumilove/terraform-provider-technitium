# NSS mode restricts TSIG algorithms to FIPS-compliant options:
# hmac-sha256, hmac-sha384, hmac-sha512
resource "technitium_tsig_key" "nss" {
  key_name  = "transfer.example.mil"
  algorithm = "hmac-sha384"
}
