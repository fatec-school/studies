package main

import (
	"email/internal/adapter/dto"
	"email/internal/domain/campaign/service"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type product struct {
	ID   int
	Name string
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := service.Service{}
	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request dto.NewCampaignRequest
		render.DecodeJSON(r.Body, &request)
		id, err := service.Create(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusCreated)
	})

	http.ListenAndServe(":8080", r)
}
