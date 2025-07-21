# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability in grd, please follow these steps:

1. **Do NOT disclose the vulnerability publicly** (no GitHub issues for security vulnerabilities).
2. Email the maintainer directly at [contact@devoria.me] with details about the vulnerability.
3. Provide as much information as possible about the vulnerability, including:
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if you have one)

## What to Expect

When you report a vulnerability, you can expect:

1. An acknowledgment of your report within 48 hours.
2. A determination of the vulnerability's validity and severity.
3. A plan for addressing the vulnerability and releasing a fix.
4. Credit in the release notes when the vulnerability is fixed (unless you prefer to remain anonymous).

## Security Best Practices When Using grd

While grd itself focuses on error handling and should not introduce security vulnerabilities directly, here are some recommendations for secure usage:

1. Keep your Go environment and dependencies up to date.
2. Be careful with the error messages you expose in `Catch` blocks, especially in production environments, to avoid leaking sensitive information.
3. Consider sanitizing or logging important errors separately from what is returned to users.

Thank you for helping keep grd secure!
