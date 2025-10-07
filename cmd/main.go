package main

import (
	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/controllers"

	_ "github.com/remithomasn7/qualitinvest/docs" // Import indirect pour Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Investment API
// @version 1.0
// @description API for analyzing company financials based on Alpha Vantage data.
// @BasePath /api/v1
func main() {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	// Instanciation de l'API client
	apiClient := alpha_vantage.NewClient()

	// Groupes de routes pour les endpoints de l'API
	v1 := router.Group("/api/v1/companies/:symbol")
	{
		v1.GET("/overview", func(c *gin.Context) {
			controllers.GetCompanyOverview(c, apiClient)
		})
		v1.GET("/dividends", func(c *gin.Context) {
			controllers.GetDividends(c, apiClient)
		})
		v1.GET("/splits", func(c *gin.Context) {
			controllers.GetSplits(c, apiClient)
		})
		v1.GET("/income", func(c *gin.Context) {
			controllers.GetIncomeStatements(c, apiClient)
		})
		v1.GET("/balancesheet", func(c *gin.Context) {
			controllers.GetBalanceSheet(c, apiClient)
		})
		v1.GET("/cashflow", func(c *gin.Context) {
			controllers.GetCashFlowStatements(c, apiClient)
		})
		v1.GET("/shares_outstandings", func(c *gin.Context) {
			controllers.GetShareOutstandings(c, apiClient)
		})
		v1.GET("/earnings", func(c *gin.Context) {
			controllers.GetEarnings(c, apiClient)
		})
	}

	etf := router.Group("/api/v1/etf/:symbol")
	{
		etf.GET("/profile", func(c *gin.Context) {
			controllers.GetETFProfile(c, apiClient)
		})
	}
	// Route pour la documentation Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Démarrage du serveur
	router.Run(":8080")
}
