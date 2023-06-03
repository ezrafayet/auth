package handlers

import (
	"fmt"
	"iam/pkg/apperrors"
	"iam/pkg/httphelpers"
	"iam/src/core/ports/primaryports"
	"net/http"
	"os"
)

func (h *AuthorizationHandler) VerifyAccessToken(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		answer, err := h.authorizationService.IsAccessTokenValid(primaryports.IsAccessTokenValidArgs{
			AuthorisationHeader: r.Header.Get("Authorization"),
		})

		if err != nil {
			switch err.Error() {
			case "Token is expired":
				httphelpers.WriteError(http.StatusForbidden, "error", apperrors.InvalidAccessToken)(w, r)
			default:
				fmt.Println(err)
				httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
			}
			return
		}

		if answer.Valid {
			nextHandler.ServeHTTP(w, r)
			return
		}

		httphelpers.WriteError(http.StatusForbidden, "error", apperrors.InvalidAccessToken)(w, r)
	})
}

func (h *AuthorizationHandler) VerifyFeatureFlags(flags []string) func(nextHandler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, f := range flags {
				if os.Getenv(f) != "true" {
					httphelpers.WriteError(http.StatusNotImplemented, "error", apperrors.FeatureNotEnabled)(w, r)
					return
				}
			}
			nextHandler.ServeHTTP(w, r)
			return
		})
	}
}