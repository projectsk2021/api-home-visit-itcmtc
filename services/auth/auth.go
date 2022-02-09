package services

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type AuthService interface {
	NewUser(model.User) (model.User, error)
	UserLogin(domain.UserLoginForm) (string, error)
	ShowUser(uuid.UUID) (model.User, error)
	ShowUserAll() ([]model.User, error)
	RemoveUser(uuid.UUID) error
}
