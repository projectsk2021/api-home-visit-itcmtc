package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
	repoAuth "github.com/kamchai-n/api-student-home-visit/repository/auth"
	repoSetting "github.com/kamchai-n/api-student-home-visit/repository/setting"
	"github.com/kamchai-n/api-student-home-visit/utils"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepo    repoAuth.AuthRepository
	settingRepo repoSetting.SettingRepository
}

func NewAuthService(authRepo repoAuth.AuthRepository, settingRepo repoSetting.SettingRepository) AuthService {
	return authService{authRepo: authRepo, settingRepo: settingRepo}
}

func (s authService) NewUser(req model.User) (model.User, error) {
	teacherPassword, _ := bcrypt.GenerateFromPassword([]byte(*req.Password), 10)
	teacherPasswordString := string(teacherPassword)
	req.Password = &teacherPasswordString

	user, err := s.authRepo.CreateUser(req)
	if err != nil {
		return user, fmt.Errorf("1310")
	}
	if err := s.settingRepo.CreateSetting(user.UserID); err != nil {
		return user, fmt.Errorf("1315")
	}

	user.Password = nil
	return req, nil
}

func (s authService) UserLogin(req domain.UserLoginForm) (string, error) {
	var token string
	user, err := s.authRepo.GetUserByUsername(req.Username)
	if err != nil {
		return token, fmt.Errorf("1206")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(req.Password)); err != nil {
		return token, fmt.Errorf("1314")
	}

	token, err = utils.CreateJWTUser(&user.UserID)
	if err != nil {
		return token, fmt.Errorf("1206")
	}

	return token, nil
}

func (s authService) ShowUserAll() ([]model.User, error) {
	users, err := s.authRepo.GetAllUser()
	if err != nil {
		return users, fmt.Errorf("1319")
	}
	return users, nil
}

func (s authService) ShowUser(id uuid.UUID) (model.User, error) {
	user, err := s.authRepo.GetUserById(id)
	if err != nil {
		return user, fmt.Errorf("1311")
	}
	return user, nil
}

func (s authService) RemoveUser(userId uuid.UUID) error {

	if err := s.authRepo.DeleteUser(userId); err != nil {
		return fmt.Errorf("1311")
	}
	return nil
}
