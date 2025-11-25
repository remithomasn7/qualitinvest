package repository

import (
	"database/sql"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type sharesOutstandingRepository struct {
	db *sql.DB
}

func NewSharesOutstandingRepository(db *sql.DB) SharesOutstandingRepository {
	return &sharesOutstandingRepository{db: db}
}

func (r *sharesOutstandingRepository) Save(companyID int, shares models.SharesOutstandings) error {
	return nil // Stub for now
}

func (r *sharesOutstandingRepository) GetByCompany(companyID int) (*models.SharesOutstandings, error) {
	return nil, nil // Stub for now
}
