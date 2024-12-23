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

func toCompanyOverview(buf []byte) (*models.CompanyOverview, error) {
	var overview models.CompanyOverview
	err := json.Unmarshal(buf, &overview)
	if err != nil {
		return nil, err
	}
	return &overview, nil
}

func (c *AlphaVantageClient) CompanyOverview(symbol string) (*models.CompanyOverview, error) {
	params := map[string]string{
		"function": "OVERVIEW",
		"symbol":   symbol,
	}

	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}

	return toCompanyOverview(body)
}

func toIncomeStatements(buf []byte) (*models.IncomeStatements, error) {
	incomeStatements := &models.IncomeStatements{}
	if err := json.Unmarshal(buf, incomeStatements); err != nil {
		return nil, err
	}
	return incomeStatements, nil
}
func (c *AlphaVantageClient) IncomeStatements(symbol string) (*models.IncomeStatements, error) {
	params := map[string]string{
		"function": "INCOME_STATEMENT",
		"symbol":   symbol,
	}

	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}

	return toIncomeStatements(body)
}

func toBalanceSheet(buf []byte) (*models.BalanceSheet, error) {
	balancesheet := &models.BalanceSheet{}
	if err := json.Unmarshal(buf, balancesheet); err != nil {
		return nil, err
	}
	return balancesheet, nil
}
func (c *AlphaVantageClient) BalanceSheet(symbol string) (*models.BalanceSheet, error) {
	params := map[string]string{
		"function": "BALANCE_SHEET",
		"symbol":   symbol,
	}

	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}

	return toBalanceSheet(body)
}

func toCashFlowStatements(buf []byte) (*models.CashFlowStatements, error) {
	cashFlowStatements := &models.CashFlowStatements{}
	if err := json.Unmarshal(buf, cashFlowStatements); err != nil {
		return nil, err
	}
	return cashFlowStatements, nil
}
func (c *AlphaVantageClient) CashFlowStatements(symbol string) (*models.CashFlowStatements, error) {
	params := map[string]string{
		"function": "CASH_FLOW",
		"symbol":   symbol,
	}

	body, err := c.fetch(params)
	if err != nil {
		return nil, err
	}

	return toCashFlowStatements(body)
}

func (c *AlphaVantageClient) GetROCE(symbol string) ([]models.ROCE, error) {
	// params := map[string]string{
	// 	"function": "ROCE",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// return []models.ROCE{}, err
	// }

	// Mapper le résultat en modèle ROCE ici
	var roce []models.ROCE
	// roce0 := roce[0]
	// Assure-toi de parser le résultat selon la structure de ton modèle ROCE
	// Par exemple :
	// roce0 = result["ROCE"].(float64)

	return roce, nil
}

func (c *AlphaVantageClient) GetMargins(symbol string) ([]models.Margins, error) {
	// params := map[string]string{
	// 	"function": "GET_MARGIN",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// 	return []models.Margins{}, err
	// }

	var margins []models.Margins
	// Mapper le résultat en modèle Margins ici
	// margins[0].GrossMargin = result["ROCE"].(float64)

	return margins, nil
}

func (c *AlphaVantageClient) GetGrowth(symbol string) ([]models.Growth, error) {
	// params := map[string]string{
	// 	"function": "GET_GROWTH",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// 	return models.Growth{}, err
	// }

	var growth []models.Growth
	// Mapper le résultat en modèle Growth ici
	// growth.RevenueGrowth = result["GROWTH"].(float64)

	return growth, nil
}

func (c *AlphaVantageClient) GetFreeCashFlow(symbol string) ([]models.FreeCashFlow, error) {
	// params := map[string]string{
	// 	"function": "GET_FCF",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// 	return models.FreeCashFlow{}, err
	// }

	var fcf []models.FreeCashFlow
	// Mapper le résultat en modèle Growth ici
	// fcf.FCFPerShare = result["FCF"].(float64)

	return fcf, nil
}

func (c *AlphaVantageClient) GetShareDilution(symbol string) ([]models.ShareDilution, error) {
	// params := map[string]string{
	// 	"function": "GET_SHARE_DILUTION",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// 	return models.ShareDilution{}, err
	// }

	var share []models.ShareDilution
	// Mapper le résultat en modèle Growth ici
	// share.SharesOutstanding = result["SHARE"].(float64)

	return share, nil
}

func (c *AlphaVantageClient) GetDebtEBITDA(symbol string) ([]models.DebtEBITDA, error) {
	// params := map[string]string{
	// 	"function": "GET_EBITDA_DEBT",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// 	return models.DebtEBITDA{}, err
	// }

	var debtEBITDA []models.DebtEBITDA
	// Mapper le résultat en modèle Growth ici
	// debtEBITDA.DebtToEBITDA = result["DEBT_EBITDA"].(float64)

	return debtEBITDA, nil
}

func (c *AlphaVantageClient) GetRDExpenses(symbol string) ([]models.RDExpenses, error) {
	// params := map[string]string{
	// 	"function": "GET_RD_EXPENSES",
	// 	"symbol":   symbol,
	// }
	// result, err := c.fetch("", params)
	// if err != nil {
	// 	return models.RDExpenses{}, err
	// }

	var rdExpenses []models.RDExpenses
	// Mapper le résultat en modèle Growth ici
	// rdExpenses.Expenses = result["RD_EXPENSES"].(float64)

	return rdExpenses, nil
}
