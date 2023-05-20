package emailprovider

import (
	"fmt"
	"iam/pkg/apperrors"
	"os"
)

type Provider struct {
	apiKey string
}

func NewEmailProvider() *Provider {
	apiKey := os.Getenv("EMAIL_PROVIDER_API_KEY")

	if apiKey == "" {
		panic(apperrors.EmailProviderNotSet)
	}

	return &Provider{apiKey: apiKey}
}

func (e *Provider) SendEmail(email string, subject string, keyValues map[string]any) error {
	if os.Getenv("ENV") == "DEV" {
		fmt.Println("[Gate Keeper] ----> Sending email to", email, "with subject", subject, "and keyValues", keyValues)
	} else {
		// todo: implement
		fmt.Println("[Gate Keeper] ----> Sending email to", email, "with subject", subject)
	}

	return nil
}
