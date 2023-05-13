package httphelpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteSuccess(code int, message string, data any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		answerJson, err := json.Marshal(ApiAnswer{
			Status:  "success",
			Message: message,
			Data:    data,
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
