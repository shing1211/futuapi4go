# Security Policy

## Reporting a Vulnerability

Found a security issue? **Do not open a public GitHub issue.**

### Option 1 — GitHub Security Advisories (Preferred)
Go to [Security Advisories](https://github.com/shing1211/futuapi4go/security/advisories/new) and click "Report a vulnerability".

### Option 2 — Email
Send to **shing1211@users.noreply.github.com** with subject: `[SECURITY] futuapi4go vulnerability report`

## Best Practices

- Never hardcode credentials — use environment variables
- Protect your OpenD port — don't expose it to untrusted networks
- Use TLS/SSH tunneling if OpenD must be accessed across networks
- Always test trading in simulate mode first
