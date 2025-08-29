package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iamtomorrow/GEOApi/internal/model"
)

func GetAllLocations(w http.ResponseWriter, r *http.Request) {

	response := []model.Location{
		{
			ID:          "1100809",
			IDMunicipio: "1100809",
			Name:        "Candeias do Jamari",
			UF:          "Rondônia",
			Centroide:   "POINT(-63.3254198532114 -8.88702392955617)",
		},
		{
			ID:          "1600212",
			IDMunicipio: "1600212",
			Name:        "Cutias",
			UF:          "Amapá",
			Centroide:   "POINT(-50.4839927545885 1.03057297634923)",
		},
		{
			ID:          "1100809",
			IDMunicipio: "1100809",
			Name:        "Candeias do Jamari",
			UF:          "Rondônia",
			Centroide:   "POINT(-63.3254198532114 -8.88702392955617)",
		},
		{
			ID:          "1100809",
			IDMunicipio: "1100809",
			Name:        "Candeias do Jamari",
			UF:          "Rondônia",
			Centroide:   "POINT(-63.3254198532114 -8.88702392955617)",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	url := r.URL
	header := r.Header

	response := model.DefaultModel{
		ID:     id,
		URL:    url,
		Header: header,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
