package controllers

import (
	"log"
	"net/http"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/services"
	"github.com/remithomasn7/qualitinvest/pkg"

	"github.com/gin-gonic/gin"
)

// GetCompanyOverview - Return the Company's overview (Internal endpoint)
func GetCompanyOverview(c *gin.Context, financialService *services.FinancialService) {
	symbol := c.Param("symbol")

	overviewData, err := financialService.FetchCompanyOverview(symbol)
	if err != nil {
		log.Printf("Error while retrieving the OverviewData for %s: %v", symbol, err)
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

// GetETFProfile - Return the ETF profile (Internal endpoint)
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
