package models

// ScreeningResult représente le résultat d'un screening d'entreprise
type ScreeningResult struct {
	CompanyID    int      `json:"company_id"`
	Symbol       string   `json:"symbol"`
	Name         string   `json:"name"`
	Sector       string   `json:"sector"`
	Industry     string   `json:"industry"`
	PERatio      *float64 `json:"pe_ratio"`
	PBRatio      *float64 `json:"pb_ratio"`
	PriceToSales *float64 `json:"price_to_sales"`
	ProfitMargin *float64 `json:"profit_margin"`
	ROE          *float64 `json:"roe"`
	ROA          *float64 `json:"roa"`
	MarketCap    *int64   `json:"market_cap"`
	Beta         *float64 `json:"beta"`
	Score        float64  `json:"score"` // Score calculé basé sur les critères
}

// CompanyAnalysis représente une analyse complète d'une entreprise
type CompanyAnalysis struct {
	Company        Company                         `json:"company"`
	Overview       CompanyOverview                 `json:"overview"`
	LatestIncome   *AnnualReportIncomeStatements   `json:"latest_income,omitempty"`
	LatestBalance  *BalanceSheetAnnualReport       `json:"latest_balance,omitempty"`
	LatestCashFlow *CashFlowStatementsAnnualReport `json:"latest_cash_flow,omitempty"`

	// Ratios calculés
	CalculatedRatios FinancialRatios `json:"calculated_ratios"`

	// Métriques dérivées
	GrowthMetrics GrowthMetrics `json:"growth_metrics"`
}

// GrowthMetrics contient les métriques de croissance calculées
type GrowthMetrics struct {
	RevenueGrowth1Y  *float64 `json:"revenue_growth_1y"`
	RevenueGrowth3Y  *float64 `json:"revenue_growth_3y"`
	EarningsGrowth1Y *float64 `json:"earnings_growth_1y"`
	EarningsGrowth3Y *float64 `json:"earnings_growth_3y"`
}

// ValuationAnalysis contient l'analyse de valorisation
type ValuationAnalysis struct {
	Symbol           string             `json:"symbol"`
	CurrentRatios    ValuationRatios    `json:"current_ratios"`
	HistoricalRatios []HistoricalRatios `json:"historical_ratios"`
	FairValue        FairValue          `json:"fair_value"`
}

// ValuationRatios contient les ratios de valorisation actuels
type ValuationRatios struct {
	PERatio             *float64 `json:"pe_ratio"`
	PBRatio             *float64 `json:"pb_ratio"`
	PriceToSales        *float64 `json:"price_to_sales"`
	PriceToCashFlow     *float64 `json:"price_to_cash_flow"`
	EnterpriseValueEBIT *float64 `json:"enterprise_value_ebit"`
}

// HistoricalRatios contient les ratios historiques
type HistoricalRatios struct {
	FiscalDate string          `json:"fiscal_date"`
	Ratios     ValuationRatios `json:"ratios"`
}

// FairValue contient l'estimation de valeur fair
type FairValue struct {
	Method         string   `json:"method"`
	EstimatedValue *float64 `json:"estimated_value"`
	Upside         *float64 `json:"upside"`
}
