package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/config"
	"github.com/remithomasn7/qualitinvest/internal/controllers"
	"github.com/remithomasn7/qualitinvest/internal/services"

	_ "github.com/remithomasn7/qualitinvest/docs" // Import indirect pour Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Investment API
// @version 1.0
// @description API for analyzing company financials based on Alpha Vantage data.
// @BasePath /
func main() {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	// Initialize database connection
	dbConfig := config.NewDatabaseConfig()
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Alpha Vantage API client
	apiClient := alpha_vantage.NewClient()

	// Initialize services
	dataCollectionService := services.NewDataCollectionService(db, apiClient)
	screeningService := services.NewScreeningService(db)
	scheduler := services.NewDataScheduler(dataCollectionService)

	// Start the data collection scheduler
	scheduler.Start()

	// Initialize controllers
	screeningController := controllers.NewScreeningController(screeningService, dataCollectionService)

	// ============================
	// NOUVEAUX ENDPOINTS DE SCREENING ET ANALYSE
	// ============================

	// Screening endpoints
	router.POST("/api/v1/screening", screeningController.ScreenCompanies)
	router.GET("/api/v1/screening/templates", screeningController.GetScreeningTemplates)

	// Company analysis endpoints
	analysis := router.Group("/api/v1/analysis")
	{
		analysis.GET("/:symbol", screeningController.GetCompanyAnalysis)
		analysis.GET("/:symbol/sector", screeningController.GetSectorComparison)
		analysis.GET("/:symbol/valuation", screeningController.GetValuationAnalysis)
	}

	// Data collection management
	router.POST("/api/v1/collect/:symbol", screeningController.CollectCompanyData)

	// ============================
	// ENDPOINTS INTERNES (non exposés publiquement)
	// ============================
	// Les anciens endpoints sont maintenant utilisés uniquement en interne
	// par les services d'analyse et de screening

	// Route pour la documentation Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Stock Screener API starting on :8080")
	log.Println("📊 Database connected and services initialized")
	log.Println("📈 Data collection scheduler started")
	log.Println("🎯 New screening endpoints available at /api/v1/screening")
	log.Println("📋 Company analysis available at /api/v1/analysis/{symbol}")

	// Configuration du serveur avec gestion d'arrêt propre
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Démarrage du serveur dans une goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Attente des signaux d'arrêt (SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutdown signal received, stopping gracefully...")

	// Arrêt propre du scheduler
	log.Println("⏹️ Stopping data collection scheduler...")
	scheduler.Stop()

	// Fermeture de la base de données
	log.Println("💾 Closing database connections...")
	db.Close()

	// Arrêt du serveur HTTP avec timeout
	log.Println("🌐 Shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("❌ Server forced to shutdown: %v", err)
		os.Exit(1)
	}

	log.Println("✅ Application stopped gracefully")
}
