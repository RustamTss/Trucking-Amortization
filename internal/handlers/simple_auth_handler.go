package handlers

import (
	"trucking-amortization/internal/models"
	"trucking-amortization/internal/services"
	"trucking-amortization/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type SimpleAuthHandler struct {
	userService *services.SimpleUserService
	jwtSecret   string
}

func NewSimpleAuthHandler(userService *services.SimpleUserService, jwtSecret string) *SimpleAuthHandler {
	return &SimpleAuthHandler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *SimpleAuthHandler) Register(c *fiber.Ctx) error {
	var req models.SimpleUserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Базовая валидация
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email, password and name are required",
		})
	}

	user, err := h.userService.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Генерируем JWT токен
	token, err := utils.GenerateToken(user.ID, user.Email, h.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.SimpleAuthResponse{
		User:  *user,
		Token: token,
	})
}

func (h *SimpleAuthHandler) Login(c *fiber.Ctx) error {
	var req models.SimpleUserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Базовая валидация
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	user, err := h.userService.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Генерируем JWT токен
	token, err := utils.GenerateToken(user.ID, user.Email, h.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	userResponse := models.SimpleUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Companies: user.Companies,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(models.SimpleAuthResponse{
		User:  userResponse,
		Token: token,
	})
}

func (h *SimpleAuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	user, err := h.userService.GetByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
