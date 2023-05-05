[Back to README](../../README.md)

# One-Time Password (OTP) Authentication

The One-Time Password (OTP) Authentication flow is a way to exchange a user's one-time password for an access token. Here we will refer to this flow as the "Magic link flow", where the user receives a link via email that they can click to log in.

## Configuration

[In progress]

## Api endpoints

### User creation

```text
Diagram
```

#### 1. Endpoint to create a user

Request:
```
POST /api/internal/v1/auth/register
{
    "authMethod": "magicLink",
    "username": "string (mandatory)"
    "email": "string, (mandatory)"
}
```

Response:
```
code: 200
{
    "status": "success",
    "message": "string",
    "data": {
        "userId": "string",
    }
}
```

Errors:
```
- 422 invalidData -> provides a list of fields
- 409 usernameExists
- 409 emailExists
- 429 rateLimited
- 500 internalError
```

#### 2. Endpoint to send a validation email

Request:
```
POST /api/internal/v1/auth/users/{userId}/email-verification
{}
```

Response:
```
code 200
{
    "status": "success",
    "message": "string"
}
```

Errors:
```
- 422 invalidData -> provides a list of fields
- 404 userNotFound
- 409 emailAlreadyVerified
- 429 rateLimited
- 500 internalError
```

### Validation of an email and login

```text
Diagram
```

#### 3. Endpoint to validate an email (link received by email)

The user must arrive on a page that will call the following endpoint to validate the email.

Request:
```text
PATCH /api/internal/v1/auth/users/{userId}/email-verification/{verificationToken}
{}
```

Response:
```text
code 200
{
    "status": "success",
    "message": "string",
    "data": {
        "loginToken": "string"
    }
}
```

Errors:
```text
- 422 invalidData -> must have a list of fields
- 404 tokenNotFound
- 401 tokenExpired -> must send userId
- 404 userNotFound
- 403 userBlocked
- 410 userDeleted
- 409 emailAlreadyVerified
- 429 rateLimited
- 500 internalError
```

### User wants to log in

```text
Diagram
```

#### 4. Endpoint to ask for a login link

Request:
```text
POST /api/internal/v1/auth/magic-link
{
    "email": "string, (mandatory)"
}
```

Response:
```text
code 200
{
    "status": "success",
    "message": "string"
}
```

Errors:
```text
- 422 invalidData -> must have a list of fields
- 404 userNotFound
- 403 emailNotVerified -> must send userId
- 403 userBlocked
- 410 userDeleted
- 429 rateLimited
- 500 internalError
```

#### 5. Endpoint to login

Request:
```text
POST /api/internal/v1/auth/login/{loginToken}
{}
```

Response:
```text
code 200
{
    "status": "success",
    "message": "string"
}
```
+ cookies (access_token, refresh_token)

Errors:
```text
- 422 invalidData -> must have a list of fields
- 404 tokenNotFound
- 403 tokenExpired
- 403 emailNotVerified -> must send userId
- 403 userBlocked
- 410 userDeleted
- 429 rateLimited
- 500 internalError
```


<br/>[Back to README](../../README.md)