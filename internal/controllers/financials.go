package controllers

import (
	"log"
	"net/http"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/services"
	"github.com/remithomasn7/qualitinvest/pkg"

	"github.com/gin-gonic/gin"
)

// GetCompanyOverview - Return the Company's overview
// @Summary Get an Overview of a company
// @Description Get the overview  for a given company.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.OverviewResponse
// @Failure 500 {object} pkg.OverviewResponse
// @Router /companies/{symbol}/overview [get]
func GetCompanyOverview(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	overviewData, err := services.FetchCompanyOverview(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the OverviewData for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.OverviewResponse{
		Status: "success",
		Data:   overviewData,
	})
}

// GetETFProfile - Return the ETF profile
// @Summary Get ETF Profile
// @Description Get the profile for a given ETF.
// @Tags Financials
// @Param symbol path string true "ETF Symbol"
// @Success 200 {object} pkg.ETFProfileResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /etf/{symbol}/profile [get]
func GetETFProfile(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	etfProfileData, err := services.FetchETFProfile(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieving the ETF profile for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.ETFProfileResponse{
		Status: "success",
		Data:   etfProfileData,
	})
}

// GetDividends - Return the historical and future (declared) dividend distributions.
// @Summary Get the company's dividends
// @Description Get the dividends for a given company.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.DividendsResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /companies/{symbol}/dividends [get]
func GetDividends(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	dividendsData, err := services.FetchDividends(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieving the dividends for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.DividendsResponse{
		Status: "success",
		Data:   dividendsData,
	})
}

// GetSplits - Return the returns historical split events.
// @Summary Get the historical stock splits
// @Description Get the splits for a given company.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.SplitsResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /companies/{symbol}/splits [get]
func GetSplits(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	splitsData, err := services.FetchSplits(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieving the splits for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.SplitsResponse{
		Status: "success",
		Data:   splitsData,
	})
}

// GetIncomeStatements - Return the company's income statements
// @Summary Get the company's income statements
// @Description Get the company's income statements.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.OverviewResponse
// @Failure 500 {object} pkg.OverviewResponse
// @Router /companies/{symbol}/income [get]
func GetIncomeStatements(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	incomeStatementsData, err := services.FetchIncomeStatements(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the income statements for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.IncomeStatementsResponse{
		Status: "success",
		Data:   incomeStatementsData,
	})
}

// GetBalanceSheet - Return the company's balance sheet
// @Summary Get the company's balance sheet
// @Description Get the company's balance sheet.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.BalanceSheetResponse
// @Failure 500 {object} pkg.BalanceSheetResponse
// @Router /companies/{symbol}/balancesheet [get]
func GetBalanceSheet(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	balanceSheetData, err := services.FetchBalanceSheet(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the balance sheet for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.BalanceSheetResponse{
		Status: "success",
		Data:   balanceSheetData,
	})
}

// GetCashFlowStatements - Return the company's Cash Flow statements
// @Summary Get the company's cash flow statements
// @Description Get the company's cash flow statements.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.CashFlowStatementsResponse
// @Failure 500 {object} pkg.CashFlowStatementsResponse
// @Router /companies/{symbol}/cashflow [get]
func GetCashFlowStatements(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	cashFlowStatementsData, err := services.FetchCashFlowStatements(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the cash flow statements for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.CashFlowStatementsResponse{
		Status: "success",
		Data:   cashFlowStatementsData,
	})
}

// GetShareOutstandings - Return the company's shares outstandings
// @Summary Get the company's shares outstandings
// @Description Get the company's shares outstandings.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.SharesOutstandingsResponse
// @Failure 500 {object} pkg.SharesOutstandingsResponse
// @Router /companies/{symbol}/shares_outstandings [get]
func GetShareOutstandings(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	sharesOutstandingsData, err := services.FetchSharesOutstandings(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the shares outstandings for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.SharesOutstandingsResponse{
		Status: "success",
		Data:   sharesOutstandingsData,
	})
}

// GetEarnings - Return the company's earnings
// @Summary Get the company's earnings
// @Description Get the company's earnings.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.EarningsResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /companies/{symbol}/earnings [get]
func GetEarnings(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	earningsData, err := services.FetchEarnings(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the earnings for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.EarningsResponse{
		Status: "success",
		Data:   earningsData,
	})
}
