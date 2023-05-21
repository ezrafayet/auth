package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/apperrors"
	"iam/pkg/httphelpers"
	"iam/src/core/ports/primaryports"
	"net/http"
)

type AuthenticationHandler struct {
	authenticationService primaryports.AuthenticationService
}

func NewAuthenticationHandler(authenticationService primaryports.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{authenticationService: authenticationService}
}

func (h *AuthenticationHandler) SendMagicLink(w http.ResponseWriter, r *http.Request) {
	var args primaryports.SendMagicLinkArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		return
	}

	answer, err := h.authenticationService.SendMagicLink(args)

	if err != nil {
		switch err.Error() {
		case apperrors.InvalidEmail:
			httphelpers.WriteError(http.StatusBadRequest, "error", err.Error())(w, r)
		case apperrors.UserNotFound:
			httphelpers.WriteError(http.StatusNotFound, "error", err.Error())(w, r)
		case apperrors.EmailNotVerified:
			// todo: consider sending a new verification code
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

	httphelpers.WriteSuccess(http.StatusAccepted, "Authentication successful", answer)(w, r)
}

func (h *AuthenticationHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var args primaryports.GetAccessTokenArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		return
	}

	answer, err := h.authenticationService.Authenticate(args)

	if err != nil {
		switch err.Error() {
		case apperrors.InvalidCode:
			httphelpers.WriteError(http.StatusBadRequest, "error", err.Error())(w, r)
		case apperrors.AuthorizationCodeExpired:
			httphelpers.WriteError(http.StatusBadRequest, "error", err.Error())(w, r)
		case apperrors.UserNotFound:
			httphelpers.WriteError(http.StatusNotFound, "error", err.Error())(w, r)
		case apperrors.EmailNotVerified:
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

	httphelpers.WriteSuccess(http.StatusAccepted, "Authentication successful", answer)(w, r)
}
