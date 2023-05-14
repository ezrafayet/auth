package handlers

import (
	"iam/src/core/ports/primaryports"
	"net/http"
)

type EmailVerificationHandler struct {
	emailVerificationService primaryports.EmailVerificationService
}

func NewEmailVerificationHandler(emailVerificationService primaryports.EmailVerificationService) *EmailVerificationHandler {
	return &EmailVerificationHandler{
		emailVerificationService: emailVerificationService,
	}
}

func (h *EmailVerificationHandler) SendVerificationEmail(w http.ResponseWriter, r *http.Request) {

}

func (h *EmailVerificationHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {

}
