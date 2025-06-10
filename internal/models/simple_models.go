package models

import "time"

// Упрощенные модели для демонстрации без MongoDB
type SimpleUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Companies []string  `json:"companies"`
	CreatedAt time.Time `json:"created_at"`
}

type SimpleCompany struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SimpleAsset struct {
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

type SimpleUserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SimpleUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SimpleUserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Companies []string  `json:"companies"`
	CreatedAt time.Time `json:"created_at"`
}

type SimpleAuthResponse struct {
	User  SimpleUserResponse `json:"user"`
	Token string             `json:"token"`
}
