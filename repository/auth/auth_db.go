package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepository{db: db}
}

func (r authRepository) CreateUser(user model.User) (model.User, error) {
	if err := r.db.Model(&model.User{}).Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r authRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r authRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	if err := r.db.Select("user_id, name, username, images").Model(&model.User{}).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r authRepository) GetUserById(ID uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.db.Select("user_id, name, username, images").Model(&model.User{}).Where("user_id = ?", ID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r authRepository) DeleteUser(userId uuid.UUID) error {
	if err := r.db.Where("user_id = ?", userId).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
