package repository

import (
	"database/sql"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type balanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) BalanceRepository {
	return &balanceRepository{db: db}
}

func (r *balanceRepository) SaveAnnual(companyID int, statements []models.BalanceSheetAnnualReport) error {
	return nil // Stub for now
}

func (r *balanceRepository) SaveQuarterly(companyID int, statements []models.BalanceSheetQuarterlyReport) error {
	return nil // Stub for now
}

func (r *balanceRepository) GetAnnualByCompany(companyID int) ([]models.BalanceSheetAnnualReport, error) {
	return nil, nil // Stub for now
}

func (r *balanceRepository) GetQuarterlyByCompany(companyID int) ([]models.BalanceSheetQuarterlyReport, error) {
	return nil, nil // Stub for now
}
