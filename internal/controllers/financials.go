package controllers

import (
	"log"
	"net/http"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/services"
	"github.com/remithomasn7/qualitinvest/pkg"

	"github.com/gin-gonic/gin"
)

// CompanyOverview - Return the Company's overview
// @Summary Get an Overview of a company
// @Description Get the overview  for a given company.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.OverviewResponse
// @Failure 500 {object} pkg.OverviewResponse
// @Router /companies/{symbol}/overview [get]
func CompanyOverview(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
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

// IncomeStatements - Return the company's income statements
// @Summary Get the company's income statements
// @Description Get the company's income statements.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.OverviewResponse
// @Failure 500 {object} pkg.OverviewResponse
// @Router /companies/{symbol}/incomestatements [get]
func IncomeStatements(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
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

// BalanceSheet - Return the company's balance sheet
// @Summary Get the company's balance sheet
// @Description Get the company's balance sheet.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.BalanceSheetResponse
// @Failure 500 {object} pkg.BalanceSheetResponse
// @Router /companies/{symbol}/balancesheet [get]
func BalanceSheet(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
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

// GetROCE - Return the ROCE (Return on Capital Employed) over the last 10 years
// @Summary Get ROCE for a company
// @Description Get the ROCE ratio for the last 10 years for a given company. ROCE is a measure of how efficiently a company is using its capital to generate profits.
// @Tags Financials
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.ROCEResponse
// @Failure 500 {object} pkg.ROCEResponse
// @Router /companies/{symbol}/roce [get]
func GetROCE(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")

	roceData, err := services.FetchROCE(apiClient, symbol)
	if err != nil {
		log.Printf("Error while retrieve the ROCE for %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.ROCEResponse{
		Status: "success",
		Data:   roceData,
	})
}

// GetMargins retourne les données de marges pour une entreprise spécifique
// @Summary Get Margins Data
// @Description Récupère les marges brute, opérationnelle, nette, et le ratio CAPEX/Résultat opérationnel pour une entreprise donnée sur les 10 dernières années.
// @Tags Financials
// @Param symbol path string true "Symbol de l'entreprise"
// @Produce json
// @Success 200 {object} pkg.MarginResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /financials/{symbol}/margins [get]
func GetMargins(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")
	marginsData, err := services.FetchMargins(apiClient, symbol)
	if err != nil {
		log.Printf("Erreur lors de la récupération des marges pour %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.MarginResponse{
		Status: "success",
		Data:   marginsData,
	})
}

// GetGrowth retourne les données de croissance pour une entreprise spécifique
// @Summary Get Growth Data
// @Description Récupère le taux de croissance du chiffre d'affaires et un indice de prévisibilité du chiffre d'affaires pour une entreprise sur les 10 dernières années.
// @Tags Financials
// @Param symbol path string true "Symbol de l'entreprise"
// @Produce json
// @Success 200 {object} pkg.GrowthResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /financials/{symbol}/growth [get]
func GetGrowth(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")
	growthData, err := services.FetchGrowth(apiClient, symbol)
	if err != nil {
		log.Printf("Erreur lors de la récupération des données de croissance pour %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.GrowthResponse{
		Status: "success",
		Data:   growthData,
	})
}

// GetFreeCashFlow retourne le Free Cash Flow par action pour une entreprise spécifique
// @Summary Get Free Cash Flow Data
// @Description Récupère le Free Cash Flow par action sur les 10 dernières années pour une entreprise donnée.
// @Tags Financials
// @Param symbol path string true "Symbol de l'entreprise"
// @Produce json
// @Success 200 {object} pkg.FreeCashFlowResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /financials/{symbol}/free-cash-flow [get]
func GetFreeCashFlow(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")
	fcfData, err := services.FetchFreeCashFlow(apiClient, symbol)
	if err != nil {
		log.Printf("Erreur lors de la récupération du Free Cash Flow pour %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.FreeCashFlowResponse{
		Status: "success",
		Data:   fcfData,
	})
}

// GetShareDilution retourne le nombre d'actions en circulation pour une entreprise spécifique
// @Summary Get Share Dilution Data
// @Description Récupère le nombre d'actions en circulation pour une entreprise donnée sur les 10 dernières années afin d'analyser la dilution des actionnaires.
// @Tags Financials
// @Param symbol path string true "Symbol de l'entreprise"
// @Produce json
// @Success 200 {object} pkg.ShareDilutionResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /financials/{symbol}/share-dilution [get]
func GetShareDilution(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")
	sharesData, err := services.FetchShareDilution(apiClient, symbol)
	if err != nil {
		log.Printf("Erreur lors de la récupération des actions en circulation pour %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.ShareDilutionResponse{
		Status: "success",
		Data:   sharesData,
	})
}

// GetDebtEBITDA retourne les données de dette et d'EBITDA pour une entreprise spécifique
// @Summary Get Debt to EBITDA Data
// @Description Récupère la dette et l'EBITDA d'une entreprise sur les 10 dernières années pour calculer le ratio dette/EBITDA.
// @Tags Financials
// @Param symbol path string true "Symbol de l'entreprise"
// @Produce json
// @Success 200 {object} pkg.DebtEBITDAResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /financials/{symbol}/debt-ebitda [get]
func GetDebtEBITDA(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")
	debtData, err := services.FetchDebtEBITDA(apiClient, symbol)
	if err != nil {
		log.Printf("Erreur lors de la récupération de la dette pour %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.DebtEBITDAResponse{
		Status: "success",
		Data:   debtData,
	})
}

// GetRDExpenses retourne les dépenses en R&D pour une entreprise spécifique
// @Summary Get R&D Expenses Data
// @Description Récupère les dépenses en recherche et développement pour une entreprise donnée sur les 10 dernières années.
// @Tags Financials
// @Param symbol path string true "Symbol de l'entreprise"
// @Produce json
// @Success 200 {object} pkg.RDExpensesResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /financials/{symbol}/rd-expenses [get]
func GetRDExpenses(c *gin.Context, apiClient *alpha_vantage.AlphaVantageClient) {
	symbol := c.Param("symbol")
	rdData, err := services.FetchRDExpenses(apiClient, symbol)
	if err != nil {
		log.Printf("Erreur lors de la récupération des dépenses en R&D pour %s: %v", symbol, err)
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, pkg.RDExpensesResponse{
		Status: "success",
		Data:   rdData,
	})
}
