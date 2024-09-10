package fakers

import (
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/rizkiromadoni/go-shop/app/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "$2a$12$MHxK3cL3GdqHM5GszA4nfuZehioKvZ3qYO1X8yM8cy/AYuey7.Cdi", // user
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
