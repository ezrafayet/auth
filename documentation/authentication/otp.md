[Back to README](../../README.md)

[Document in progress]

# One-Time Password (OTP) Authentication

The One-Time Password (OTP) Authentication flow is a way to exchange a user's one-time password for an access token. Here we will use it in a "Magic link flow", where the user logs in by clicking a link in an email they receive.

## Definitions

- **Verification code**: A verification code is a unique, time-limited code that is used to verify the user's email address.
- **Authorization code**: An authorization code is a unique, time-limited code that is used to exchange for an access token and a refresh token.
- **Access token**: An access token is a unique, time-limited JWT that is used to authenticate API requests on behalf of the user.
- **Refresh token**: A refresh token is a unique, time-limited token that is used to get a new Access Token.
- **Authorization server**: The server that issues access tokens and refresh tokens to the client after successfully authenticating the resource owner and obtaining authorization.

## Configuration

In order to be enabled, the environment must contain:
```text
ENABLE_AUTH_WITH_MAGIC_LINK = true
```

And an email service must be configured (default one is SendGrid):
```text
EMAIL_SERVICE_TOKEN = foobar12345
```

The email provider can easily be switched for another since it is injected and the core defines its interface.

## Flow to register a new user

1. The user enters their username and email address in the client.
2. The client sends a request to the server with the user's username and email.
3. The server checks if a user with this username or email exists in the system. If it does, it returns an error. If it does not, it creates a new user with the username and email, it generates a verification code and sends it to the user.
4. The user clicks on the verification link, which directs them to a specific callback URL in the client application.
5. When the client application receives the token, it sends a request to the server to exchange the validation code for an authorization code.
6. The server validates the token. If it's valid and has not expired, the server returns an authorization code to the client application.
7. The client application stores the authorization code securely and uses it to authenticate subsequent API requests on behalf of the user.

## Flow to login

1. The user enters their email address in the client. 
2. The client sends a request to the authorization server with the user's email.
3. The authorization server checks if a user with this email exists in the system. If it does, it generates a unique, time-limited "magic link" containing a one time authorization code. 
4. The server sends the magic link to the user's email address. 
5. The user clicks on the magic link, which directs them to a specific callback URL in the client application. 
6. When the client application receives the token, it sends a request to the authorization server to exchange the code for an access token and a refresh token. 
7. The authorization server validates the token. If it's valid and has not expired, the server returns an access token and a refresh token to the client application. 
8. The client application stores the access token and refresh token securely and uses it to authenticate subsequent API requests on behalf of the user.

## Api endpoints

The API has 5 endpoints to support those flows:
1. To create a user with a username and an email
2. To ask the server for a "validation email" with a validation code
3. To validate a user's account with a validation code
4. To ask the server for a "login email" with an authorization code
5. To exchange an authorization code for an access token and a refresh token

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
- 422 invalidData -> provides a list of fields
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
- 422 invalidData -> provides a list of fields
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
- 422 invalidData -> provides a list of fields
- 404 tokenNotFound
- 401 tokenExpired
- 403 emailNotVerified -> must send userId
- 404 userNotFound
- 403 userBlocked
- 410 userDeleted
- 429 rateLimited
- 500 internalError
```


<br/>[Back to README](../../README.md)