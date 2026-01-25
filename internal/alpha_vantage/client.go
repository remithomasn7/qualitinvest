package alpha_vantage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

const alphaVantageBaseURL = "https://www.alphavantage.co/query"

type AlphaVantageClient struct {
	apiKey string
}

func NewClient() *AlphaVantageClient {
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("ALPHA_VANTAGE_KEY") // fallback pour compatibilité
	}
	return &AlphaVantageClient{apiKey: apiKey}
}

func (c *AlphaVantageClient) fetch(params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", alphaVantageBaseURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", c.apiKey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: expected %d, got %d",
			http.StatusOK, resp.StatusCode)
	}

	return body, nil
}

func (c *AlphaVantageClient) CompanyOverview(symbol string) (*models.CompanyOverview, error) {
	overview := &models.CompanyOverview{}

	params := map[string]string{
		"function": "OVERVIEW",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}

	// Vérifier si la réponse contient un message d'erreur/demo
	if strings.Contains(string(body), "Information") && strings.Contains(string(body), "demo") {
		return nil, fmt.Errorf("demo API key used - please use a valid API key")
	}

	err = json.Unmarshal(body, overview)
	if err != nil {
		return nil, err
	}
	return overview, nil
}

func (c *AlphaVantageClient) ETFProfile(symbol string) (*models.ETFProfile, error) {
	ETFProfile := &models.ETFProfile{}

	params := map[string]string{
		"function": "ETF_PROFILE",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, ETFProfile)
	if err != nil {
		return nil, err
	}
	return ETFProfile, nil
}

func (c *AlphaVantageClient) Dividends(symbol string) (*models.Dividends, error) {
	dividends := &models.Dividends{}

	params := map[string]string{
		"function": "DIVIDENDS",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, dividends)
	if err != nil {
		return nil, err
	}
	return dividends, nil
}

func (c *AlphaVantageClient) SharesOutstandings(symbol string) (*models.SharesOutstandings, error) {
	SharesOutstandings := &models.SharesOutstandings{}

	params := map[string]string{
		"function": "SHARES_OUTSTANDING",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, SharesOutstandings)
	if err != nil {
		return nil, err
	}
	return SharesOutstandings, nil
}

func (c *AlphaVantageClient) Earnings(symbol string) (*models.Earnings, error) {
	earnings := &models.Earnings{}

	params := map[string]string{
		"function": "EARNINGS",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, earnings)
	if err != nil {
		return nil, err
	}
	return earnings, nil
}

func (c *AlphaVantageClient) Splits(symbol string) (*models.Splits, error) {
	splits := &models.Splits{}

	params := map[string]string{
		"function": "SPLITS",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, splits)
	if err != nil {
		return nil, err
	}
	return splits, nil
}

func (c *AlphaVantageClient) IncomeStatements(symbol string) (*models.IncomeStatements, error) {
	incomeStatements := &models.IncomeStatements{}

	params := map[string]string{
		"function": "INCOME_STATEMENT",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, incomeStatements); err != nil {
		return nil, err
	}
	return incomeStatements, nil
}

func (c *AlphaVantageClient) BalanceSheet(symbol string) (*models.BalanceSheet, error) {
	balancesheet := &models.BalanceSheet{}

	params := map[string]string{
		"function": "BALANCE_SHEET",
		"symbol":   symbol,
	}
	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, balancesheet); err != nil {
		return nil, err
	}
	return balancesheet, nil
}

func (c *AlphaVantageClient) CashFlowStatements(symbol string) (*models.CashFlowStatements, error) {
	cashFlowStatements := &models.CashFlowStatements{}

	params := map[string]string{
		"function": "CASH_FLOW",
		"symbol":   symbol,
	}

	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, cashFlowStatements); err != nil {
		return nil, err
	}
	return cashFlowStatements, nil
}

// SymbolSearch recherche des symboles d'entreprises via Alpha Vantage SYMBOL_SEARCH
func (c *AlphaVantageClient) SymbolSearch(keywords string) (*models.SymbolSearchResponse, error) {
	searchResponse := &models.SymbolSearchResponse{}

	params := map[string]string{
		"function": "SYMBOL_SEARCH",
		"keywords": keywords,
	}

	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}

	// Alpha Vantage retourne soit "bestMatches" soit "Error Message" ou "Note"
	var rawResponse map[string]interface{}
	if err = json.Unmarshal(body, &rawResponse); err != nil {
		return nil, err
	}

	// Vérifier si c'est une erreur
	if note, ok := rawResponse["Note"].(string); ok {
		return nil, fmt.Errorf("API note: %s", note)
	}
	if errorMsg, ok := rawResponse["Error Message"].(string); ok {
		return nil, fmt.Errorf("API error: %s", errorMsg)
	}

	// Parser les bestMatches
	if bestMatches, ok := rawResponse["bestMatches"].([]interface{}); ok {
		for _, match := range bestMatches {
			if matchMap, ok := match.(map[string]interface{}); ok {
				result := models.SymbolSearchResult{}
				if symbol, ok := matchMap["1. symbol"].(string); ok {
					result.Symbol = symbol
				}
				if name, ok := matchMap["2. name"].(string); ok {
					result.Name = name
				}
				if matchScore, ok := matchMap["3. type"].(string); ok {
					result.Type = matchScore
				}
				if region, ok := matchMap["4. region"].(string); ok {
					result.Region = region
				}
				if marketOpen, ok := matchMap["5. marketOpen"].(string); ok {
					result.MarketOpen = marketOpen
				}
				if marketClose, ok := matchMap["6. marketClose"].(string); ok {
					result.MarketClose = marketClose
				}
				if timezone, ok := matchMap["7. timezone"].(string); ok {
					result.Timezone = timezone
				}
				if currency, ok := matchMap["8. currency"].(string); ok {
					result.Currency = currency
				}
				if matchScore, ok := matchMap["9. matchScore"].(string); ok {
					// Convertir le score en float64
					if score, err := strconv.ParseFloat(matchScore, 64); err == nil {
						result.MatchScore = score
					}
				}
				searchResponse.BestMatches = append(searchResponse.BestMatches, result)
			}
		}
	}

	return searchResponse, nil
}
