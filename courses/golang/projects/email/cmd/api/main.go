package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		status := map[string]string{"status": "ok"}
		w.Header().Set("Content-Type", "application/json")

		resp, _ := json.Marshal(status)
		w.Write(resp)
	})

	// Path param
	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")

		w.Write([]byte(param))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		if param != "" {
			w.Write([]byte(param))
		} else {
			w.Write([]byte("hello"))
		}

	})

	http.ListenAndServe(":8080", r)
}
