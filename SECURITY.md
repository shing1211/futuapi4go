# Security Policy

## Reporting a Vulnerability

Found a security issue? **Do not open a public GitHub issue.**

### Option 1 — GitHub Security Advisories (Preferred)
Go to [Security Advisories](https://github.com/shing1211/futuapi4go/security/advisories/new) and click "Report a vulnerability".

### Option 2 — Email
Send to **shing1211@users.noreply.github.com** with subject: `[SECURITY] futuapi4go vulnerability report`

## Supported Versions

Only the latest release receives security fixes. The SDK tracks the current
Futu OpenD protocol — see [docs/VERSION_MAP.md](docs/VERSION_MAP.md).

| Version | Supported |
|---------|-----------|
| Latest `v0.20.x` | ✅ |
| Older releases | ❌ |

## Response Targets

- Acknowledgement within **7 days**.
- Triage and severity assessment within **14 days**.
- A fix is released as a patch; the reporter is credited in the release notes
  unless they prefer to remain anonymous.

## Scope

- **In scope:** this repository's Go code — connection handling, crypto
  (RSA/AES), packet parsing, and request construction.
- **Out of scope:** the Futu OpenD binary, Futu's servers, and user
  misconfiguration (e.g. exposing the OpenD port or hardcoding credentials).

## Best Practices

- Never hardcode credentials — use environment variables
- Protect your OpenD port — don't expose it to untrusted networks
- Use TLS/SSH tunneling if OpenD must be accessed across networks
- Always test trading in simulate mode first
