package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type SelectLogFileRequest struct {
	Name string `json:"name"`
}

func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var s SelectLogFileRequest

		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		if s.Name == "" {
			http.Error(w, "Name field is required", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "name", s.Name)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
