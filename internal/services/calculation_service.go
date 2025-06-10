package services

import (
	"math"
	"time"
	"trucking-amortization/internal/models"
)

type CalculationService struct{}

func NewCalculationService() *CalculationService {
	return &CalculationService{}
}

// CalculateAmortizationSchedule рассчитывает график погашения кредита
func (s *CalculationService) CalculateAmortizationSchedule(asset *models.Asset) *models.AmortizationSchedule {
	if asset.LoanInfo == nil {
		return nil
	}

	loan := asset.LoanInfo
	monthlyRate := loan.InterestRate / 100 / 12
	termMonths := float64(loan.LoanTerm)

	// Рассчитываем ежемесячный платеж по формуле аннуитета
	monthlyPayment := loan.LoanAmount * (monthlyRate * math.Pow(1+monthlyRate, termMonths)) /
		(math.Pow(1+monthlyRate, termMonths) - 1)

	entries := make([]models.AmortizationEntry, 0, loan.LoanTerm)
	balance := loan.LoanAmount
	cumulativeInterest := 0.0
	currentDate := loan.StartDate

	for month := 1; month <= loan.LoanTerm; month++ {
		interest := balance * monthlyRate
		principal := monthlyPayment - interest
		balance -= principal
		cumulativeInterest += interest

		// Корректируем последний платеж, если есть незначительный остаток
		if month == loan.LoanTerm && balance > 0 {
			principal += balance
			balance = 0
		}

		entry := models.AmortizationEntry{
			Month:              month,
			Payment:            monthlyPayment,
			Principal:          principal,
			Interest:           interest,
			Balance:            balance,
			CumulativeInterest: cumulativeInterest,
			Date:               currentDate,
		}
		entries = append(entries, entry)

		// Переходим к следующему месяцу
		currentDate = currentDate.AddDate(0, 1, 0)
	}

	return &models.AmortizationSchedule{
		AssetID:        asset.ID.Hex(),
		TotalAmount:    loan.LoanAmount,
		MonthlyPayment: monthlyPayment,
		TotalInterest:  cumulativeInterest,
		Entries:        entries,
	}
}

// CalculateDepreciationSchedule рассчитывает график амортизации актива (прямолинейный метод)
func (s *CalculationService) CalculateDepreciationSchedule(asset *models.Asset) *models.DepreciationSchedule {
	// Для траков и трейлеров обычно используется срок полезного использования 5-7 лет
	usefulLife := 7
	if asset.Type == "trailer" {
		usefulLife = 10 // Трейлеры служат дольше
	}

	// Остаточная стоимость 10% от первоначальной стоимости
	salvageValue := asset.PurchasePrice * 0.1
	depreciableAmount := asset.PurchasePrice - salvageValue
	annualDepreciation := depreciableAmount / float64(usefulLife)

	entries := make([]models.DepreciationEntry, 0, usefulLife)
	accumulatedDepreciation := 0.0
	currentDate := asset.PurchaseDate

	for year := 1; year <= usefulLife; year++ {
		accumulatedDepreciation += annualDepreciation
		bookValue := asset.PurchasePrice - accumulatedDepreciation

		entry := models.DepreciationEntry{
			Year:                    year,
			DepreciationAmount:      annualDepreciation,
			AccumulatedDepreciation: accumulatedDepreciation,
			BookValue:               bookValue,
			Date:                    time.Date(currentDate.Year()+year, currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location()),
		}
		entries = append(entries, entry)
	}

	return &models.DepreciationSchedule{
		AssetID:            asset.ID.Hex(),
		InitialValue:       asset.PurchasePrice,
		UsefulLife:         usefulLife,
		AnnualDepreciation: annualDepreciation,
		Entries:            entries,
	}
}

// CalculateCurrentLoanBalance рассчитывает текущий остаток по кредиту
func (s *CalculationService) CalculateCurrentLoanBalance(asset *models.Asset) float64 {
	if asset.LoanInfo == nil {
		return 0
	}

	loan := asset.LoanInfo
	monthlyRate := loan.InterestRate / 100 / 12
	termMonths := float64(loan.LoanTerm)

	// Количество прошедших месяцев с начала кредита
	monthsPassed := int(time.Since(loan.StartDate).Hours() / 24 / 30.44) // Среднее количество дней в месяце
	if monthsPassed <= 0 {
		return loan.LoanAmount
	}
	if monthsPassed >= loan.LoanTerm {
		return 0
	}

	// Рассчитываем ежемесячный платеж
	monthlyPayment := loan.LoanAmount * (monthlyRate * math.Pow(1+monthlyRate, termMonths)) /
		(math.Pow(1+monthlyRate, termMonths) - 1)

	// Рассчитываем остаток после определенного количества платежей
	balance := loan.LoanAmount
	for i := 0; i < monthsPassed; i++ {
		interest := balance * monthlyRate
		principal := monthlyPayment - interest
		balance -= principal
	}

	if balance < 0 {
		balance = 0
	}

	return balance
}

// CalculateRemainingMonths рассчитывает количество оставшихся месяцев до погашения кредита
func (s *CalculationService) CalculateRemainingMonths(asset *models.Asset) int {
	if asset.LoanInfo == nil {
		return 0
	}

	loan := asset.LoanInfo
	monthsPassed := int(time.Since(loan.StartDate).Hours() / 24 / 30.44)
	remaining := loan.LoanTerm - monthsPassed

	if remaining < 0 {
		return 0
	}

	return remaining
}
