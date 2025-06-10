package main

import (
	"log"
	"trucking-amortization/internal/config"
	"trucking-amortization/internal/handlers"
	"trucking-amortization/internal/middleware"
	"trucking-amortization/internal/services"
	"trucking-amortization/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к MongoDB
	db, err := database.NewMongoDB(cfg.MongoURI, cfg.DatabaseName)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Close()

	// Инициализируем сервисы
	userService := services.NewUserService(db)
	calcService := services.NewCalculationService()

	// Инициализируем обработчики
	authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)

	// Создаем приложение Fiber
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
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Маршруты
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Get("/me", middleware.AuthMiddleware(cfg.JWTSecret), authHandler.Me)

	// Companies routes (будут добавлены позже)
	companies := api.Group("/companies")
	companies.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	// Assets routes (будут добавлены позже)
	assets := api.Group("/assets")
	assets.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	// Schedules routes (будут добавлены позже)
	schedules := api.Group("/schedules")
	schedules.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Trucking Amortization API is running",
		})
	})

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
