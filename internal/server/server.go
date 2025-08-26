package server

import (
	"fmt"
	"net/http"

	"github.com/iamtomorrow/GEOApi/internal/routes"
)

func StartServer() {
	router := routes.SetupRoutes()

	fmt.Println("Server running at http://localhost:4000")
	http.ListenAndServe(":4000", router)
}
