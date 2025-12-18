package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"assignment/config"
	"assignment/internal/handler"
	"assignment/internal/logger"
	"assignment/internal/middleware"
	"assignment/internal/repository"
	"assignment/internal/routes"
)

func main() {
	logger.Init()
	defer logger.Sync()

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	db := config.NewDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	routes.RegisterUserRoutes(app, userHandler)

	log.Println("Running at 8000")
	log.Fatal(app.Listen(":8000"))
}
