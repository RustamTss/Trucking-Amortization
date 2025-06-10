package models

import "time"

type AmortizationEntry struct {
	Month              int       `json:"month"`
	Payment            float64   `json:"payment"`
	Principal          float64   `json:"principal"`
	Interest           float64   `json:"interest"`
	Balance            float64   `json:"balance"`
	CumulativeInterest float64   `json:"cumulative_interest"`
	Date               time.Time `json:"date"`
}

type AmortizationSchedule struct {
	AssetID        string              `json:"asset_id"`
	TotalAmount    float64             `json:"total_amount"`
	MonthlyPayment float64             `json:"monthly_payment"`
	TotalInterest  float64             `json:"total_interest"`
	Entries        []AmortizationEntry `json:"entries"`
}

type DepreciationEntry struct {
	Year                    int       `json:"year"`
	DepreciationAmount      float64   `json:"depreciation_amount"`
	AccumulatedDepreciation float64   `json:"accumulated_depreciation"`
	BookValue               float64   `json:"book_value"`
	Date                    time.Time `json:"date"`
}

type DepreciationSchedule struct {
	AssetID            string              `json:"asset_id"`
	InitialValue       float64             `json:"initial_value"`
	UsefulLife         int                 `json:"useful_life"`
	AnnualDepreciation float64             `json:"annual_depreciation"`
	Entries            []DepreciationEntry `json:"entries"`
}

type BusinessDebtSummary struct {
	CompanyID       string  `json:"company_id"`
	TotalLoanAmount float64 `json:"total_loan_amount"`
	TotalBalance    float64 `json:"total_balance"`
	MonthlyPayment  float64 `json:"monthly_payment"`
	AssetsCount     int     `json:"assets_count"`
}

type BusinessDebtDetail struct {
	AssetID         string  `json:"asset_id"`
	AssetType       string  `json:"asset_type"`
	Make            string  `json:"make"`
	Model           string  `json:"model"`
	Year            int     `json:"year"`
	LoanAmount      float64 `json:"loan_amount"`
	CurrentBalance  float64 `json:"current_balance"`
	MonthlyPayment  float64 `json:"monthly_payment"`
	InterestRate    float64 `json:"interest_rate"`
	RemainingMonths int     `json:"remaining_months"`
}

type BusinessDebtSchedule struct {
	Summary BusinessDebtSummary  `json:"summary"`
	Details []BusinessDebtDetail `json:"details"`
}
