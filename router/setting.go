package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kamchai-n/api-student-home-visit/handler/setting"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	repo "github.com/kamchai-n/api-student-home-visit/repository/setting"
	services "github.com/kamchai-n/api-student-home-visit/services/setting"
	"gorm.io/gorm"
)

func SettingRouter(router fiber.Router, db *gorm.DB) {
	settingRepository := repo.NewSettingRepository(db)
	settingServices := services.NewSettingService(settingRepository)
	settingHandler := handler.NewSettingHandler(settingServices)

	router.Use(middlewares.UserProtected())
	router.Get("/", settingHandler.ShowSetting)
	router.Put("/", settingHandler.EditIsFromSdcSetting)
}
