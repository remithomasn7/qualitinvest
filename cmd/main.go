package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/config"
	"github.com/remithomasn7/qualitinvest/internal/controllers"
	"github.com/remithomasn7/qualitinvest/internal/services"
	"github.com/remithomasn7/qualitinvest/pkg"

	_ "github.com/remithomasn7/qualitinvest/docs" // Import indirect pour Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"

	otelog "go.opentelemetry.io/otel/log"
)

// otelWriter implémente io.Writer pour rediriger les logs Gin vers OpenTelemetry
type otelWriter struct {
	ctx    context.Context
	logger *pkg.Logger
}

func (w *otelWriter) Write(p []byte) (n int, err error) {
	message := strings.TrimSpace(string(p))

	if strings.Contains(message, "ERROR") || strings.Contains(message, "PANIC") {
		w.logger.Error(w.ctx, message, otelog.String("source", "gin"))
	} else {
		w.logger.Info(w.ctx, message, otelog.String("source", "gin"))
	}

	return len(p), nil
}

// @title My Investment API
// @version 1.0
// @description API for analyzing company financials based on Alpha Vantage data.
// @BasePath /
func main() {
	// Initialiser le contexte global
	ctx := context.Background()

	// Initialiser OpenTelemetry logging
	otelShutdown, err := config.InitOpenTelemetryLogger(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize OpenTelemetry logger: %v", err)
	}
	defer otelShutdown()

	// Créer le logger OpenTelemetry
	appLogger := pkg.NewLogger("qualitinvest-api")

	// Configurer Gin pour utiliser OpenTelemetry logging en production
	if os.Getenv("ENV") == "production" {
		gin.DefaultWriter = &otelWriter{ctx: ctx, logger: appLogger}
		gin.DefaultErrorWriter = &otelWriter{ctx: ctx, logger: appLogger}
	}

	// Configurer Gin pour utiliser OpenTelemetry logging en production
	if os.Getenv("ENV") == "production" {
		gin.DefaultWriter = &otelWriter{ctx: ctx, logger: appLogger}
		gin.DefaultErrorWriter = &otelWriter{ctx: ctx, logger: appLogger}
	}

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
	dataCollectionService := services.NewDataCollectionService(db, apiClient, appLogger)
	screeningService := services.NewScreeningService(db, appLogger)
	scheduler := services.NewDataScheduler(dataCollectionService)

	// Start the data collection scheduler
	scheduler.Start()

	// Initialize controllers
	screeningController := controllers.NewScreeningController(screeningService, dataCollectionService, appLogger)

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

	appLogger.Info(ctx, "🚀 Stock Screener API starting on :8080",
		otelog.String("port", ":8080"),
		otelog.String("version", "1.0.0"))
	appLogger.Info(ctx, "📊 Database connected and services initialized")
	appLogger.Info(ctx, "📈 Data collection scheduler started")
	appLogger.Info(ctx, "🎯 New screening endpoints available",
		otelog.String("screening_endpoint", "/api/v1/screening"))
	appLogger.Info(ctx, "📋 Company analysis available",
		otelog.String("analysis_endpoint", "/api/v1/analysis/{symbol}"))

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

	appLogger.Info(ctx, "🛑 Shutdown signal received, stopping gracefully...")

	// Arrêt propre du scheduler
	appLogger.Info(ctx, "⏹️ Stopping data collection scheduler...")
	scheduler.Stop()

	// Fermeture de la base de données
	appLogger.Info(ctx, "💾 Closing database connections...")
	db.Close()

	// Arrêt du serveur HTTP avec timeout
	appLogger.Info(ctx, "🌐 Shutting down HTTP server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		appLogger.Error(ctx, "❌ Server forced to shutdown",
			otelog.String("error", err.Error()))
		os.Exit(1)
	}

	appLogger.Info(ctx, "✅ Application stopped gracefully")
}
