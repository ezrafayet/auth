
*Note: this project is in its early stages. This is an implementation of an authentication and authorization service that adheres to industry standards and best practices.*

# Introduction

"Therefore IAM" provides an authentication and authorization solution for other services within a distributed system. It is meant to be generic and can be seamlessly integrated across multiple projects.

# Quick start

1. Within the folder services/iam, copy .env.example to .env and fill in the required values.


2. Run with docker:
```
docker-compose build && docker-compose up
```

Then navigate to: http://localhost:5050

# High level architecture

The project is composed of two main services:
- The API gateway
- The IAM service

Additionally, it includes two dummy services to facilitate tests and operations:
- The Client service
- The Core service

The following draft diagram illustrates the overall architecture of the project:
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
                                            +--->|  Core   |
                                                 +---------+
                                                    :8000
```

# Authentication

Authentication is the process of verifying the identity of a user, device, or system. In today's digital landscape, authentication is an essential aspect of any application, encompassing a variety of modern flows and paradigms.

## Authentication flows

We should support various authentication flows, including user access and programmatic access.

1. [Resource Owner Password Credentials (ROPC)](documentation/authentication/ropc.md)
2. [One-Time Password (OTP) Authentication](documentation/authentication/otp.md)
3. [Bearer Token Authentication](documentation/authentication/bearer_token.md)
4. Social Authentication Flows (OAuth 2.0)
5. Single Sign-On (SSO) Flows

Additionally, we should describe switching between user flows.

# Authorisation

Authorization determines the level of access or permissions a verified user has within a system, dictating what actions they can perform or resources they can access.

## Authorization paradigms

1. Role-based access control (RBAC)
2. Attribute-based access control (ABAC)

# Security

This section outlines various additional authentication factors and protective measures to prevent specific attacks.

## Additional authentication factors

1. CAPTCHA
2. Multi-Factor Authentication (MFA)

## Protections about specific attacks

1. Brute Force Attacks
2. Credential Stuffing Attacks
3. Phishing Attacks
4. Man-in-the-middle (MITM) Attacks
5. Cross-Site Request Forgery (CSRF) Attacks
6. Cross-Site Scripting (XSS) Attacks
7. SQL Injection Attacks
8. Session Hijacking Attacks
9. Clickjacking Attacks
10. Denial-of-Service (DoS) Attacks
11. Eavesdropping Attacks
12. Insecure Direct Object References (IDOR) Attacks

## Zero Trust model

[In progress]

# Contributing

[In progress]

# Releases

[In progress]

# Resources

- https://datatracker.ietf.org/doc/html/rfc6749
