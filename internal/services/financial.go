package services

import (
	"database/sql"
	"log"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/models"
	"github.com/remithomasn7/qualitinvest/internal/repository"
)

type FinancialService struct {
	apiClient     *alpha_vantage.AlphaVantageClient
	companyRepo   repository.CompanyRepository
	overviewRepo  repository.CompanyOverviewRepository
	incomeRepo    repository.IncomeRepository
	balanceRepo   repository.BalanceRepository
	cashFlowRepo  repository.CashFlowRepository
	dividendsRepo repository.DividendsRepository
	splitsRepo    repository.SplitsRepository
	sharesRepo    repository.SharesOutstandingRepository
	earningsRepo  repository.EarningsRepository
}

func NewFinancialService(db *sql.DB, apiClient *alpha_vantage.AlphaVantageClient) *FinancialService {
	return &FinancialService{
		apiClient:     apiClient,
		companyRepo:   repository.NewCompanyRepository(db),
		overviewRepo:  repository.NewCompanyOverviewRepository(db),
		incomeRepo:    repository.NewIncomeRepository(db),
		balanceRepo:   repository.NewBalanceRepository(db),
		cashFlowRepo:  repository.NewCashFlowRepository(db),
		dividendsRepo: repository.NewDividendsRepository(db),
		splitsRepo:    repository.NewSplitsRepository(db),
		sharesRepo:    repository.NewSharesOutstandingRepository(db),
		earningsRepo:  repository.NewEarningsRepository(db),
	}
}

func (s *FinancialService) FetchCompanyOverview(symbol string) (models.CompanyOverview, error) {
	// Try to get from database first
	if dbData, err := s.overviewRepo.GetBySymbol(symbol); err == nil && dbData != nil {
		log.Printf("Overview data successfully found in database for: %s", symbol)
		return *dbData, nil
	}

	// If not in database, call AlphaVantage API
	log.Printf("Overview data not found in database for %s, calling AlphaVantage API", symbol)
	overviewData, err := s.apiClient.CompanyOverview(symbol)
	if err != nil {
		log.Printf("Failed to fetch overview from AlphaVantage API for %s: %v", symbol, err)
		return models.CompanyOverview{}, err
	}

	// Create or update company record
	_, err = s.companyRepo.CreateOrUpdate(
		overviewData.Symbol,
		overviewData.Name,
		overviewData.Exchange,
		overviewData.Sector,
		overviewData.Industry,
		overviewData.Country,
		overviewData.Currency,
	)
	if err != nil {
		log.Printf("Warning: Failed to save company info for %s: %v", symbol, err)
	}

	// Persist overview data to database
	err = s.overviewRepo.Save(*overviewData)
	if err != nil {
		log.Printf("Warning: Failed to persist overview data for %s: %v", symbol, err)
		// Continue anyway - we can still return the data
	}

	log.Printf("Overview data saved to database for: %s", symbol)
	return *overviewData, nil
}

func FetchETFProfile(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.ETFProfile, error) {
	log.Printf("Fetching ETF Profile from AlphaVantage API for: %s", symbol)
	etfProfileData, err := apiClient.ETFProfile(symbol)
	if err != nil {
		log.Printf("Failed to fetch ETF Profile from AlphaVantage API for %s: %v", symbol, err)
		return models.ETFProfile{}, err
	}

	log.Printf("ETF Profile data successfully fetched for: %s", symbol)
	return *etfProfileData, nil
}

func FetchDividends(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.Dividends, error) {
	log.Printf("Fetching Dividends from AlphaVantage API for: %s", symbol)
	dividendsData, err := apiClient.Dividends(symbol)
	if err != nil {
		log.Printf("Failed to fetch Dividends from AlphaVantage API for %s: %v", symbol, err)
		return models.Dividends{}, err
	}

	log.Printf("Dividends data successfully fetched for: %s", symbol)
	return *dividendsData, nil
}

func FetchSplits(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.Splits, error) {
	log.Printf("Fetching Splits from AlphaVantage API for: %s", symbol)
	splitsData, err := apiClient.Splits(symbol)
	if err != nil {
		log.Printf("Failed to fetch Splits from AlphaVantage API for %s: %v", symbol, err)
		return models.Splits{}, err
	}

	log.Printf("Splits data successfully fetched for: %s", symbol)
	return *splitsData, nil
}

func FetchIncomeStatements(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.IncomeStatements, error) {
	log.Printf("Fetching Income Statements from AlphaVantage API for: %s", symbol)
	incomeData, err := apiClient.IncomeStatements(symbol)
	if err != nil {
		log.Printf("Failed to fetch Income Statements from AlphaVantage API for %s: %v", symbol, err)
		return models.IncomeStatements{}, err
	}

	log.Printf("Income Statements data successfully fetched for: %s", symbol)
	return *incomeData, nil
}

func FetchBalanceSheet(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.BalanceSheet, error) {
	log.Printf("Fetching Balance Sheet from AlphaVantage API for: %s", symbol)
	balanceSheetData, err := apiClient.BalanceSheet(symbol)
	if err != nil {
		log.Printf("Failed to fetch Balance Sheet from AlphaVantage API for %s: %v", symbol, err)
		return models.BalanceSheet{}, err
	}

	log.Printf("Balance Sheet data successfully fetched for: %s", symbol)
	return *balanceSheetData, nil
}

func FetchCashFlowStatements(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.CashFlowStatements, error) {
	log.Printf("Fetching Cash Flow Statements from AlphaVantage API for: %s", symbol)
	cashFlowData, err := apiClient.CashFlowStatements(symbol)
	if err != nil {
		log.Printf("Failed to fetch Cash Flow Statements from AlphaVantage API for %s: %v", symbol, err)
		return models.CashFlowStatements{}, err
	}

	log.Printf("Cash Flow Statements data successfully fetched for: %s", symbol)
	return *cashFlowData, nil
}

func FetchSharesOutstandings(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.SharesOutstandings, error) {
	log.Printf("Fetching Shares Outstanding from AlphaVantage API for: %s", symbol)
	sharesOutstandingsData, err := apiClient.SharesOutstandings(symbol)
	if err != nil {
		log.Printf("Failed to fetch Shares Outstanding from AlphaVantage API for %s: %v", symbol, err)
		return models.SharesOutstandings{}, err
	}

	log.Printf("Shares Outstanding data successfully fetched for: %s", symbol)
	return *sharesOutstandingsData, nil
}

func FetchEarnings(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.Earnings, error) {
	log.Printf("Fetching Earnings from AlphaVantage API for: %s", symbol)
	earningsData, err := apiClient.Earnings(symbol)
	if err != nil {
		log.Printf("Failed to fetch Earnings from AlphaVantage API for %s: %v", symbol, err)
		return models.Earnings{}, err
	}

	log.Printf("Earnings data successfully fetched for: %s", symbol)
	return *earningsData, nil
}
