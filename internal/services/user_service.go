package services

import (
	"context"
	"errors"
	"time"
	"trucking-amortization/internal/models"
	"trucking-amortization/pkg/database"
	"trucking-amortization/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	db         *database.MongoDB
	collection *mongo.Collection
}

func NewUserService(db *database.MongoDB) *UserService {
	return &UserService{
		db:         db,
		collection: db.GetCollection("users"),
	}
}

func (s *UserService) Register(req *models.UserRegisterRequest) (*models.UserResponse, error) {
	// Проверяем, существует ли пользователь с таким email
	var existingUser models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": req.Email}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}
	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Создаем нового пользователя
	user := &models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		Name:      req.Name,
		Companies: []primitive.ObjectID{},
		CreatedAt: time.Now(),
	}

	result, err := s.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return &models.UserResponse{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		Name:      user.Name,
		Companies: user.Companies,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *UserService) Login(req *models.UserLoginRequest) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Проверяем пароль
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (s *UserService) GetByID(userID string) (*models.UserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var user models.User
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		Name:      user.Name,
		Companies: user.Companies,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *UserService) AddCompanyToUser(userID string, companyID primitive.ObjectID) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	_, err = s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$addToSet": bson.M{"companies": companyID}},
	)
	return err
}
