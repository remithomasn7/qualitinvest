package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type companyOverviewRepository struct {
	db *sql.DB
}

func NewCompanyOverviewRepository(db *sql.DB) CompanyOverviewRepository {
	return &companyOverviewRepository{db: db}
}

func (r *companyOverviewRepository) Save(overview models.CompanyOverview) error {
	query := `
		INSERT INTO company_overview (
			company_id, symbol, pe_ratio, peg_ratio, pb_ratio, price_to_sales_ratio_ttm,
			profit_margin, operating_margin_ttm, return_on_assets_ttm, return_on_equity_ttm,
			quarterly_earnings_growth_yoy, quarterly_revenue_growth_yoy,
			market_capitalization, beta, shares_outstanding, updated_at
		) VALUES (
			(SELECT id FROM companies WHERE symbol = $1), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		) ON CONFLICT (symbol)
		DO UPDATE SET
			pe_ratio = EXCLUDED.pe_ratio,
			peg_ratio = EXCLUDED.peg_ratio,
			pb_ratio = EXCLUDED.pb_ratio,
			price_to_sales_ratio_ttm = EXCLUDED.price_to_sales_ratio_ttm,
			profit_margin = EXCLUDED.profit_margin,
			operating_margin_ttm = EXCLUDED.operating_margin_ttm,
			return_on_assets_ttm = EXCLUDED.return_on_assets_ttm,
			return_on_equity_ttm = EXCLUDED.return_on_equity_ttm,
			quarterly_earnings_growth_yoy = EXCLUDED.quarterly_earnings_growth_yoy,
			quarterly_revenue_growth_yoy = EXCLUDED.quarterly_revenue_growth_yoy,
			market_capitalization = EXCLUDED.market_capitalization,
			beta = EXCLUDED.beta,
			shares_outstanding = EXCLUDED.shares_outstanding,
			updated_at = EXCLUDED.updated_at`

	_, err := r.db.Exec(query,
		overview.Symbol,
		float64Ptr(overview.PERatio), float64Ptr(overview.PEGRatio),
		float64Ptr(overview.PriceToBookRatio), float64Ptr(overview.PriceToSalesRatioTTM),
		float64Ptr(overview.ProfitMargin), float64Ptr(overview.OperatingMarginTTM),
		float64Ptr(overview.ReturnOnAssetsTTM), float64Ptr(overview.ReturnOnEquityTTM),
		float64Ptr(overview.QuarterlyEarningsGrowthYOY), float64Ptr(overview.QuarterlyRevenueGrowthYOY),
		int64Ptr(int64(overview.MarketCapitalization)), float64Ptr(overview.Beta),
		int64Ptr(int64(overview.SharesOutstanding)), time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to save company overview: %w", err)
	}

	return nil
}

func (r *companyOverviewRepository) GetBySymbol(symbol string) (*models.CompanyOverview, error) {
	query := `
		SELECT co.symbol, co.pe_ratio, co.peg_ratio, co.pb_ratio, co.price_to_sales_ratio_ttm,
		       co.profit_margin, co.operating_margin_ttm, co.return_on_assets_ttm, co.return_on_equity_ttm,
		       co.quarterly_earnings_growth_yoy, co.quarterly_revenue_growth_yoy,
		       co.market_capitalization, co.beta, co.shares_outstanding,
		       c.name, c.exchange, c.sector, c.industry, c.country, c.currency
		FROM company_overview co
		JOIN companies c ON co.company_id = c.id
		WHERE co.symbol = $1`

	var overview models.CompanyOverview
	var dbOverview models.DBCompanyOverview

	err := r.db.QueryRow(query, symbol).Scan(
		&overview.Symbol, &dbOverview.PERatio, &dbOverview.PEGRatio, &dbOverview.PBRatio,
		&dbOverview.PriceToSalesRatioTTM, &dbOverview.ProfitMargin, &dbOverview.OperatingMarginTTM,
		&dbOverview.ReturnOnAssetsTTM, &dbOverview.ReturnOnEquityTTM,
		&dbOverview.QuarterlyEarningsGrowthYOY, &dbOverview.QuarterlyRevenueGrowthYOY,
		&dbOverview.MarketCapitalization, &dbOverview.Beta, &dbOverview.SharesOutstanding,
		&overview.Name, &overview.Exchange, &overview.Sector, &overview.Industry,
		&overview.Country, &overview.Currency,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Overview not found
		}
		return nil, fmt.Errorf("failed to get company overview: %w", err)
	}

	// Convert nullable fields to the overview struct
	if dbOverview.PERatio != nil {
		overview.PERatio = *dbOverview.PERatio
	}
	if dbOverview.PEGRatio != nil {
		overview.PEGRatio = *dbOverview.PEGRatio
	}
	if dbOverview.PBRatio != nil {
		overview.PriceToBookRatio = *dbOverview.PBRatio
	}
	if dbOverview.MarketCapitalization != nil {
		overview.MarketCapitalization = int(*dbOverview.MarketCapitalization)
	}
	if dbOverview.SharesOutstanding != nil {
		overview.SharesOutstanding = int(*dbOverview.SharesOutstanding)
	}
	// Add other field mappings as needed...

	return &overview, nil
}

func (r *companyOverviewRepository) Update(symbol string, overview models.CompanyOverview) error {
	return r.Save(overview) // For now, reuse the Save method which handles upsert
}

// Helper functions for nullable types
func float64Ptr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

func int64Ptr(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}
