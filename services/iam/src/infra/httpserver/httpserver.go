package httpserver

import (
	"encoding/json"
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
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.Post("/api/internal/v1/auth/register", func(w http.ResponseWriter, r *http.Request) {
		type RegisterRequest struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Method   string `json:"method"`
		}

		type RegisterResponse struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Data    struct {
				UserId string `json:"userId"`
			}
			Error struct {
				Code string `json:"code"`
			}
		}

		var registerRequest RegisterRequest
		_ = json.NewDecoder(r.Body).Decode(&registerRequest)

		fmt.Println("headers", r.Header.Get("X-Request-Id"), r.Header.Get("X-Initiator-Type"), r.Header.Get("X-Initiator-Id"))
		fmt.Println("body", registerRequest.Method, registerRequest.Email, registerRequest.Username)

		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.Post("/api/internal/v1/auth/users/email-verification/send", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.Patch("/api/internal/v1/auth/users/email-verification/confirm", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	// get a magic link
	r.Post("/api/internal/v1/auth/magic-link", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	// perform actual login
	r.Post("/api/internal/v1/auth/token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.Post("/api/internal/v1/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.Post("/api/internal/v1/auth/refresh", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	// This endpoint is used by the API Gateway to check if the access token is valid
	r.Get("/api/internal/v1/auth/authorize", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	r.Get("/api/internal/v1/auth/foobar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("headers", r.Header.Get("X-Request-Id"))
		w.Write([]byte("{\"status\":\"success\"}"))
	})

	fmt.Println("Server started on port 7777")
	return http.ListenAndServe(":7777", r)
}
