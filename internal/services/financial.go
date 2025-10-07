package services

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/models"
)

// Initialisation du cache avec une expiration de 24h et un nettoyage toutes les heures
var financialCache = cache.New(24*time.Hour, 1*time.Hour)

func FetchCompanyOverview(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.CompanyOverview, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("overview_" + symbol); found {
		log.Printf("Overview data successfuly found in cache for: %s", symbol)
		return cachedData.(models.CompanyOverview), nil
	}

	// If data are not in cache, call AlphaVantage API
	overviewData, err := apiClient.CompanyOverview(symbol)
	if err != nil {
		return models.CompanyOverview{}, err
	}

	// Store/Update data in cache
	financialCache.Set("overview_"+symbol, *overviewData, cache.DefaultExpiration)

	return *overviewData, nil
}

func FetchETFProfile(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.ETFProfile, error) {
	// Vérifie si les données sont en cache
	if cachedData, found := financialCache.Get("etfprofile_" + symbol); found {
		log.Printf("ETF Profile data found in cache for: %s", symbol)
		return cachedData.(models.ETFProfile), nil
	}

	// Si non, appelle l'API AlphaVantage
	etfProfileData, err := apiClient.ETFProfile(symbol)
	if err != nil {
		return models.ETFProfile{}, err
	}

	// Stocke les données en cache
	financialCache.Set("etfprofile_"+symbol, *etfProfileData, cache.DefaultExpiration)

	return *etfProfileData, nil
}

func FetchDividends(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.Dividends, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("dividends_" + symbol); found {
		log.Printf("Dividends data found in cache for: %s", symbol)
		return cachedData.(models.Dividends), nil
	}

	// If data are not in cache, call AlphaVantage API
	dividendsData, err := apiClient.Dividends(symbol)
	if err != nil {
		return models.Dividends{}, err
	}

	// Store/Update data in cache
	financialCache.Set("dividends_"+symbol, *dividendsData, cache.DefaultExpiration)

	return *dividendsData, nil
}

func FetchSplits(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.Splits, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("splits_" + symbol); found {
		log.Printf("Splits data found in cache for: %s", symbol)
		return cachedData.(models.Splits), nil
	}

	// If data are not in cache, call AlphaVantage API
	splitsData, err := apiClient.Splits(symbol)
	if err != nil {
		return models.Splits{}, err
	}

	// Store/Update data in cache
	financialCache.Set("splits_"+symbol, *splitsData, cache.DefaultExpiration)

	return *splitsData, nil
}

func FetchIncomeStatements(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.IncomeStatements, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("income_" + symbol); found {
		log.Printf("Income Statements data successfuly found in cache for: %s", symbol)
		return cachedData.(models.IncomeStatements), nil
	}

	// If data are not in cache, call AlphaVantage API
	incomeData, err := apiClient.IncomeStatements(symbol)
	if err != nil {
		return models.IncomeStatements{}, err
	}

	// Store/Update data in cache
	financialCache.Set("income_"+symbol, *incomeData, cache.DefaultExpiration)

	return *incomeData, nil
}

func FetchBalanceSheet(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.BalanceSheet, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("balancesheet_" + symbol); found {
		log.Printf("Balance Sheet data successfuly found in cache for: %s", symbol)
		return cachedData.(models.BalanceSheet), nil
	}

	// If data are not in cache, call AlphaVantage API
	balanceSheetData, err := apiClient.BalanceSheet(symbol)
	if err != nil {
		return models.BalanceSheet{}, err
	}

	// Store/Update data in cache
	financialCache.Set("balancesheet_"+symbol, *balanceSheetData, cache.DefaultExpiration)

	return *balanceSheetData, nil
}

func FetchCashFlowStatements(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.CashFlowStatements, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("cashflow_" + symbol); found {
		log.Printf("Cash Flow Statements successfuly found in cache for: %s", symbol)
		return cachedData.(models.CashFlowStatements), nil
	}

	// If data are not in cache, call AlphaVantage API
	cashFlowData, err := apiClient.CashFlowStatements(symbol)
	if err != nil {
		return models.CashFlowStatements{}, err
	}

	// Store/Update data in cache
	financialCache.Set("cashflow_"+symbol, *cashFlowData, cache.DefaultExpiration)

	return *cashFlowData, nil
}

func FetchSharesOutstandings(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.SharesOutstandings, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("shares_outstandings_" + symbol); found {
		log.Printf("Shares Outstanding data successfuly found in cache for: %s", symbol)
		return cachedData.(models.SharesOutstandings), nil
	}

	// If data are not in cache, call AlphaVantage API
	sharesOutstandingsData, err := apiClient.SharesOutstandings(symbol)
	if err != nil {
		return models.SharesOutstandings{}, err
	}

	// Store/Update data in cache
	financialCache.Set("shares_outstandings_"+symbol, *sharesOutstandingsData, cache.DefaultExpiration)

	return *sharesOutstandingsData, nil
}

func FetchEarnings(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.Earnings, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("earnings_" + symbol); found {
		log.Printf("Earnings data successfuly found in cache for: %s", symbol)
		return cachedData.(models.Earnings), nil
	}

	// If data are not in cache, call AlphaVantage API
	earningsData, err := apiClient.Earnings(symbol)
	if err != nil {
		return models.Earnings{}, err
	}

	// Store/Update data in cache
	financialCache.Set("earnings_"+symbol, *earningsData, cache.DefaultExpiration)

	return *earningsData, nil
}
