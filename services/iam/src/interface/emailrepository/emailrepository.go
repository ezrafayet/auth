package emailrepository

import (
	"fmt"
	"iam/src/core/domain/types"
	"iam/src/infra/emailprovider"
)

type EmailRepository struct {
	EmailProvider *emailprovider.Provider
}

func NewEmailRepository(emailProvider *emailprovider.Provider) *EmailRepository {
	return &EmailRepository{EmailProvider: emailProvider}
}

func (e *EmailRepository) WelcomeNewUser(email types.Email, username types.Username) error {
	m := make(map[string]any)

	m["username"] = username

	return e.EmailProvider.SendEmail(string(email), "Welcome to IAM", m)
}

func (e *EmailRepository) SendVerificationCode(email types.Email, username types.Username, verificationCodeForURL string) error {
	m := make(map[string]any)

	m["username"] = username
	m["verificationCode"] = verificationCodeForURL

	fmt.Println(fmt.Printf("http://localhost:5050/auth/email-verification/%s", verificationCodeForURL))

	return e.EmailProvider.SendEmail(string(email), "Verify your email", m)
}

func (e *EmailRepository) SendMagicLink(email types.Email, username types.Username, authorizationCodeForURL string) error {
	m := make(map[string]any)

	m["username"] = username
	m["authorizationCode"] = authorizationCodeForURL

	fmt.Println(fmt.Printf("http://localhost:5050/login/email-verification/%s", authorizationCodeForURL))

	return e.EmailProvider.SendEmail(string(email), "Connect with your magic link", m)
}
