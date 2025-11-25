package repository

import (
	"database/sql"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type earningsRepository struct {
	db *sql.DB
}

func NewEarningsRepository(db *sql.DB) EarningsRepository {
	return &earningsRepository{db: db}
}

func (r *earningsRepository) SaveAnnual(companyID int, earnings []models.AnnualEarning) error {
	return nil // Stub for now
}

func (r *earningsRepository) SaveQuarterly(companyID int, earnings []models.QuarterlyEarning) error {
	return nil // Stub for now
}

func (r *earningsRepository) GetAnnualByCompany(companyID int) ([]models.AnnualEarning, error) {
	return nil, nil // Stub for now
}

func (r *earningsRepository) GetQuarterlyByCompany(companyID int) ([]models.QuarterlyEarning, error) {
	return nil, nil // Stub for now
}
