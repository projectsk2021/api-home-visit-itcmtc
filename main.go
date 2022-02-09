package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kamchai-n/api-student-home-visit/model"
	"github.com/kamchai-n/api-student-home-visit/router"
	"github.com/spf13/viper"
)

func main() {

	// Loading Config
	initConfig()

	// Loading Database
	DB := model.InitDatabase()

	// Init fiber
	app := fiber.New(fiber.Config{
		AppName: "Home Visit ITCMTC v1",
	})
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api := app.Group("/api")

	router.StudentRouter(api.Group("/student"), DB)
	router.AuthRouter(api.Group("/auth"), DB)
	router.SettingRouter(api.Group("/setting"), DB)
	router.VisitRouter(api.Group("/visit"), DB)
	router.StaticRouter(api.Group("/static"), DB)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("post")))
}

func initConfig() {
	if os.Getenv("RUN_MODE") == "production" {
		fmt.Println("production")
		viper.SetConfigName("production")
	} else {
		fmt.Println("development")
		viper.SetConfigName("development")
	}
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
