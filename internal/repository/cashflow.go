package repository

import (
	"database/sql"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type cashFlowRepository struct {
	db *sql.DB
}

func NewCashFlowRepository(db *sql.DB) CashFlowRepository {
	return &cashFlowRepository{db: db}
}

func (r *cashFlowRepository) SaveAnnual(companyID int, statements []models.CashFlowStatementsAnnualReport) error {
	return nil // Stub for now
}

func (r *cashFlowRepository) SaveQuarterly(companyID int, statements []models.CashFlowStatementsQuarterlyReport) error {
	return nil // Stub for now
}

func (r *cashFlowRepository) GetAnnualByCompany(companyID int) ([]models.CashFlowStatementsAnnualReport, error) {
	return nil, nil // Stub for now
}

func (r *cashFlowRepository) GetQuarterlyByCompany(companyID int) ([]models.CashFlowStatementsQuarterlyReport, error) {
	return nil, nil // Stub for now
}
