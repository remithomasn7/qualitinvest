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

	// Instanciation de l'API client
	apiClient := alpha_vantage.NewClient()

	// Groupes de routes pour les endpoints de l'API
	v1 := router.Group("/api/v1/companies/:symbol")
	{
		v1.GET("/overview", func(c *gin.Context) {
			controllers.CompanyOverview(c, apiClient)
		})
		v1.GET("/incomestatements", func(c *gin.Context) {
			controllers.IncomeStatements(c, apiClient)
		})
		v1.GET("/balancesheet", func(c *gin.Context) {
			controllers.BalanceSheet(c, apiClient)
		})
		v1.GET("/roce", func(c *gin.Context) {
			controllers.GetROCE(c, apiClient)
		})
		v1.GET("/margins", func(c *gin.Context) {
			controllers.GetMargins(c, apiClient)
		})
		v1.GET("/growth", func(c *gin.Context) {
			controllers.GetGrowth(c, apiClient)
		})
		v1.GET("/cashflow", func(c *gin.Context) {
			controllers.GetFreeCashFlow(c, apiClient)
		})
		v1.GET("/shares", func(c *gin.Context) {
			controllers.GetShareDilution(c, apiClient)
		})
		v1.GET("/debt", func(c *gin.Context) {
			controllers.GetDebtEBITDA(c, apiClient)
		})
		v1.GET("/rd", func(c *gin.Context) {
			controllers.GetRDExpenses(c, apiClient)
		})
	}

	// Route pour la documentation Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Démarrage du serveur
	router.Run(":8080")
}
