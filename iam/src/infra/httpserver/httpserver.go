package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"net/http"
)

var proxyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	targetURL := "http://iam:7777"

	// Create a new HTTP request with the same method, headers, and body as the incoming request
	proxyReq, err := http.NewRequest(r.Method, targetURL+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyReq.Header = r.Header

	// Send the request to the target URL using the default HTTP client
	client := http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response status, headers, and body from the target URL to the response writer
	w.WriteHeader(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	io.Copy(w, resp.Body)
})

func Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome1"))
	})
	r.Get("/api/auth/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome2"))
	})

	fmt.Println("Server started on port 7777")
	return http.ListenAndServe(":7777", r)
}
