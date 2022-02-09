package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kamchai-n/api-student-home-visit/handler/visit"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	repo "github.com/kamchai-n/api-student-home-visit/repository/visit"
	services "github.com/kamchai-n/api-student-home-visit/services/visit"
	"gorm.io/gorm"
)

func VisitRouter(router fiber.Router, db *gorm.DB) {
	visitRepository := repo.NewVisitRepository(db)
	visitServices := services.NewStudentService(visitRepository)
	visitHandler := handler.NewVisitHandler(visitServices)

	router.Use(middlewares.UserProtected())
	router.Get("/getAll", visitHandler.ListAllVisit)
	router.Get("/detail/:id", visitHandler.ListDetailById)
	router.Post("/add", visitHandler.SaveVisit)
	router.Put("/update", visitHandler.EditVisit)
	router.Delete("/delete/:visit_id", visitHandler.DeleteVisit)
}
