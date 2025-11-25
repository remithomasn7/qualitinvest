package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/remithomasn7/qualitinvest/internal/services"
	"github.com/remithomasn7/qualitinvest/pkg"
)

// ScreeningController gère les endpoints de screening
type ScreeningController struct {
	screeningService     *services.ScreeningService
	dataCollectionService *services.DataCollectionService
}

// NewScreeningController crée un nouveau contrôleur de screening
func NewScreeningController(screeningSvc *services.ScreeningService, dataSvc *services.DataCollectionService) *ScreeningController {
	return &ScreeningController{
		screeningService:     screeningSvc,
		dataCollectionService: dataSvc,
	}
}

// @Summary Screen companies based on fundamental criteria
// @Description Filter companies based on valuation, profitability, and financial health criteria
// @Tags Screening
// @Accept json
// @Produce json
// @Param criteria body services.ScreeningCriteria true "Screening criteria"
// @Success 200 {object} pkg.ScreeningResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /api/v1/screening [post]
func (sc *ScreeningController) ScreenCompanies(c *gin.Context) {
	var criteria services.ScreeningCriteria
	if err := c.ShouldBindJSON(&criteria); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Status:  "error",
			Message: "Invalid screening criteria: " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	results, err := sc.screeningService.ScreenCompanies(criteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: "Screening failed: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, pkg.ScreeningResponse{
		Status: "success",
		Data:   results,
		Count:  len(results),
	})
}

// @Summary Get detailed company analysis
// @Description Get comprehensive fundamental analysis for a specific company
// @Tags Analysis
// @Accept json
// @Produce json
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.CompanyAnalysisResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /api/v1/analysis/{symbol} [get]
func (sc *ScreeningController) GetCompanyAnalysis(c *gin.Context) {
	symbol := c.Param("symbol")

	// Vérifier si les données existent
	stale, err := sc.dataCollectionService.IsDataStale(symbol, 0)
	if err != nil {
		// Erreur technique lors de la vérification
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: "Failed to check data availability: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	if stale {
		// Les données n'existent pas ou sont trop vieilles
		c.JSON(http.StatusNotFound, pkg.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("No data available for symbol %s. Please collect data first using POST /api/v1/collect/%s", symbol, symbol),
			Code:    http.StatusNotFound,
		})
		return
	}

	// Les données existent, récupérer l'analyse
	analysis, err := sc.screeningService.GetCompanyAnalysis(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve company analysis: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, pkg.CompanyAnalysisResponse{
		Status: "success",
		Data:   analysis,
	})
}

// @Summary Get sector comparison for a company
// @Description Compare company metrics with sector averages and peers
// @Tags Analysis
// @Accept json
// @Produce json
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.SectorComparisonResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /api/v1/analysis/{symbol}/sector [get]
func (sc *ScreeningController) GetSectorComparison(c *gin.Context) {
	symbol := c.Param("symbol")

	comparison, err := sc.screeningService.GetSectorComparison(symbol)
	if err != nil {
		c.JSON(http.StatusNotFound, pkg.ErrorResponse{
			Status:  "error",
			Message: "Sector comparison not available: " + err.Error(),
			Code:    http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, pkg.SectorComparisonResponse{
		Status: "success",
		Data:   comparison,
	})
}

// @Summary Get valuation analysis for a company
// @Description Get detailed valuation analysis with fair value estimates
// @Tags Analysis
// @Accept json
// @Produce json
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.ValuationAnalysisResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /api/v1/analysis/{symbol}/valuation [get]
func (sc *ScreeningController) GetValuationAnalysis(c *gin.Context) {
	// TODO: Implémenter la logique d'analyse de valorisation
	// Pour l'instant, retourner une réponse temporaire
	c.JSON(http.StatusNotImplemented, pkg.ErrorResponse{
		Status:  "error",
		Message: "Valuation analysis not yet implemented",
		Code:    http.StatusNotImplemented,
	})
}

// @Summary Trigger data collection for a company
// @Description Manually trigger data collection from Alpha Vantage for a specific company
// @Tags Data Collection
// @Accept json
// @Produce json
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /api/v1/collect/{symbol} [post]
func (sc *ScreeningController) CollectCompanyData(c *gin.Context) {
	symbol := c.Param("symbol")

	// Déclencher la collecte et vérifier le résultat
	err := sc.dataCollectionService.CollectCompanyData(symbol)
	if err != nil {
		// Déterminer le type d'erreur pour le code HTTP approprié
		errorMessage := err.Error()
		statusCode := http.StatusInternalServerError

		// Si c'est un symbole invalide, retourner 400 Bad Request
		if strings.Contains(errorMessage, "invalid symbol") || strings.Contains(errorMessage, "company not found") {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, pkg.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Data collection failed for %s: %s", symbol, errorMessage),
			Code:    statusCode,
		})
		return
	}

	c.JSON(http.StatusOK, pkg.SuccessResponse{
		Status:  "success",
		Message: "Data collection completed for " + symbol,
	})
}

// @Summary Get screening templates
// @Description Get predefined screening templates for common investment strategies
// @Tags Screening
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ScreeningTemplatesResponse
// @Router /api/v1/screening/templates [get]
func (sc *ScreeningController) GetScreeningTemplates(c *gin.Context) {
	templates := []pkg.ScreeningTemplate{
		{
			Name:        "Value Investing",
			Description: "Classic value investing criteria",
			Criteria: services.ScreeningCriteria{
				MaxPERatio:      floatPtr(15),
				MaxPBRatio:      floatPtr(1.5),
				MinROE:          floatPtr(0.10),
				MaxDebtToEquity: floatPtr(1.0),
				Limit:           100,
			},
		},
		{
			Name:        "Growth Investing",
			Description: "High growth companies",
			Criteria: services.ScreeningCriteria{
				MinRevenueGrowth: floatPtr(0.15),
				MinEPSGrowth:     floatPtr(0.15),
				MinROE:           floatPtr(0.15),
				Limit:            100,
			},
		},
		{
			Name:        "Quality Investing",
			Description: "High quality, stable companies",
			Criteria: services.ScreeningCriteria{
				MinProfitMargin:  floatPtr(0.10),
				MinROE:           floatPtr(0.15),
				MaxDebtToEquity:   floatPtr(0.5),
				MinCurrentRatio:   floatPtr(1.5),
				Limit:             100,
			},
		},
	}

	c.JSON(http.StatusOK, pkg.ScreeningTemplatesResponse{
		Status:    "success",
		Templates: templates,
	})
}

// Helper functions
func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}
