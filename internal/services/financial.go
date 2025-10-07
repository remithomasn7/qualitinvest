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

func FetchROCE(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.ROCE, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("roce_" + symbol); found {
		log.Printf("Found ROCE data in cache : %v", cachedData)
		return cachedData.([]models.ROCE), nil
	}

	// If data are not in cache, call AlphaVantage API
	roceData, err := apiClient.GetROCE(symbol)
	log.Printf("GetROCE successfully called %v", roceData)
	if err != nil {
		return []models.ROCE{}, err
	}

	// Store/Update data in cache
	financialCache.Set("roce_"+symbol, roceData, cache.DefaultExpiration)

	return roceData, nil
}

// FetchMargins récupère les données de marge brute, opérationnelle, nette, et CAPEX/RO
func FetchMargins(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.Margins, error) {
	if cachedData, found := financialCache.Get("margins_" + symbol); found {
		return cachedData.([]models.Margins), nil
	}

	marginData, err := apiClient.GetMargins(symbol) // Appel à l'API pour les marges
	if err != nil {
		return []models.Margins{}, err
	}

	// Stocke les données en cache
	financialCache.Set("margins_"+symbol, marginData, cache.DefaultExpiration)

	return marginData, nil
}

// FetchGrowth récupère les données de croissance des revenus et un indice de prévisibilité
func FetchGrowth(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.Growth, error) {
	if cachedData, found := financialCache.Get("growth_" + symbol); found {
		return cachedData.([]models.Growth), nil
	}

	growthData, err := apiClient.GetGrowth(symbol) // Appel à l'API pour la croissance
	if err != nil {
		return []models.Growth{}, err
	}

	// Stocke les données en cache
	financialCache.Set("growth_"+symbol, growthData, cache.DefaultExpiration)

	return growthData, nil
}

// FetchFreeCashFlow récupère le Free Cash Flow par action sur les 10 dernières années
func FetchFreeCashFlow(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.FreeCashFlow, error) {
	if cachedData, found := financialCache.Get("fcf_" + symbol); found {
		return cachedData.([]models.FreeCashFlow), nil
	}

	fcfData, err := apiClient.GetFreeCashFlow(symbol) // Appel à l'API pour le Free Cash Flow
	if err != nil {
		return []models.FreeCashFlow{}, err
	}

	// Stocke les données en cache
	financialCache.Set("fcf_"+symbol, fcfData, cache.DefaultExpiration)

	return fcfData, nil
}

// FetchDebtEBITDA récupère la dette et l'EBITDA pour calculer le ratio dette/EBITDA
func FetchDebtEBITDA(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.DebtEBITDA, error) {
	if cachedData, found := financialCache.Get("debt_ebitda_" + symbol); found {
		return cachedData.([]models.DebtEBITDA), nil
	}

	debtData, err := apiClient.GetDebtEBITDA(symbol) // Appel à l'API pour la dette et l'EBITDA
	if err != nil {
		return []models.DebtEBITDA{}, err
	}

	// Stocke les données en cache
	financialCache.Set("debt_ebitda_"+symbol, debtData, cache.DefaultExpiration)

	return debtData, nil
}

// FetchRDExpenses récupère les dépenses en R&D sur les 10 dernières années
func FetchRDExpenses(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.RDExpenses, error) {
	if cachedData, found := financialCache.Get("rd_" + symbol); found {
		return cachedData.([]models.RDExpenses), nil
	}

	rdData, err := apiClient.GetRDExpenses(symbol) // Appel à l'API pour les dépenses en R&D
	if err != nil {
		return []models.RDExpenses{}, err
	}

	// Stocke les données en cache
	financialCache.Set("rd_"+symbol, rdData, cache.DefaultExpiration)

	return rdData, nil
}
