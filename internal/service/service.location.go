package service

import (
	"github.com/iamtomorrow/GEOApi/internal/model"
	"github.com/iamtomorrow/GEOApi/internal/repository"
)

func GetAllLocations() []model.Location {
	return repository.Locations
}
