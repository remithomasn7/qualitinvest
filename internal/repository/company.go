package repository

import (
	"database/sql"
	"fmt"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) CreateOrUpdate(symbol, name, exchange, sector, industry, country, currency string) (*models.Company, error) {
	query := `
		INSERT INTO companies (symbol, name, exchange, sector, industry, country, currency)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (symbol)
		DO UPDATE SET
			name = EXCLUDED.name,
			exchange = EXCLUDED.exchange,
			sector = EXCLUDED.sector,
			industry = EXCLUDED.industry,
			country = EXCLUDED.country,
			currency = EXCLUDED.currency
		RETURNING id, symbol, name, exchange, sector, industry, country, currency`

	var company models.Company
	err := r.db.QueryRow(query, symbol, name, exchange, sector, industry, country, currency).Scan(
		&company.ID, &company.Symbol, &company.Name, &company.Exchange,
		&company.Sector, &company.Industry, &company.Country, &company.Currency,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create/update company: %w", err)
	}

	return &company, nil
}

func (r *companyRepository) GetBySymbol(symbol string) (*models.Company, error) {
	query := `SELECT id, symbol, name, exchange, sector, industry, country, currency FROM companies WHERE symbol = $1`

	var company models.Company
	err := r.db.QueryRow(query, symbol).Scan(
		&company.ID, &company.Symbol, &company.Name, &company.Exchange,
		&company.Sector, &company.Industry, &company.Country, &company.Currency,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Company not found
		}
		return nil, fmt.Errorf("failed to get company by symbol: %w", err)
	}

	return &company, nil
}

func (r *companyRepository) GetByID(id int) (*models.Company, error) {
	query := `SELECT id, symbol, name, exchange, sector, industry, country, currency FROM companies WHERE id = $1`

	var company models.Company
	err := r.db.QueryRow(query, id).Scan(
		&company.ID, &company.Symbol, &company.Name, &company.Exchange,
		&company.Sector, &company.Industry, &company.Country, &company.Currency,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Company not found
		}
		return nil, fmt.Errorf("failed to get company by ID: %w", err)
	}

	return &company, nil
}

func (r *companyRepository) List(limit, offset int) ([]models.Company, error) {
	query := `SELECT id, symbol, name, exchange, sector, industry, country, currency FROM companies ORDER BY symbol LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list companies: %w", err)
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var company models.Company
		err := rows.Scan(
			&company.ID, &company.Symbol, &company.Name, &company.Exchange,
			&company.Sector, &company.Industry, &company.Country, &company.Currency,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan company: %w", err)
		}
		companies = append(companies, company)
	}

	return companies, nil
}
