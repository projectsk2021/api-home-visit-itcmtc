package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kamchai-n/api-student-home-visit/handler/auth"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	repoAuth "github.com/kamchai-n/api-student-home-visit/repository/auth"
	repoSetting "github.com/kamchai-n/api-student-home-visit/repository/setting"
	services "github.com/kamchai-n/api-student-home-visit/services/auth"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {
	authRepository := repoAuth.NewAuthRepository(db)
	settingRepository := repoSetting.NewSettingRepository(db)
	authServices := services.NewAuthService(authRepository, settingRepository)
	studentHandler := handler.NewAuthHandler(authServices)

	router.Post("/login", studentHandler.UserLogin)
	router.Use(middlewares.UserProtected())
	router.Post("/register", studentHandler.UserRegister)
	router.Get("/getAll", studentHandler.GetAllUser)
	router.Get("/info", studentHandler.UserInfo)
	router.Delete("/deleteUser/:id", studentHandler.DeleteUser)
}
