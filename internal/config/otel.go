package config

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
)

// InitOpenTelemetryLogger initialise le système de logging OpenTelemetry
func InitOpenTelemetryLogger(ctx context.Context) (func(), error) {
	// Créer l'exporteur (console pour le développement)
	exporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	// Créer la ressource avec les informations du service
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("qualitinvest-api"),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String(getOtelEnvOrDefault("ENV", "development")),
		),
	)
	if err != nil {
		return nil, err
	}

	// Créer le LoggerProvider
	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(res),
	)

	// Définir le LoggerProvider global
	global.SetLoggerProvider(loggerProvider)

	log.Println("✅ OpenTelemetry logging initialized")

	// Retourner une fonction de nettoyage
	return func() {
		if err := loggerProvider.Shutdown(ctx); err != nil {
			log.Printf("❌ Error shutting down OpenTelemetry logger: %v", err)
		}
		log.Println("🛑 OpenTelemetry logger shut down")
	}, nil
}

// getOtelEnvOrDefault retourne la valeur d'une variable d'environnement ou une valeur par défaut
func getOtelEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
