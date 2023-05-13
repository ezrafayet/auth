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

	//r.Get("/api/internal/v1/auth/whoami", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})

	r.Post("/api/internal/v1/auth/register", reg.UsersHandler.RegisterHandler)

	//r.Post("/api/internal/v1/auth/users/email-verification/send", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//r.Patch("/api/internal/v1/auth/users/email-verification/verify", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//// get a magic link
	//r.Post("/api/internal/v1/auth/magic-link", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//// perform actual login
	//r.Post("/api/internal/v1/auth/token", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//r.Post("/api/internal/v1/auth/logout", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//r.Post("/api/internal/v1/auth/refresh", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//// This endpoint is used by the API Gateway to check if the access token is valid
	//r.Get("/api/internal/v1/auth/authorize", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})
	//
	//r.Get("/api/internal/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("headers", r.Header.Get("X-Request-Id"))
	//	w.Write([]byte("{\"status\":\"success\"}"))
	//})

	r.HandleFunc("/*", httphelpers.WriteError(http.StatusNotFound, "error", "NOT_FOUND"))

	fmt.Println("Server started on port 7777")
	return http.ListenAndServe(":7777", r)
}
