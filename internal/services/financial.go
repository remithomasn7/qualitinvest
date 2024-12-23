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
		log.Printf("Found overview data in cache : %v", cachedData)
		return cachedData.(models.CompanyOverview), nil
	}

	// If data are not in cache, call AlphaVantage API
	overviewData, err := apiClient.CompanyOverview(symbol)
	log.Printf("FetchCompanyOverview successfully called %v", overviewData)
	if err != nil {
		return models.CompanyOverview{}, err
	}

	// Store/Update data in cache
	financialCache.Set("overview_"+symbol, *overviewData, cache.DefaultExpiration)

	return *overviewData, nil
}

func FetchIncomeStatements(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.IncomeStatements, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("incomestatements_" + symbol); found {
		log.Printf("Found income statements data in cache : %v", cachedData)
		return cachedData.(models.IncomeStatements), nil
	}

	// If data are not in cache, call AlphaVantage API
	incomestatementsData, err := apiClient.IncomeStatements(symbol)
	log.Printf("FetchIncomeStatements successfully called %v", incomestatementsData)
	if err != nil {
		return models.IncomeStatements{}, err
	}

	// Store/Update data in cache
	financialCache.Set("incomestatements_"+symbol, *incomestatementsData, cache.DefaultExpiration)

	return *incomestatementsData, nil
}

func FetchBalanceSheet(apiClient *alpha_vantage.AlphaVantageClient, symbol string) (models.BalanceSheet, error) {
	// Check if data are in cache
	if cachedData, found := financialCache.Get("balancesheet_" + symbol); found {
		log.Printf("Found balance sheet data in cache : %v", cachedData)
		return cachedData.(models.BalanceSheet), nil
	}

	// If data are not in cache, call AlphaVantage API
	balanceSheetData, err := apiClient.BalanceSheet(symbol)
	log.Printf("FetchBalanceSheet successfully called %v", balanceSheetData)
	if err != nil {
		return models.BalanceSheet{}, err
	}

	// Store/Update data in cache
	financialCache.Set("balancesheet_"+symbol, *balanceSheetData, cache.DefaultExpiration)

	return *balanceSheetData, nil
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

// FetchShareDilution récupère le nombre d'actions en circulation sur les 10 dernières années
func FetchShareDilution(apiClient *alpha_vantage.AlphaVantageClient, symbol string) ([]models.ShareDilution, error) {
	if cachedData, found := financialCache.Get("shares_" + symbol); found {
		return cachedData.([]models.ShareDilution), nil
	}

	sharesData, err := apiClient.GetShareDilution(symbol) // Appel à l'API pour les actions en circulation
	if err != nil {
		return []models.ShareDilution{}, err
	}

	// Stocke les données en cache
	financialCache.Set("shares_"+symbol, sharesData, cache.DefaultExpiration)

	return sharesData, nil
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
