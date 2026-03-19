# Security Policy

## Supported Versions

| Version | Supported          |
|---------|--------------------|
| latest  | Yes                |
| < latest | No                |

## Reporting a Vulnerability

If you discover a security vulnerability in this provider, please report it responsibly.

**Do NOT open a public GitHub issue for security vulnerabilities.**

Instead, please use [GitHub's private vulnerability reporting](https://github.com/darkhonor/terraform-provider-technitium/security/advisories/new) to submit your report. This keeps the conversation private between you and the maintainers until a fix is available.

In your report, please include:

1. A description of the vulnerability
2. Steps to reproduce the issue
3. Potential impact assessment
4. Any suggested fixes (optional)

You should receive a response within 72 hours. We will work with you to understand the issue and coordinate a fix and disclosure timeline.

## Scope

This security policy covers:

- The Terraform provider binary and its source code
- The STIG compliance validation engine
- Credential handling (API tokens, TSIG keys)

This policy does **not** cover:

- The Technitium DNS Server itself (report to [TechnitiumSoftware](https://github.com/TechnitiumSoftware/DnsServer))
- HashiCorp Terraform core (report to [HashiCorp](https://www.hashicorp.com/security))
