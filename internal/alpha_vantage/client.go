package alpha_vantage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

const alphaVantageBaseURL = "https://www.alphavantage.co/query"

type AlphaVantageClient struct {
	apiKey string
}

func NewClient() *AlphaVantageClient {
	return &AlphaVantageClient{apiKey: os.Getenv("ALPHA_VANTAGE_KEY")}
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
