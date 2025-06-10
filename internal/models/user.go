package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Email     string               `bson:"email" json:"email"`
	Password  string               `bson:"password" json:"-"`
	Name      string               `bson:"name" json:"name"`
	Companies []primitive.ObjectID `bson:"companies" json:"companies"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string               `json:"id"`
	Email     string               `json:"email"`
	Name      string               `json:"name"`
	Companies []primitive.ObjectID `json:"companies"`
	CreatedAt time.Time            `json:"created_at"`
}

type AuthResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
