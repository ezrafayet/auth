# iam

[In progress]

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

## Quick start

```
docker-compose build
docker-compose up
```

## Use cases

Regarding users and authentication, we have the following use cases:
- Support for password user creation / access (Password flow)
- Support for magic link user creation / access (Magic link flow)
- Support for OAuth2 user creation / access
- Support for Recaptcha
- Support for MFA

Regarding authorisations:
- Support for role-based access control (RBAC)
- Support for attribute-based access control (ABAC)
