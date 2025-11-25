package models

import "time"

// Company represents a company in the database
type Company struct {
	ID       int    `json:"id" db:"id"`
	Symbol   string `json:"symbol" db:"symbol"`
	Name     string `json:"name" db:"name"`
	Exchange string `json:"exchange" db:"exchange"`
	Sector   string `json:"sector" db:"sector"`
	Industry string `json:"industry" db:"industry"`
	Country  string `json:"country" db:"country"`
	Currency string `json:"currency" db:"currency"`
}

// PriceData represents historical price data
type PriceData struct {
	Date          string  `json:"date" db:"date"`
	Open          float64 `json:"open" db:"open"`
	High          float64 `json:"high" db:"high"`
	Low           float64 `json:"low" db:"low"`
	Close         float64 `json:"close" db:"close"`
	AdjustedClose float64 `json:"adjusted_close" db:"adjusted_close"`
	Volume        int64   `json:"volume" db:"volume"`
	SMA50         float64 `json:"sma_50" db:"sma_50"`
	SMA200        float64 `json:"sma_200" db:"sma_200"`
}

// FinancialRatios represents calculated financial ratios
type FinancialRatios struct {
	FiscalDate         string  `json:"fiscal_date" db:"fiscal_date"`
	GrossMargin        float64 `json:"gross_margin" db:"gross_margin"`
	OperatingMargin    float64 `json:"operating_margin" db:"operating_margin"`
	NetMargin          float64 `json:"net_margin" db:"net_margin"`
	AssetTurnover      float64 `json:"asset_turnover" db:"asset_turnover"`
	InventoryTurnover  float64 `json:"inventory_turnover" db:"inventory_turnover"`
	DebtToEquity       float64 `json:"debt_to_equity" db:"debt_to_equity"`
	DebtToAssets       float64 `json:"debt_to_assets" db:"debt_to_assets"`
	PERatio            float64 `json:"pe_ratio" db:"pe_ratio"`
	PBRatio            float64 `json:"pb_ratio" db:"pb_ratio"`
	PriceToSalesRatio  float64 `json:"price_to_sales_ratio" db:"price_to_sales_ratio"`
	RevenueGrowthYOY   float64 `json:"revenue_growth_yoy" db:"revenue_growth_yoy"`
	EarningsGrowthYOY  float64 `json:"earnings_growth_yoy" db:"earnings_growth_yoy"`
}

// Database models for repository operations
type DBCompanyOverview struct {
	ID               int       `db:"id"`
	CompanyID        int       `db:"company_id"`
	Symbol           string    `db:"symbol"`
	PERatio          *float64  `db:"pe_ratio"`
	PEGRatio         *float64  `db:"peg_ratio"`
	PBRatio          *float64  `db:"pb_ratio"`
	PriceToSalesRatioTTM *float64 `db:"price_to_sales_ratio_ttm"`
	ProfitMargin     *float64  `db:"profit_margin"`
	OperatingMarginTTM *float64 `db:"operating_margin_ttm"`
	ReturnOnAssetsTTM *float64 `db:"return_on_assets_ttm"`
	ReturnOnEquityTTM *float64 `db:"return_on_equity_ttm"`
	QuarterlyEarningsGrowthYOY *float64 `db:"quarterly_earnings_growth_yoy"`
	QuarterlyRevenueGrowthYOY *float64 `db:"quarterly_revenue_growth_yoy"`
	MarketCapitalization *int64 `db:"market_capitalization"`
	Beta              *float64  `db:"beta"`
	SharesOutstanding *int64    `db:"shares_outstanding"`
	UpdatedAt        time.Time `db:"updated_at"`
}
