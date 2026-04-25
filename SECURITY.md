# Security Policy

## Supported Versions

| Version | Status |
|---------|--------|
| **v0.3.x** | ✅ Actively maintained |
| v0.2.x | ✅ Receives security patches |
| v0.1.x | ✅ Receives security patches |
| < v0.1 | ❌ Not supported |

## Reporting a Vulnerability

Found a security issue? **Do not open a public GitHub issue.** Instead:

### Option 1 — GitHub Security Advisories (Preferred)

1. Go to [Security Advisories](https://github.com/shing1211/futuapi4go/security/advisories/new)
2. Click "Report a vulnerability"
3. Fill in as much detail as possible

### Option 2 — Email

Send to **shing1211@users.noreply.github.com** with:
- Subject: `[SECURITY] futuapi4go vulnerability report`
- Steps to reproduce
- Potential impact
- Any fix suggestions (optional)

## Response Timeline

| Stage | Timeline |
|-------|----------|
| **Acknowledgment** | Within 48 hours |
| **Initial assessment** | Within 7 days |
| **Critical fix** | Within 30 days |
| **Public disclosure** | Coordinated with reporter |

With your permission, we'll credit you in the advisory. Anonymity is fully supported.

## Security Best Practices for SDK Users

- **Never hardcode credentials** — use environment variables or a secrets manager
- **Protect your OpenD port** — don't expose it to untrusted networks
- **Use TLS/SSH tunneling** if OpenD must be accessed across networks
- **Rotate credentials** regularly
- **Run with minimal permissions** — don't run the trading client as root/admin
- **Always test trading in simulate mode** first

## Scope

This policy covers the futuapi4go codebase. Issues with Futu's servers, Futu OpenD, or the official API should be reported to [Futu directly](https://www.futunn.com/).

## Security Design Notes

futuapi4go is a client library — security is a shared responsibility:

| Layer | Responsibility |
|-------|---------------|
| **Futu OpenD** | Authentication, encryption, server-side security |
| **futuapi4go** | Safe protobuf marshaling, TCP connection management, no credential storage |
| **Your application** | Credential security, network configuration, trading safeguards |
