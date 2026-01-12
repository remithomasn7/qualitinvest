package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remithomasn7/qualitinvest/internal/services"
	"github.com/remithomasn7/qualitinvest/pkg"
	"go.opentelemetry.io/otel/log"
)

// ScreeningController gère les endpoints de screening
type ScreeningController struct {
	logger                *pkg.Logger
	screeningService      *services.ScreeningService
	dataCollectionService *services.DataCollectionService
}

// NewScreeningController crée un nouveau contrôleur de screening
func NewScreeningController(screeningSvc *services.ScreeningService, dataSvc *services.DataCollectionService, logger *pkg.Logger) *ScreeningController {
	return &ScreeningController{
		logger:                logger,
		screeningService:      screeningSvc,
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
// @Description Get comprehensive fundamental analysis for a specific company. Data collection is automatic if not available.
// @Tags Analysis
// @Accept json
// @Produce json
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.CompanyAnalysisResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /api/v1/analysis/{symbol} [get]
func (sc *ScreeningController) GetCompanyAnalysis(c *gin.Context) {
	ctx := context.Background()
	startTime := time.Now()
	symbol := c.Param("symbol")

	sc.logger.Info(ctx, "📥 Requête d'analyse d'entreprise reçue",
		log.String("symbol", symbol),
		log.String("endpoint", "/api/v1/analysis/"+symbol),
		log.String("method", "GET"))

	// Vérifier si les données existent
	sc.logger.Debug(ctx, "🔍 Vérification de la disponibilité des données",
		log.String("symbol", symbol))

	stale, err := sc.dataCollectionService.IsDataStale(symbol, 0)
	if err != nil {
		// Erreur technique lors de la vérification
		sc.logger.Error(ctx, "❌ Erreur lors de la vérification des données",
			log.String("symbol", symbol),
			log.String("error", err.Error()),
			log.Float64("duration_ms", float64(time.Since(startTime).Milliseconds())))

		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: "Failed to check data availability: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	if stale {
		// Les données n'existent pas - collecte automatique
		sc.logger.Info(ctx, "🔄 Données manquantes - démarrage collecte automatique",
			log.String("symbol", symbol))

		collectStart := time.Now()
		err := sc.dataCollectionService.CollectCompanyData(symbol)
		collectDuration := time.Since(collectStart)

		if err != nil {
			sc.logger.Error(ctx, "❌ Échec de la collecte automatique",
				log.String("symbol", symbol),
				log.String("error", err.Error()),
				log.Float64("collect_duration_ms", float64(collectDuration.Milliseconds())),
				log.Float64("total_duration_ms", float64(time.Since(startTime).Milliseconds())))

			// Déterminer le code d'erreur approprié
			errorMessage := err.Error()
			statusCode := http.StatusInternalServerError
			if strings.Contains(errorMessage, "invalid symbol") || strings.Contains(errorMessage, "company not found") {
				statusCode = http.StatusBadRequest
			}

			c.JSON(statusCode, pkg.ErrorResponse{
				Status:  "error",
				Message: fmt.Sprintf("Failed to collect data for %s: %s", symbol, errorMessage),
				Code:    statusCode,
			})
			return
		}

		sc.logger.Info(ctx, "✅ Collecte automatique réussie - données prêtes pour analyse",
			log.String("symbol", symbol),
			log.Float64("collect_duration_ms", float64(collectDuration.Milliseconds())),
			log.Float64("total_duration_ms", float64(time.Since(startTime).Milliseconds())))
	}

	sc.logger.Debug(ctx, "✅ Données disponibles, récupération de l'analyse",
		log.String("symbol", symbol))

	// Les données existent, récupérer l'analyse
	analysis, err := sc.screeningService.GetCompanyAnalysis(symbol)
	if err != nil {
		sc.logger.Error(ctx, "❌ Erreur lors de la récupération de l'analyse",
			log.String("symbol", symbol),
			log.String("error", err.Error()),
			log.Float64("duration_ms", float64(time.Since(startTime).Milliseconds())))

		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve company analysis: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	duration := time.Since(startTime)
	sc.logger.Info(ctx, "✅ Analyse d'entreprise terminée avec succès",
		log.String("symbol", symbol),
		log.Float64("duration_ms", float64(duration.Milliseconds())),
		log.String("company_name", analysis.Overview.Name),
		log.String("sector", analysis.Overview.Sector))

	c.JSON(http.StatusOK, pkg.CompanyAnalysisResponse{
		Status: "success",
		Data:   analysis,
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

// @Summary Trigger data collection for a company (Admin/Testing)
// @Description Manually trigger data collection from Alpha Vantage. Note: Data collection is automatic in analysis endpoints.
// @Tags Data Collection
// @Accept json
// @Produce json
// @Param symbol path string true "Company Symbol"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400 {object} pkg.ErrorResponse
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
			Name:        "Value Investing (Adapted)",
			Description: "Classic value investing - adapted for AAPL/MSFT data",
			Criteria: services.ScreeningCriteria{
				MaxPERatio: floatPtr(40), // AAPL=37.04, MSFT=33.66
				MaxPBRatio: floatPtr(60), // AAPL=54.41
				Limit:      100,
			},
		},
		{
			Name:        "Quality Investing (Adapted)",
			Description: "High quality companies - adapted for current data",
			Criteria: services.ScreeningCriteria{
				MinProfitMargin: floatPtr(0.25), // MSFT=0.357 (35.7%)
				Limit:           100,
			},
		},
		{
			Name:        "All Companies",
			Description: "Show all companies in database",
			Criteria: services.ScreeningCriteria{
				Limit: 100,
			},
		},
		{
			Name:        "Technology Sector",
			Description: "All technology companies",
			Criteria: services.ScreeningCriteria{
				Sectors: []string{"TECHNOLOGY"},
				Limit:   100,
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
