package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type product struct {
	ID   int
	Name string
}

func main() {
	r := chi.NewRouter()

	r.Use(myMiddleware)

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

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var product product
		request, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("erro ao ler corpo da request"))
			return
		}

		err = json.Unmarshal(request, &product)
		if err != nil {
			w.Write([]byte("erro ao deserializar json"))
			return
		}

		w.Write([]byte(product.Name))
	})

	http.ListenAndServe(":8080", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before")
		next.ServeHTTP(w, r)
		fmt.Println("after")
	})
}
