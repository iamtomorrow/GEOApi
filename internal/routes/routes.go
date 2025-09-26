package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/iamtomorrow/GEOApi/internal/handler"
	"github.com/iamtomorrow/GEOApi/internal/middleware"
)

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()

	// testing
	router.Get("/v1/ping", handler.Ping)
	router.Get("/v1/test", handler.Test)

	// user
	router.Get("/v1/signup", middleware.Authenticate)

	// locations
	router.Get("/v1/timezone/{lat}{lon}", handler.Handler)
	router.Get("/v1/places/nearby/{lat}{lon}{radius}", handler.Handler)
	router.Get("/v1/geocode/{address}", handler.Handler)
	router.Get("/v1/address/{lat}{lon}", handler.Handler)
	router.Get("/v1/distance/{fromLat}{fromLon}{toLat}{toLon}{unit}", handler.Handler)

	router.Get("/v1/locations", handler.GetAllLocations)
	router.Get("/v1/location/{id}", handler.GetLocation)

	return router
}
