package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iamtomorrow/GEOApi/internal/model"
)

type Status struct {
	Status string
}

func SingUp(w http.ResponseWriter, r *http.Request) {
	var response = model.Status{
		Status: "Ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((response))
}
