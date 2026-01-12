package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/remithomasn7/qualitinvest/internal/models"
	"github.com/remithomasn7/qualitinvest/internal/repository"
	"github.com/remithomasn7/qualitinvest/pkg"
	"go.opentelemetry.io/otel/log"
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
	MinROE          *float64 `json:"min_roe"`
	MinROA          *float64 `json:"min_roa"`
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
	logger       *pkg.Logger
	db           *sql.DB
	overviewRepo repository.CompanyOverviewRepository
	companyRepo  repository.CompanyRepository
	incomeRepo   repository.IncomeRepository
	balanceRepo  repository.BalanceRepository
	cashFlowRepo repository.CashFlowRepository
}

// NewScreeningService crée un nouveau service de screening
func NewScreeningService(db *sql.DB, logger *pkg.Logger) *ScreeningService {
	return &ScreeningService{
		logger:       logger,
		db:           db,
		overviewRepo: repository.NewCompanyOverviewRepository(db),
		companyRepo:  repository.NewCompanyRepository(db),
		incomeRepo:   repository.NewIncomeRepository(db),
		balanceRepo:  repository.NewBalanceRepository(db),
		cashFlowRepo: repository.NewCashFlowRepository(db),
	}
}

// ScreenCompanies applique des critères de filtrage aux entreprises
func (s *ScreeningService) ScreenCompanies(criteria ScreeningCriteria) ([]models.ScreeningResult, error) {
	ctx := context.Background()

	s.logger.Info(ctx, "🔍 Début du screening d'entreprises",
		log.String("criteria", fmt.Sprintf("%+v", criteria)))

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
		WHERE c.symbol != ''
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

	// Exécuter la requête
	s.logger.Debug(ctx, "🗃️ Exécution de la requête SQL",
		log.String("query", query),
		log.String("args", fmt.Sprintf("%v", args)))

	rows, err := s.db.Query(query, args...)
	if err != nil {
		s.logger.Error(ctx, "❌ Erreur lors de l'exécution de la requête",
			log.String("error", err.Error()))
		return nil, fmt.Errorf("failed to execute screening query: %w", err)
	}
	defer rows.Close()

	var results []models.ScreeningResult
	for rows.Next() {
		var result models.ScreeningResult
		err := rows.Scan(
			&result.CompanyID,
			&result.Symbol,
			&result.Name,
			&result.Sector,
			&result.Industry,
			&result.PERatio,
			&result.PBRatio,
			&result.PriceToSales,
			&result.ProfitMargin,
			&result.ROE,
			&result.ROA,
			&result.MarketCap,
			&result.Beta,
		)
		if err != nil {
			s.logger.Error(ctx, "❌ Erreur lors du scan d'un résultat",
				log.String("error", err.Error()))
			return nil, fmt.Errorf("failed to scan screening result: %w", err)
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error(ctx, "❌ Erreur lors de l'itération des résultats",
			log.String("error", err.Error()))
		return nil, fmt.Errorf("error iterating screening results: %w", err)
	}

	s.logger.Info(ctx, "✅ Screening terminé",
		log.Int("results_count", len(results)),
		log.String("final_query", query))

	return results, nil
}

// GetCompanyAnalysis retourne une analyse complète d'une entreprise
func (s *ScreeningService) GetCompanyAnalysis(symbol string) (*models.CompanyAnalysis, error) {
	ctx := context.Background()

	s.logger.Info(ctx, "📊 Début de l'analyse de l'entreprise",
		log.String("symbol", symbol))

	// Récupérer l'overview
	s.logger.Debug(ctx, "🗃️ Récupération des données overview depuis la base de données",
		log.String("symbol", symbol))

	overview, err := s.overviewRepo.GetBySymbol(symbol)
	if err != nil {
		s.logger.Error(ctx, "❌ Erreur lors de la récupération des données overview",
			log.String("symbol", symbol),
			log.String("error", err.Error()))
		return nil, err
	}
	if overview == nil {
		s.logger.Warn(ctx, "⚠️ Aucune donnée overview trouvée pour l'entreprise",
			log.String("symbol", symbol))
		return nil, fmt.Errorf("company not found: %s", symbol)
	}

	// Récupérer la compagnie
	s.logger.Debug(ctx, "🗃️ Récupération des informations de l'entreprise",
		log.String("symbol", symbol))

	company, err := s.companyRepo.GetBySymbol(symbol)
	if err != nil {
		s.logger.Error(ctx, "❌ Erreur lors de la récupération des informations de l'entreprise",
			log.String("symbol", symbol),
			log.String("error", err.Error()))
		return nil, err
	}
	if company == nil {
		s.logger.Warn(ctx, "⚠️ Aucune information d'entreprise trouvée",
			log.String("symbol", symbol))
		return nil, fmt.Errorf("company not found: %s", symbol)
	}

	s.logger.Info(ctx, "✅ Données récupérées avec succès, construction de l'analyse",
		log.String("symbol", symbol),
		log.String("company_name", overview.Name))

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
