[DRAFT]

Note this document is a draft and is subject to change.

# Architecture

This is a draft of the overall architecture of the project:
```
                            +---------+
                        +-->| CLient  |
                        |   +---------+
        +---------+     |      :3000             +---------+
        |  Proxy  +-----+                   +--->|   IAM   |
        +---------+     |                   |    +---------+
           :5000        |                   |       :7777
                        +-------------------+
                                            |
                                            |    +---------+
                                            +--->|  Core   | <-- Any services
                                                 +---------+
                                                    :8000
```

[In progress]

# Authentication

## Authentication flows

- [Password flow](authentication/password_flow.md)
- Magic link flow
- OAuth2 flows
- SSO flows

# Authorisation

- Role-based access control (RBAC)
- Attribute-based access control (ABAC)

## Security

### Additional authentication factors

- Recaptcha
- MFA

### Protections about specific attacks

- Brute force attacks
- Credential stuffing attacks
- Phishing attacks
- Man-in-the-middle (MITM) attacks
- Cross-site request forgery (CSRF) attacks
- Cross-site scripting (XSS) attacks
- SQL injection attacks
- Session hijacking attacks
- Clickjacking attacks
- Denial-of-service (DoS) attacks
- Eavesdropping attacks
- Insecure Direct Object References (IDOR) attacks
