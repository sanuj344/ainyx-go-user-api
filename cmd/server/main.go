package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/sanuj344/ainyx-go-user-api/config"
	"github.com/sanuj344/ainyx-go-user-api/internal/handler"
	"github.com/sanuj344/ainyx-go-user-api/internal/logger"
	"github.com/sanuj344/ainyx-go-user-api/internal/middleware"
	"github.com/sanuj344/ainyx-go-user-api/internal/repository"
	"github.com/sanuj344/ainyx-go-user-api/internal/routes"
	"github.com/sanuj344/ainyx-go-user-api/internal/service"
)

func main() {
	cfg := config.Load()

	logr := logger.New()
	defer logr.Sync()

	ctx := context.Background()

	db := repository.NewDB(ctx, cfg.DBUrl)
	defer db.Close()

	repo := repository.NewRepository(db)
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logr))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	routes.Register(app, userHandler)

	log.Println("🚀 Server running on :3000")
	log.Fatal(app.Listen(":3000"))
}
