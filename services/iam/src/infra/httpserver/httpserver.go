package httpserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"iam/pkg/env"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "staging" {
		env.LoadFromFile()
	}
}

func Start() error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

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

		type NewUser struct {
			Id                  string    `json:"id"`
			Email               string    `json:"email"`
			Username            string    `json:"username"`
			UsernameFingerprint string    `json:"usernameFingerprint"`
			CreatedAt           time.Time `json:"createdAt"`
		}

		var newUser NewUser

		var registerRequest RegisterRequest
		_ = json.NewDecoder(r.Body).Decode(&registerRequest)

		fmt.Println("headers", r.Header.Get("X-Request-Id"), r.Header.Get("X-Initiator-Type"), r.Header.Get("X-Initiator-Id"))
		fmt.Println("body", registerRequest.Email, registerRequest.Username)

		newUser.Id = uuid.New().String()
		newUser.Email = registerRequest.Email
		newUser.Username = registerRequest.Username
		newUser.UsernameFingerprint = registerRequest.Username
		newUser.CreatedAt = time.Now().UTC()

		fmt.Println(newUser)

		_, err = db.Exec("INSERT INTO users(id, created_at, username, username_fingerprint, email) VALUES ($1, $2, $3, $4, $5)", newUser.Id, newUser.CreatedAt, newUser.Username, newUser.UsernameFingerprint, newUser.Email)

		if err != nil {
			fmt.Println(err)
			w.Write([]byte("{\"status\":\"error\",\"error\":{\"code\":\"already_exists\"}}"))
			return
		}

		w.Write([]byte("{\"status\":\"success\",\"data\":{\"userId\":\"" + newUser.Id + "\"}}"))
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
