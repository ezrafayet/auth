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
                                            +--->|  Core   | <-- TODO
                                                 +---------+
                                                    :8000
```

# Getting started

```
docker-compose build
docker-compose up
docker-compose stop
```

# Use cases

- Create a user using password (flow 1). 
  - Variant: verification email logs in directly
  - Variant: sends an email with a link or a code to verify the email
  - Variant: phone login
  - Variant: with / without recaptha
- Create a user using magic link (flow 2)
- OAuth2: (flow 3)
  - Google auth (flow 3)
  - Facebook auth (flow 4)
  - GitHub auth (flow 5)

- MFA
  - Email
  - SMS
  - Google Authenticator
  - Yubikey

- Recaptcha

- Switching between flows
