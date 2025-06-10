package main

import (
	"log"
	"trucking-amortization/internal/config"
	"trucking-amortization/internal/handlers"
	"trucking-amortization/internal/middleware"
	"trucking-amortization/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Инициализируем простые сервисы (без MongoDB)
	userService := services.NewSimpleUserService()

	// Инициализируем обработчики
	authHandler := handlers.NewSimpleAuthHandler(userService, cfg.JWTSecret)

	// Создаем Fiber приложение
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// API маршруты
	api := app.Group("/api/v1")

	// Аутентификация (публичные маршруты)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Защищенные маршруты
	protected := api.Group("/", middleware.AuthMiddleware(cfg.JWTSecret))
	protected.Get("/profile", authHandler.GetProfile)

	// Здоровье приложения
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "trucking-amortization-simple",
		})
	})

	// Запускаем сервер
	log.Printf("Simple server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
