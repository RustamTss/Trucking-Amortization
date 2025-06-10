package main

import (
	"log"
	"trucking-amortization/internal/config"
	"trucking-amortization/internal/handlers"
	"trucking-amortization/internal/middleware"
	"trucking-amortization/internal/models"
	"trucking-amortization/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Инициализируем сервисы (упрощенные версии)
	userService := services.NewSimpleUserService()
	calcService := services.NewCalculationService()

	// Инициализируем обработчики
	authHandler := handlers.NewSimpleAuthHandler(userService, cfg.JWTSecret)

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
			"message": "Trucking Amortization API is running (Simple Version)",
		})
	})

	// Демонстрационный endpoint для расчета амортизации
	app.Get("/demo/amortization", func(c *fiber.Ctx) error {
		// Пример расчета амортизации
		demoAsset := &models.SimpleAsset{
			ID:            "demo-1",
			Type:          "truck",
			Make:          "Freightliner",
			Model:         "Cascadia",
			Year:          2020,
			PurchasePrice: 150000,
			LoanInfo: &models.LoanInfo{
				LoanAmount:   120000,
				InterestRate: 5.5,
				LoanTerm:     60, // 5 лет
			},
		}

		// Здесь будет расчет амортизации
		_ = calcService
		_ = demoAsset

		return c.JSON(fiber.Map{
			"message": "Demo amortization calculation",
			"asset":   demoAsset,
		})
	})

	log.Printf("Server starting on port %s (Simple Version)", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
