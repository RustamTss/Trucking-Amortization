package services

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"trucking-amortization/internal/models"
	"trucking-amortization/pkg/utils"
)

type SimpleUserService struct {
	users     map[string]*models.SimpleUser
	mutex     sync.RWMutex
	idCounter int
}

func NewSimpleUserService() *SimpleUserService {
	return &SimpleUserService{
		users:     make(map[string]*models.SimpleUser),
		mutex:     sync.RWMutex{},
		idCounter: 1,
	}
}

func (s *SimpleUserService) Register(req *models.SimpleUserRegisterRequest) (*models.SimpleUserResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	userID := fmt.Sprintf("%d", s.idCounter)
	s.idCounter++

	user := &models.SimpleUser{
		ID:        userID,
		Email:     req.Email,
		Password:  hashedPassword,
		Name:      req.Name,
		Companies: []string{},
		CreatedAt: time.Now(),
	}

	s.users[userID] = user

	return &models.SimpleUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Companies: user.Companies,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *SimpleUserService) Login(req *models.SimpleUserLoginRequest) (*models.SimpleUser, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Ищем пользователя по email
	for _, user := range s.users {
		if user.Email == req.Email {
			// Проверяем пароль
			if !utils.CheckPasswordHash(req.Password, user.Password) {
				return nil, errors.New("invalid credentials")
			}
			return user, nil
		}
	}

	return nil, errors.New("invalid credentials")
}

func (s *SimpleUserService) GetByID(userID string) (*models.SimpleUserResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}

	return &models.SimpleUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Companies: user.Companies,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *SimpleUserService) AddCompanyToUser(userID string, companyID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user, exists := s.users[userID]
	if !exists {
		return errors.New("user not found")
	}

	// Проверяем, не добавлена ли уже компания
	for _, id := range user.Companies {
		if id == companyID {
			return nil // Уже добавлена
		}
	}

	user.Companies = append(user.Companies, companyID)
	return nil
}
