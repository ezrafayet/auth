package httpserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"iam/pkg/envreader"
	"iam/pkg/httphelpers"
	"iam/src/infra/database"
	"iam/src/infra/emailprovider"
	"iam/src/infra/registry"
)

func init() {
	if os.Getenv("ENV") != "PRODUCTION" && os.Getenv("ENV") != "STAGING" {
		envreader.LoadFromFile()
	}

	if err := envreader.CheckRequiredEnv(); err != nil {
		panic(err)
	}
}

func Start() error {
	db := database.ConnectDB()

	defer database.CloseDB(db)

	emailProvider := emailprovider.NewEmailProvider()

	r := registry.NewRegistry(db, emailProvider)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// Route: GET /api/internal/v1/auth/whoami
	// It:    Provides detailed information about the authenticated user, or return anonymous.
	router.Get("/api/v1/auth/whoami", nil /*todo: implement*/)

	// Route: POST /api/internal/v1/auth/register
	// It:    Creates a new user.
	router.
		With(r.AuthorizationHandler.VerifyCaptcha).
		Post("/api/v1/auth/register", r.UsersHandler.Register)

	// Route: POST /api/internal/v1/auth/email-verification/send
	// It:    Sends a verification email to the user.
	//
	// /!\    This endpoint can issue a verification code through email
	router.Post("/api/v1/auth/email-verification/send", r.VerificationCodeHandler.SendVerificationEmail)

	// Route: PATCH /api/internal/v1/auth/email-verification/confirm
	// It:    Confirms the email address of the user.
	//
	// /!\    This endpoint can issue an authorization code in its answer
	router.Patch("/api/v1/auth/email-verification/confirm", r.VerificationCodeHandler.ConfirmEmail)

	// Route: POST /api/internal/v1/auth/magic-link
	// It:    Sends a magic link to the user.
	//
	// /!\    This endpoint can issue an authorization code via email
	router.
		With(r.AuthorizationHandler.VerifyFeatureFlags([]string{"ENABLE_AUTH_WITH_MAGIC_LINK"})).
		Post("/api/v1/auth/magic-link", r.AuthenticationHandler.SendMagicLink)

	// Route: POST /api/internal/v1/auth/token
	// It:    Authenticates the user by trading an access token and a refresh token an authorization code for.
	//
	// /!\    This endpoint can issue an access token and a refresh token in exchange for an authorization code
	// /!\    This is the only endpoint that can issue those tokens, in exchange for an authorization code
	router.Post("/api/v1/auth/token", r.AuthenticationHandler.Authenticate)

	// Route: POST /api/v1/auth/token/authorize
	// It:    Verifies the validity of an access token.
	router.
		With(r.AuthorizationHandler.VerifyAccessToken).
		Post("/api/v1/auth/token/authorize", func(w http.ResponseWriter, r *http.Request) {
			httphelpers.WriteSuccess(http.StatusAccepted, "Access token valid", struct{}{})(w, r)
		})

	// todo: logout

	// Route: POST /api/v1/auth/token/refresh
	// It:    Refreshes the access token of the user.
	//
	// /!\ This endpoint can issue an authorization code
	router.Post("/api/v1/auth/token/refresh", r.AuthorizationHandler.Refresh)

	router.
		With(r.AuthorizationHandler.VerifyAccessToken).
		Get("/api/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("headers", r.Header.Get("X-Request-Id"))
			httphelpers.WriteSuccess(http.StatusOK, "Foobar", struct{}{})(w, r)
		})

	router.HandleFunc("/*", httphelpers.WriteError(http.StatusNotFound, "error", "NOT_FOUND"))

	fmt.Println("Server started on port " + os.Getenv("PORT"))

	return http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
