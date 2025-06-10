package services

import (
	"errors"
	"time"
	"trucking-amortization/internal/models"
	"trucking-amortization/pkg/utils"
)

// Простая in-memory реализация для демонстрации
type SimpleUserService struct {
	users []models.SimpleUser
}

func NewSimpleUserService() *SimpleUserService {
	return &SimpleUserService{
		users: make([]models.SimpleUser, 0),
	}
}

func (s *SimpleUserService) Register(req *models.SimpleUserRegisterRequest) (*models.SimpleUserResponse, error) {
	// Проверяем, существует ли пользователь с таким email
	for _, user := range s.users {
		if user.Email == req.Email {
			return nil, errors.New("user with this email already exists")
		}
	}

	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Создаем нового пользователя
	user := models.SimpleUser{
		ID:        "user_" + req.Email, // Простой ID для демонстрации
		Email:     req.Email,
		Password:  hashedPassword,
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	s.users = append(s.users, user)

	return &models.SimpleUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *SimpleUserService) Login(req *models.SimpleUserLoginRequest) (*models.SimpleUser, error) {
	for _, user := range s.users {
		if user.Email == req.Email {
			// Проверяем пароль
			if !utils.CheckPasswordHash(req.Password, user.Password) {
				return nil, errors.New("invalid credentials")
			}
			return &user, nil
		}
	}
	return nil, errors.New("invalid credentials")
}

func (s *SimpleUserService) GetByID(userID string) (*models.SimpleUserResponse, error) {
	for _, user := range s.users {
		if user.ID == userID {
			return &models.SimpleUserResponse{
				ID:        user.ID,
				Email:     user.Email,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *SimpleUserService) AddCompanyToUser(userID string, companyID string) error {
	// Implementation needed
	return errors.New("method not implemented")
}
