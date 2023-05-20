package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/apperrors"
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
		httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		return
	}

	err = h.emailVerificationService.Send(args)

	if err != nil {
		switch err.Error() {
		case apperrors.InvalidId:
			httphelpers.WriteError(http.StatusBadRequest, "error", err.Error())(w, r)
		case apperrors.UserNotFound:
			httphelpers.WriteError(http.StatusNotFound, "error", err.Error())(w, r)
		case apperrors.EmailAlreadyVerified:
			httphelpers.WriteError(http.StatusForbidden, "error", err.Error())(w, r)
		case apperrors.LimitExceeded:
			httphelpers.WriteError(http.StatusTooManyRequests, "error", err.Error())(w, r)
		default:
			fmt.Println(err)
			httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		}
		return
	}

	httphelpers.WriteSuccess(http.StatusNoContent, "Verification code sent successfully", struct{}{})(w, r)
}

func (h *EmailVerificationHandler) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	var args primaryports.ConfirmEmailArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		return
	}

	answer, err := h.emailVerificationService.Confirm(args)

	if err != nil {
		if err != nil {
			switch err.Error() {
			case apperrors.InvalidCode:
				httphelpers.WriteError(http.StatusBadRequest, "error", err.Error())(w, r)
			case apperrors.VerificationCodeNotFound:
				httphelpers.WriteError(http.StatusNotFound, "error", err.Error())(w, r)
			case apperrors.VerificationCodeExpired:
				httphelpers.WriteError(http.StatusGone, "error", err.Error())(w, r)
			case apperrors.UserNotFound:
				httphelpers.WriteError(http.StatusNotFound, "error", err.Error())(w, r)
			case apperrors.EmailAlreadyVerified:
				httphelpers.WriteError(http.StatusForbidden, "error", err.Error())(w, r)
			case apperrors.UserBlocked:
				httphelpers.WriteError(http.StatusForbidden, "error", err.Error())(w, r)
			case apperrors.UserDeleted:
				httphelpers.WriteError(http.StatusForbidden, "error", err.Error())(w, r)
			default:
				fmt.Println(err)
				httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
			}
			return
		}
	}

	httphelpers.WriteSuccess(http.StatusAccepted, "Email verified successfully", answer)(w, r)
}
