package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/iamtomorrow/GEOApi/internal/handler"
)

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", handler.Ping)
	router.Get("/test", handler.Test)

	// locations
	router.Get("/locations", handler.GetAllLocations)
	router.Get("/location/{id}", handler.GetLocation)

	return router
}
