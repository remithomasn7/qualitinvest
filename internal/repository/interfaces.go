package repository

import (
	"github.com/remithomasn7/qualitinvest/internal/models"
)

// CompanyRepository interface for company operations
type CompanyRepository interface {
	CreateOrUpdate(symbol, name, exchange, sector, industry, country, currency string) (*models.Company, error)
	GetBySymbol(symbol string) (*models.Company, error)
	GetByID(id int) (*models.Company, error)
	List(limit, offset int) ([]models.Company, error)
}

// EarningsRepository interface for earnings operations
type EarningsRepository interface {
	SaveAnnual(companyID int, earnings []models.AnnualEarning) error
	SaveQuarterly(companyID int, earnings []models.QuarterlyEarning) error
	GetAnnualByCompany(companyID int) ([]models.AnnualEarning, error)
	GetQuarterlyByCompany(companyID int) ([]models.QuarterlyEarning, error)
}

// IncomeRepository interface for income statement operations
type IncomeRepository interface {
	SaveAnnual(companyID int, statements []models.AnnualReportIncomeStatements) error
	SaveQuarterly(companyID int, statements []models.QuarterlyReportIncomeStatements) error
	GetAnnualByCompany(companyID int) ([]models.AnnualReportIncomeStatements, error)
	GetQuarterlyByCompany(companyID int) ([]models.QuarterlyReportIncomeStatements, error)
}

// BalanceRepository interface for balance sheet operations
type BalanceRepository interface {
	SaveAnnual(companyID int, statements []models.BalanceSheetAnnualReport) error
	SaveQuarterly(companyID int, statements []models.BalanceSheetQuarterlyReport) error
	GetAnnualByCompany(companyID int) ([]models.BalanceSheetAnnualReport, error)
	GetQuarterlyByCompany(companyID int) ([]models.BalanceSheetQuarterlyReport, error)
}

// CashFlowRepository interface for cash flow operations
type CashFlowRepository interface {
	SaveAnnual(companyID int, statements []models.CashFlowStatementsAnnualReport) error
	SaveQuarterly(companyID int, statements []models.CashFlowStatementsQuarterlyReport) error
	GetAnnualByCompany(companyID int) ([]models.CashFlowStatementsAnnualReport, error)
	GetQuarterlyByCompany(companyID int) ([]models.CashFlowStatementsQuarterlyReport, error)
}

// CompanyOverviewRepository interface for company overview operations
type CompanyOverviewRepository interface {
	Save(overview models.CompanyOverview) error
	GetBySymbol(symbol string) (*models.CompanyOverview, error)
	Update(symbol string, overview models.CompanyOverview) error
}

// PriceHistoryRepository interface for price history operations
type PriceHistoryRepository interface {
	SavePrices(companyID int, prices []models.PriceData) error
	GetByCompanyAndDateRange(companyID int, startDate, endDate string) ([]models.PriceData, error)
	GetLatestByCompany(companyID int, limit int) ([]models.PriceData, error)
}

// DividendsRepository interface for dividend operations
type DividendsRepository interface {
	Save(companyID int, dividends models.Dividends) error
	GetByCompany(companyID int) (*models.Dividends, error)
}

// SplitsRepository interface for stock split operations
type SplitsRepository interface {
	Save(companyID int, splits models.Splits) error
	GetByCompany(companyID int) (*models.Splits, error)
}

// SharesOutstandingRepository interface for shares outstanding operations
type SharesOutstandingRepository interface {
	Save(companyID int, shares models.SharesOutstandings) error
	GetByCompany(companyID int) (*models.SharesOutstandings, error)
}

// FinancialRatiosRepository interface for calculated ratios
type FinancialRatiosRepository interface {
	Save(companyID int, fiscalDate string, ratios models.FinancialRatios) error
	GetByCompany(companyID int) ([]models.FinancialRatios, error)
	GetLatestByCompany(companyID int) (*models.FinancialRatios, error)
}
