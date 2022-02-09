package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type AuthRepository interface {
	CreateUser(model.User) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserById(uuid.UUID) (model.User, error)
	GetAllUser() ([]model.User, error)
	DeleteUser(uuid.UUID) error
}
