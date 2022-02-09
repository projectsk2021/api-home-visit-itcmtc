package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kamchai-n/api-student-home-visit/handler/static"
	repo "github.com/kamchai-n/api-student-home-visit/repository/static"
	services "github.com/kamchai-n/api-student-home-visit/services/static"
	"gorm.io/gorm"
)

func StaticRouter(router fiber.Router, db *gorm.DB) {
	staticRepository := repo.NewStaticRepository(db)
	staticServices := services.NewStaticService(staticRepository)
	staticHandler := handler.NewStaticHandler(staticServices)

	// router.Use(middlewares.UserProtected())
	router.Get("/education", staticHandler.ShowEducation)
}
