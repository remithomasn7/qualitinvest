package pkg

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
)

// Logger fournit une interface simplifiée pour le logging OpenTelemetry
type Logger struct {
	logger log.Logger
}

// NewLogger crée un nouveau logger avec le nom spécifié
func NewLogger(name string) *Logger {
	return &Logger{
		logger: global.Logger(name),
	}
}

// Info enregistre un message d'information
func (l *Logger) Info(ctx context.Context, message string, attrs ...log.KeyValue) {
	record := log.Record{}
	record.SetBody(log.StringValue(message))
	record.SetSeverity(log.SeverityInfo)
	record.AddAttributes(attrs...)

	l.logger.Emit(ctx, record)
}

// Error enregistre un message d'erreur
func (l *Logger) Error(ctx context.Context, message string, attrs ...log.KeyValue) {
	record := log.Record{}
	record.SetBody(log.StringValue(message))
	record.SetSeverity(log.SeverityError)
	record.AddAttributes(attrs...)

	l.logger.Emit(ctx, record)
}

// Warn enregistre un message d'avertissement
func (l *Logger) Warn(ctx context.Context, message string, attrs ...log.KeyValue) {
	record := log.Record{}
	record.SetBody(log.StringValue(message))
	record.SetSeverity(log.SeverityWarn)
	record.AddAttributes(attrs...)

	l.logger.Emit(ctx, record)
}

// Debug enregistre un message de debug
func (l *Logger) Debug(ctx context.Context, message string, attrs ...log.KeyValue) {
	record := log.Record{}
	record.SetBody(log.StringValue(message))
	record.SetSeverity(log.SeverityDebug)
	record.AddAttributes(attrs...)

	l.logger.Emit(ctx, record)
}

// Infof enregistre un message d'information formaté (compatible avec fmt.Printf)
func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.Info(ctx, message)
}

// Errorf enregistre un message d'erreur formaté
func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.Error(ctx, message)
}

// Warnf enregistre un message d'avertissement formaté
func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.Warn(ctx, message)
}

// Debugf enregistre un message de debug formaté
func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.Debug(ctx, message)
}
