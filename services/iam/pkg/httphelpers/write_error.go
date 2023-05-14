package httphelpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteError is a helper function to write an error to the response writer
func WriteError(code int, status, errorCode string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		answerJson, err := json.Marshal(ApiAnswer{
			Status: status,
			Error: &AnswerError{
				Code: errorCode,
			},
		})

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"status":"error","message":"Internal Server Error"}`))
			return
		}

		w.WriteHeader(code)

		_, err = w.Write(answerJson)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"status":"error","message":"Internal Server Error"}`))
			return
		}
	}
}
