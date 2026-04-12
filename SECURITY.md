# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| v0.6.x  | :white_check_mark: |
| v0.5.x  | :white_check_mark: |
| < 0.5   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability within futuapi4go, please report it responsibly.

**Please do NOT report security vulnerabilities through public GitHub Issues.**

Instead, please report them via one of the following methods:

### Preferred: GitHub Security Advisories

1. Navigate to the [Security Advisories](https://github.com/shing1211/futuapi4go/security/advisories) page
2. Click "Report a vulnerability"
3. Fill out the vulnerability report form with as much detail as possible

### Alternative: Email

Send an email to **shing1211@users.noreply.github.com** with:
- Subject: `[SECURITY] futuapi4go vulnerability report`
- Description of the vulnerability
- Steps to reproduce
- Potential impact assessment
- Any suggested fixes (optional)

## What to Expect

- **Acknowledgment**: We aim to acknowledge your report within **48 hours**
- **Initial Assessment**: We will assess the severity and impact within **7 days**
- **Resolution**: We target a fix within **30 days** for critical issues
- **Disclosure**: We will coordinate disclosure with you before any public announcement
- **Credit**: With your permission, we will credit you in the security advisory (unless you prefer anonymity)

## Security Best Practices for Users

When using futuapi4go:

- **Never hardcode credentials** — Use environment variables or secure vaults
- **Protect your Futu OpenD connection** — Do not expose the OpenD port publicly
- **Use TLS/SSH tunneling** if connecting across untrusted networks
- **Rotate API credentials** regularly
- **Run with minimal permissions** — Do not run the trading client as root/administrator

## Security Considerations

futuapi4go is a client library that communicates with Futu OpenD. Security responsibilities are shared between:

- **Futu OpenD**: Handles authentication, encryption, and server-side security
- **This library**: Safely marshals/unmarshals protobuf data and manages TCP connections
- **The user**: Securing their trading password, account credentials, and network environment

## Scope

This security policy covers the futuapi4go codebase. Issues with Futu's servers, Futu OpenD, or the official Futu API should be reported to Futu directly.
