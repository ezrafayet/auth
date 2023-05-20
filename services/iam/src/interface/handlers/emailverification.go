package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/httphelpers"
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
	var args primaryports.SendVerificationCodeArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}

	err = h.emailVerificationService.Send(args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}

	httphelpers.WriteSuccess(http.StatusNoContent, "Verification code sent successfully", struct{}{})(w, r)
}

func (h *EmailVerificationHandler) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	var args primaryports.ConfirmEmailArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	fmt.Println("Body")
	fmt.Println(r.Body)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}

	answer, err := h.emailVerificationService.Confirm(args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}

	httphelpers.WriteSuccess(http.StatusAccepted, "Email verified successfully", answer)(w, r)
}
