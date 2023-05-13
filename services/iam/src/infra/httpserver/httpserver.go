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
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "staging" {
		env.LoadFromFile()
	}
}

func Start() error {
	db := database.ConnectDB()
	defer database.CloseDB(db)

	reg := registry.NewRegistry(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// User management endpoints
	r.Get("/api/internal/v1/auth/whoami", nil)
	r.Post("/api/internal/v1/auth/register", reg.UsersHandler.RegisterHandler)

	// Email verification endpoints
	r.Post("/api/internal/v1/auth/email-verification/send", nil)
	r.Patch("/api/internal/v1/auth/email-verification/verify", nil)

	// Authentication endpoints
	r.Post("/api/internal/v1/auth/magic-link", nil)
	r.Post("/api/internal/v1/auth/token", nil)
	r.Post("/api/internal/v1/auth/token/authorize", nil)
	r.Post("/api/internal/v1/auth/token/refresh", nil)
	r.Post("/api/internal/v1/auth/token/revoke", nil)

	// Other
	r.Get("/api/internal/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.HandleFunc("/*", httphelpers.WriteError(http.StatusNotFound, "error", "NOT_FOUND"))

	fmt.Println("Server started on port 7777")
	return http.ListenAndServe(":7777", r)
}
