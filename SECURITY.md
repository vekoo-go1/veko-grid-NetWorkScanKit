# Security Policy

## Supported Versions

Currently supported versions of Veko Grid:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability in Veko Grid, please report it responsibly:

### How to Report

1. **Email**: Create an issue on GitHub with the label "security"
2. **Include**: Detailed description of the vulnerability
3. **Provide**: Steps to reproduce (if applicable)
4. **Timeline**: We aim to respond within 48 hours

### What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact assessment
- Suggested fix (if you have one)

### What NOT to Report as Security Issues

- Issues related to misuse of the tool for unauthorized scanning
- Reports about the tool being "too powerful" - this is by design for legitimate security research
- General feature requests or bugs that don't pose security risks

## Security Features

Veko Grid includes several built-in security and ethical safeguards:

### Rate Limiting
- Built-in delays between requests to prevent DoS-style scanning
- Configurable threading limits
- Timeout mechanisms to prevent hanging connections

### Legal Safeguards
- Prominent warnings about legal usage
- Built-in documentation about ethical guidelines
- No features designed for destructive attacks

### Privacy Protection
- Support for TOR and proxy routing for legitimate privacy needs
- No data collection or telemetry
- Results stored locally only

## Responsible Usage

This tool is designed for:
- ✅ Security research and education
- ✅ Authorized penetration testing
- ✅ Bug bounty programs within scope
- ✅ Auditing your own infrastructure

This tool should NOT be used for:
- ❌ Unauthorized scanning of systems you don't own
- ❌ Circumventing security measures
- ❌ Any illegal activities

## Dependencies Security

We regularly monitor our Go dependencies for known vulnerabilities:

- Go standard library (latest stable version)
- Third-party dependencies are minimized and regularly updated
- No runtime dependencies on external services

## Secure Development Practices

- Code is open source and publicly auditable
- No hardcoded secrets or backdoors
- Minimal external dependencies
- Static analysis during development
- Regular security reviews

## Disclosure Timeline

- **T+0**: Vulnerability reported
- **T+48h**: Initial response and acknowledgment
- **T+7d**: Investigation and assessment complete
- **T+30d**: Fix developed and tested
- **T+45d**: Public disclosure (if appropriate)

We believe in responsible disclosure and will work with security researchers to ensure vulnerabilities are addressed promptly and professionally.