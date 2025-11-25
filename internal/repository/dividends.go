package repository

import (
	"database/sql"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type dividendsRepository struct {
	db *sql.DB
}

func NewDividendsRepository(db *sql.DB) DividendsRepository {
	return &dividendsRepository{db: db}
}

func (r *dividendsRepository) Save(companyID int, dividends models.Dividends) error {
	return nil // Stub for now
}

func (r *dividendsRepository) GetByCompany(companyID int) (*models.Dividends, error) {
	return nil, nil // Stub for now
}
