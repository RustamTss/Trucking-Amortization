package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asset struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CompanyID     primitive.ObjectID `bson:"company_id" json:"company_id"`
	Type          string             `bson:"type" json:"type"`
	Make          string             `bson:"make" json:"make"`
	Model         string             `bson:"model" json:"model"`
	Year          int                `bson:"year" json:"year"`
	VIN           string             `bson:"vin" json:"vin"`
	PurchaseDate  time.Time          `bson:"purchase_date" json:"purchase_date"`
	PurchasePrice float64            `bson:"purchase_price" json:"purchase_price"`
	LoanInfo      *LoanInfo          `bson:"loan_info,omitempty" json:"loan_info"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
}

type LoanInfo struct {
	LoanAmount     float64   `bson:"loan_amount" json:"loan_amount"`
	InterestRate   float64   `bson:"interest_rate" json:"interest_rate"`
	LoanTerm       int       `bson:"loan_term" json:"loan_term"`
	StartDate      time.Time `bson:"start_date" json:"start_date"`
	MonthlyPayment float64   `bson:"monthly_payment" json:"monthly_payment"`
}

type AssetRequest struct {
	CompanyID     string    `json:"company_id"`
	Type          string    `json:"type"`
	Make          string    `json:"make"`
	Model         string    `json:"model"`
	Year          int       `json:"year"`
	VIN           string    `json:"vin"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice float64   `json:"purchase_price"`
	LoanInfo      *LoanInfo `json:"loan_info"`
}

type AssetResponse struct {
	ID            string    `json:"id"`
	CompanyID     string    `json:"company_id"`
	Type          string    `json:"type"`
	Make          string    `json:"make"`
	Model         string    `json:"model"`
	Year          int       `json:"year"`
	VIN           string    `json:"vin"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice float64   `json:"purchase_price"`
	LoanInfo      *LoanInfo `json:"loan_info"`
	CreatedAt     time.Time `json:"created_at"`
}
