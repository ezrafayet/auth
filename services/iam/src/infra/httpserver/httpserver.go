package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"iam/pkg/envreader"
	"iam/pkg/httphelpers"
	"iam/src/infra/database"
	"iam/src/infra/emailprovider"
	"iam/src/infra/registry"
	"net/http"
	"os"
)

func init() {
	if os.Getenv("ENV") != "PRODUCTION" && os.Getenv("ENV") != "STAGING" {
		envreader.LoadFromFile()
	}

	err := envreader.CheckRequiredEnv()

	if err != nil {
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

	// User management endpoints

	router.Get("/api/internal/v1/auth/whoami", nil)

	// /!\ This endpoint can issue a userId
	// It might become an endpoint for magic links only
	router.Post("/api/internal/v1/auth/register", r.UsersHandler.Register)

	// Email verification endpoints

	// /!\ This endpoint can issue a verification code
	router.Post("/api/internal/v1/auth/email-verification/send", r.VerificationCodeHandler.SendVerificationEmail)

	// /!\ This endpoint can issue an authorization code
	router.Patch("/api/internal/v1/auth/email-verification/confirm", r.VerificationCodeHandler.ConfirmEmail)

	// Authentication endpoints

	// /!\ This endpoint can issue an authorization code
	// /!\ This endpoint can issue a userId if email is not verified
	router.Post("/api/internal/v1/auth/magic-link", r.AuthenticationHandler.SendMagicLink)

	// /!\ This endpoint can issue an access and a refresh token
	// /!\ This is the only endpoint that can issue those tokens, in exchange for an authorization code
	router.Post("/api/internal/v1/auth/token", r.AuthenticationHandler.Authenticate)

	router.With(r.AuthorizationHandler.CheckAccessTokenMiddleware).Post(
		"/api/internal/v1/auth/token/authorize", func(w http.ResponseWriter, r *http.Request) {
			httphelpers.WriteSuccess(http.StatusAccepted, "Authorized successfully", struct{}{})(w, r)
		})

	// /!\ This endpoint can issue an authorization code
	router.Post("/api/internal/v1/auth/token/refresh", nil)

	router.Post("/api/internal/v1/auth/token/revoke", nil)

	// Other endpoints

	router.With(r.AuthorizationHandler.CheckAccessTokenMiddleware).Get("/api/internal/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	router.HandleFunc("/*", httphelpers.WriteError(http.StatusNotFound, "error", "NOT_FOUND"))

	fmt.Println("Server started on port " + os.Getenv("PORT"))

	return http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
