package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"iam/pkg/env"
	"iam/pkg/httphelpers"
	"iam/src/infra/database"
	"iam/src/infra/registry"
	"net/http"
	"os"
)

func init() {
	if os.Getenv("ENV") != "PRODUCTION" && os.Getenv("ENV") != "STAGING" {
		env.LoadFromFile()
	}

	err := env.CheckRequiredEnv()

	if err != nil {
		panic(err)
	}
}

func Start() error {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	r := registry.NewRegistry(db)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// User management endpoints
	router.Get("/api/internal/v1/auth/whoami", nil)
	router.Post("/api/internal/v1/auth/register", r.UsersHandler.Register)

	// Email verification endpoints
	router.Post("/api/internal/v1/auth/email-verification/send", r.VerificationCodeHandler.SendVerificationEmail)
	router.Patch("/api/internal/v1/auth/email-verification/confirm", r.VerificationCodeHandler.VerifyEmail)

	// Authentication endpoints
	router.Post("/api/internal/v1/auth/magic-link", nil)
	router.Post("/api/internal/v1/auth/token", nil)
	router.Post("/api/internal/v1/auth/token/authorize", nil)
	router.Post("/api/internal/v1/auth/token/refresh", nil)
	router.Post("/api/internal/v1/auth/token/revoke", nil)

	// Other
	router.Get("/api/internal/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	router.HandleFunc("/*", httphelpers.WriteError(http.StatusNotFound, "error", "NOT_FOUND"))

	fmt.Println("Server started on port " + os.Getenv("PORT"))
	return http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
