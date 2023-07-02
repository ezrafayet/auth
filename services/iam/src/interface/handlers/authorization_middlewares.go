package handlers

import (
	"fmt"
	"iam/pkg/apperrors"
	"iam/pkg/httphelpers"
	"iam/src/core/ports/primaryports"
	"net/http"
)

// VerifyFeatureFlags refuses access if a feature or features are not enabled
func (h *AuthorizationHandler) VerifyFeatureFlags(flags []string) func(nextHandler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			answer, err := h.authorizationService.AreFeaturesEnabled(primaryports.AreFeaturesEnabledArgs{
				FlagsNeeded: flags,
			})

			if err != nil {
				fmt.Println(err)

				httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
			}

			if answer.Active == true {
				nextHandler.ServeHTTP(w, r)

				return
			}

			httphelpers.WriteError(http.StatusServiceUnavailable, "error", apperrors.FeatureDisabled)(w, r)
		})
	}
}

// VerifyAccessToken refuses access if there is no access token or if it is not valid
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

// VerifyPermissions refuses access if a given actor does not have required authorizations to perform an action on a given resource
func (h *AuthorizationHandler) VerifyPermissions() func(nextHandler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// todo: implement abac & policies

			nextHandler.ServeHTTP(w, r)

			return
		})
	}
}

func (h *AuthorizationHandler) VerifyCaptcha(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		answer, err := h.authorizationService.IsCaptchaValid(primaryports.IsCaptchaValidArgs{
			CaptchaResponse: r.Header.Get("X-Captcha-Response"),
		})

		if err != nil {
			fmt.Println(err)

			httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)

			return
		}

		if answer.Valid {
			nextHandler.ServeHTTP(w, r)

			return
		}

		httphelpers.WriteError(http.StatusForbidden, "error", apperrors.InvalidCaptcha)(w, r)

		return
	})
}
