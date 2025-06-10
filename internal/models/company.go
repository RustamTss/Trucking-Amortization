package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	OwnerID   primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type CompanyRequest struct {
	Name string `json:"name"`
}

type CompanyResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}
