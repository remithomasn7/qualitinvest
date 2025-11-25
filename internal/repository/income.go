package repository

import (
	"database/sql"
	"fmt"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

type incomeRepository struct {
	db *sql.DB
}

func NewIncomeRepository(db *sql.DB) IncomeRepository {
	return &incomeRepository{db: db}
}

func (r *incomeRepository) SaveAnnual(companyID int, statements []models.AnnualReportIncomeStatements) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO income_annual (
			company_id, fiscal_date, reported_currency, gross_profit, total_revenue,
			cost_of_revenue, operating_income, selling_general_and_administrative,
			research_and_development, net_income, ebit, ebitda, income_tax_expense, interest_expense
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		ON CONFLICT (company_id, fiscal_date) DO UPDATE SET
			gross_profit = EXCLUDED.gross_profit,
			total_revenue = EXCLUDED.total_revenue,
			cost_of_revenue = EXCLUDED.cost_of_revenue,
			operating_income = EXCLUDED.operating_income,
			selling_general_and_administrative = EXCLUDED.selling_general_and_administrative,
			research_and_development = EXCLUDED.research_and_development,
			net_income = EXCLUDED.net_income,
			ebit = EXCLUDED.ebit,
			ebitda = EXCLUDED.ebitda,
			income_tax_expense = EXCLUDED.income_tax_expense,
			interest_expense = EXCLUDED.interest_expense`)

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, statement := range statements {
		_, err = stmt.Exec(
			companyID, statement.FiscalDateEnding, statement.ReportedCurrency,
			int(statement.GrossProfit), int(statement.TotalRevenue), int(statement.CostOfRevenue),
			int(statement.OperatingIncome), int(statement.SellingGeneralAndAdministrative),
			int(statement.ResearchAndDevelopment), int(statement.NetIncome),
			int(statement.Ebit), int(statement.Ebitda),
			int(statement.IncomeTaxExpense), int(statement.InterestExpense),
		)
		if err != nil {
			return fmt.Errorf("failed to insert income statement: %w", err)
		}
	}

	return tx.Commit()
}

func (r *incomeRepository) SaveQuarterly(companyID int, statements []models.QuarterlyReportIncomeStatements) error {
	// Similar implementation for quarterly data
	return nil // Stub for now
}

func (r *incomeRepository) GetAnnualByCompany(companyID int) ([]models.AnnualReportIncomeStatements, error) {
	return nil, nil // Stub for now
}

func (r *incomeRepository) GetQuarterlyByCompany(companyID int) ([]models.QuarterlyReportIncomeStatements, error) {
	return nil, nil // Stub for now
}
