package service

import (
	"github.com/iamtomorrow/GEOApi/internal/model"
	"github.com/iamtomorrow/GEOApi/internal/repository"
)

func GetAllUsers() []model.User {
	return repository.Users
}
