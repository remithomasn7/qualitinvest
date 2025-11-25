package repository

import (
	"database/sql"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type splitsRepository struct {
	db *sql.DB
}

func NewSplitsRepository(db *sql.DB) SplitsRepository {
	return &splitsRepository{db: db}
}

func (r *splitsRepository) Save(companyID int, splits models.Splits) error {
	return nil // Stub for now
}

func (r *splitsRepository) GetByCompany(companyID int) (*models.Splits, error) {
	return nil, nil // Stub for now
}
