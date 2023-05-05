package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/internal/v1/auth/whoami", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed whoami"))
	})

	r.Post("/api/internal/v1/auth/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed register"))
	})

	r.Post("/api/internal/v1/auth/users/{userId}/email-verification", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed email verification sending"))
	})

	r.Patch("/api/internal/v1/auth/users/{userId}/email-verification/{verificationCode}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed email verification code"))
	})

	r.Post("/api/internal/v1/auth/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed login"))
	})

	r.Post("/api/internal/v1/auth/refresh", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed refreshing token"))
	})

	r.Get("/api/internal/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("accessed foobar"))
	})

	fmt.Println("Server started on port 7777")
	return http.ListenAndServe(":7777", r)
}
