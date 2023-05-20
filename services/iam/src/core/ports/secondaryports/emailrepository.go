package secondaryports

import (
	"iam/src/core/domain/types"
)

type EmailRepository interface {
	WelcomeNewUser(email types.Email, username types.Username) error
	SendVerificationCode(email types.Email, username types.Username, verificationCodeForURL string) error
	SendMagicLink(email types.Email, username types.Username, authorizationCodeForURL string) error
}
