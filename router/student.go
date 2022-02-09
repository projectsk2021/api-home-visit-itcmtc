package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/kamchai-n/api-student-home-visit/handler/student"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	repoSetting "github.com/kamchai-n/api-student-home-visit/repository/setting"
	repoStudent "github.com/kamchai-n/api-student-home-visit/repository/student"
	services "github.com/kamchai-n/api-student-home-visit/services/student"
	"gorm.io/gorm"
)

func StudentRouter(router fiber.Router, db *gorm.DB) {
	studentRepository := repoStudent.NewStudentRepository(db)
	settingRepository := repoSetting.NewSettingRepository(db)
	studentServices := services.NewStudentService(studentRepository, settingRepository)
	studentHandler := handler.NewStudentHandler(studentServices)
	router.Post("/add", studentHandler.SaveStudent)
	router.Use(middlewares.UserProtected())
	router.Post("/add-from-excel", studentHandler.SaveStudentFromExcel)
	router.Post("/getAll", studentHandler.ShowInfoStudent)
	router.Post("/getOne", studentHandler.ShowInfoStudentById)
	router.Post("/delete", studentHandler.DeleteStudentById)
}
