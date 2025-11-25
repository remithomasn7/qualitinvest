package services

import (
	"database/sql"
	"fmt"

	"github.com/remithomasn7/qualitinvest/internal/models"
	"github.com/remithomasn7/qualitinvest/internal/repository"
)

// ScreeningCriteria définit les critères de filtrage pour le screening
type ScreeningCriteria struct {
	// Valuation ratios
	MinPERatio      *float64 `json:"min_pe_ratio"`
	MaxPERatio      *float64 `json:"max_pe_ratio"`
	MinPBRatio      *float64 `json:"min_pb_ratio"`
	MaxPBRatio      *float64 `json:"max_pb_ratio"`
	MinPriceToSales *float64 `json:"min_price_to_sales"`

	// Profitability
	MinROE         *float64 `json:"min_roe"`
	MinROA         *float64 `json:"min_roa"`
	MinProfitMargin *float64 `json:"min_profit_margin"`

	// Financial health
	MaxDebtToEquity *float64 `json:"max_debt_to_equity"`
	MinCurrentRatio *float64 `json:"min_current_ratio"`

	// Growth
	MinRevenueGrowth *float64 `json:"min_revenue_growth"`
	MinEPSGrowth     *float64 `json:"min_eps_growth"`

	// Size
	MinMarketCap *int64 `json:"min_market_cap"`
	MaxMarketCap *int64 `json:"max_market_cap"`

	// Sectors (optionnel)
	Sectors []string `json:"sectors"`

	// Limit results
	Limit int `json:"limit"`
}

// ScreeningService gère les opérations de screening et d'analyse
type ScreeningService struct {
	overviewRepo repository.CompanyOverviewRepository
	companyRepo  repository.CompanyRepository
	incomeRepo   repository.IncomeRepository
	balanceRepo  repository.BalanceRepository
	cashFlowRepo repository.CashFlowRepository
}

// NewScreeningService crée un nouveau service de screening
func NewScreeningService(db *sql.DB) *ScreeningService {
	return &ScreeningService{
		overviewRepo: repository.NewCompanyOverviewRepository(db),
		companyRepo:  repository.NewCompanyRepository(db),
		incomeRepo:   repository.NewIncomeRepository(db),
		balanceRepo:  repository.NewBalanceRepository(db),
		cashFlowRepo: repository.NewCashFlowRepository(db),
	}
}

// ScreenCompanies applique des critères de filtrage aux entreprises
func (s *ScreeningService) ScreenCompanies(criteria ScreeningCriteria) ([]models.ScreeningResult, error) {
	// Pour l'instant, on utilise une requête simple
	// TODO: Implémenter une vraie logique de screening avec jointures

	query := `
		SELECT
			c.id, c.symbol, c.name, c.sector, c.industry,
			co.pe_ratio, co.pb_ratio, co.price_to_sales_ratio_ttm,
			co.profit_margin, co.return_on_equity_ttm, co.return_on_assets_ttm,
			co.market_capitalization, co.beta
		FROM companies c
		JOIN company_overview co ON c.id = co.company_id
		WHERE 1=1
	`

	args := []interface{}{}
	argCount := 0

	// Appliquer les filtres
	if criteria.MinPERatio != nil {
		argCount++
		query += fmt.Sprintf(" AND co.pe_ratio >= $%d", argCount)
		args = append(args, *criteria.MinPERatio)
	}

	if criteria.MaxPERatio != nil {
		argCount++
		query += fmt.Sprintf(" AND (co.pe_ratio <= $%d OR co.pe_ratio IS NULL)", argCount)
		args = append(args, *criteria.MaxPERatio)
	}

	if criteria.MinPBRatio != nil {
		argCount++
		query += fmt.Sprintf(" AND co.pb_ratio >= $%d", argCount)
		args = append(args, *criteria.MinPBRatio)
	}

	if criteria.MaxPBRatio != nil {
		argCount++
		query += fmt.Sprintf(" AND (co.pb_ratio <= $%d OR co.pb_ratio IS NULL)", argCount)
		args = append(args, *criteria.MaxPBRatio)
	}

	if criteria.MinROE != nil {
		argCount++
		query += fmt.Sprintf(" AND co.return_on_equity_ttm >= $%d", argCount)
		args = append(args, *criteria.MinROE)
	}

	if criteria.MinMarketCap != nil {
		argCount++
		query += fmt.Sprintf(" AND co.market_capitalization >= $%d", argCount)
		args = append(args, *criteria.MinMarketCap)
	}

	if len(criteria.Sectors) > 0 {
		placeholders := ""
		for i, sector := range criteria.Sectors {
			if i > 0 {
				placeholders += ","
			}
			argCount++
			placeholders += fmt.Sprintf("$%d", argCount)
			args = append(args, sector)
		}
		query += fmt.Sprintf(" AND c.sector IN (%s)", placeholders)
	}

	// Limit
	limit := 50 // default
	if criteria.Limit > 0 && criteria.Limit <= 500 {
		limit = criteria.Limit
	}
	query += fmt.Sprintf(" LIMIT %d", limit)

	// Pour l'instant, retourner un résultat vide (implémentation simplifiée)
	// TODO: Implémenter la vraie logique de screening avec jointures
	return []models.ScreeningResult{}, nil
}

// GetCompanyAnalysis retourne une analyse complète d'une entreprise
func (s *ScreeningService) GetCompanyAnalysis(symbol string) (*models.CompanyAnalysis, error) {
	// Récupérer l'overview
	overview, err := s.overviewRepo.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	if overview == nil {
		return nil, fmt.Errorf("company not found: %s", symbol)
	}

	// Récupérer la compagnie
	company, err := s.companyRepo.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, fmt.Errorf("company not found: %s", symbol)
	}

	// Construire l'analyse
	analysis := &models.CompanyAnalysis{
		Company:  *company,
		Overview: *overview,
	}

	// Ajouter les derniers états financiers
	if latestIncome, err := s.incomeRepo.GetAnnualByCompany(company.ID); err == nil && len(latestIncome) > 0 {
		analysis.LatestIncome = &latestIncome[0]
	}

	if latestBalance, err := s.balanceRepo.GetAnnualByCompany(company.ID); err == nil && len(latestBalance) > 0 {
		analysis.LatestBalance = &latestBalance[0]
	}

	if latestCashFlow, err := s.cashFlowRepo.GetAnnualByCompany(company.ID); err == nil && len(latestCashFlow) > 0 {
		analysis.LatestCashFlow = &latestCashFlow[0]
	}

	return analysis, nil
}

// GetSectorComparison compare une entreprise avec son secteur
func (s *ScreeningService) GetSectorComparison(symbol string) (*models.SectorComparison, error) {
	// Récupérer les données de l'entreprise
	company, err := s.companyRepo.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, fmt.Errorf("company not found: %s", symbol)
	}

	// Calculer les moyennes sectorielles
	// TODO: Implémenter la logique de calcul des moyennes sectorielles

	return &models.SectorComparison{
		Company: *company,
		// SectorAverages: sectorAverages,
	}, nil
}
