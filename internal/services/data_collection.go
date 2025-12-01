package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/repository"
	"github.com/remithomasn7/qualitinvest/pkg"
	otelog "go.opentelemetry.io/otel/log"
)

// DataCollectionService gère la collecte intelligente des données
type DataCollectionService struct {
	logger       *pkg.Logger
	apiClient    *alpha_vantage.AlphaVantageClient
	companyRepo  repository.CompanyRepository
	overviewRepo repository.CompanyOverviewRepository
	incomeRepo   repository.IncomeRepository
	balanceRepo  repository.BalanceRepository
	cashFlowRepo repository.CashFlowRepository
	earningsRepo repository.EarningsRepository
}

// NewDataCollectionService crée un nouveau service de collecte
func NewDataCollectionService(db *sql.DB, apiClient *alpha_vantage.AlphaVantageClient, logger *pkg.Logger) *DataCollectionService {
	return &DataCollectionService{
		logger:       logger,
		apiClient:    apiClient,
		companyRepo:  repository.NewCompanyRepository(db),
		overviewRepo: repository.NewCompanyOverviewRepository(db),
		incomeRepo:   repository.NewIncomeRepository(db),
		balanceRepo:  repository.NewBalanceRepository(db),
		cashFlowRepo: repository.NewCashFlowRepository(db),
		earningsRepo: repository.NewEarningsRepository(db),
	}
}

// CollectCompanyData collecte toutes les données d'une entreprise depuis Alpha Vantage
// Utilisé pour les nouvelles entreprises ou mises à jour complètes
func (s *DataCollectionService) CollectCompanyData(symbol string) error {
	ctx := context.Background()

	s.logger.Info(ctx, "🚀 Début de la collecte de données depuis Alpha Vantage",
		otelog.String("symbol", symbol))

	// 1. Collecter les données overview (métadonnées + ratios)
	s.logger.Debug(ctx, "📡 Appel API Alpha Vantage - Company Overview",
		otelog.String("symbol", symbol))

	overview, err := s.apiClient.CompanyOverview(symbol)
	if err != nil {
		s.logger.Error(ctx, "❌ Échec de récupération des données overview depuis Alpha Vantage",
			otelog.String("symbol", symbol),
			otelog.String("error", err.Error()))
		return fmt.Errorf("failed to fetch company overview: %w", err)
	}

	s.logger.Info(ctx, "✅ Données overview récupérées avec succès",
		otelog.String("symbol", symbol),
		otelog.String("company_name", overview.Name))

	// 2. Valider que les données sont valides
	// Alpha Vantage peut retourner des données vides pour les symboles invalides
	// ou un message d'information si la clé API est démo
	if overview.Symbol == "" || overview.Name == "" {
		log.Printf("Invalid or empty data received for symbol %s - company not found", symbol)
		return fmt.Errorf("invalid symbol or company not found: %s", symbol)
	}

	// 3. Vérifier que le symbole retourné correspond à celui demandé
	// (parfois Alpha Vantage peut corriger ou retourner un symbole différent)
	if strings.ToUpper(overview.Symbol) != strings.ToUpper(symbol) {
		log.Printf("Symbol mismatch: requested %s but got %s", symbol, overview.Symbol)
		return fmt.Errorf("symbol mismatch: requested %s but API returned %s", symbol, overview.Symbol)
	}

	// 3. Créer/mettre à jour la compagnie
	company, err := s.companyRepo.CreateOrUpdate(
		overview.Symbol,
		overview.Name,
		overview.Exchange,
		overview.Sector,
		overview.Industry,
		overview.Country,
		overview.Currency,
	)
	if err != nil {
		log.Printf("Failed to create company %s: %v", symbol, err)
		return fmt.Errorf("failed to save company data: %w", err)
	}

	// Sauvegarder l'overview
	if err := s.overviewRepo.Save(*overview); err != nil {
		log.Printf("Warning: Failed to save overview for %s: %v", symbol, err)
	}

	// 2. Collecter les états financiers (synchrone pour l'instant)
	s.collectFinancialStatements(symbol, company.ID)

	log.Printf("Data collection completed for %s", symbol)
	return nil
}

// collectFinancialStatements collecte tous les états financiers en arrière-plan
func (s *DataCollectionService) collectFinancialStatements(symbol string, companyID int) {
	log.Printf("Collecting financial statements for %s", symbol)

	// Earnings
	if earnings, err := s.apiClient.Earnings(symbol); err == nil {
		if err := s.earningsRepo.SaveAnnual(companyID, earnings.AnnualEarnings); err != nil {
			log.Printf("Failed to save earnings for %s: %v", symbol, err)
		}
		if err := s.earningsRepo.SaveQuarterly(companyID, earnings.QuarterlyEarnings); err != nil {
			log.Printf("Failed to save quarterly earnings for %s: %v", symbol, err)
		}
	} else {
		log.Printf("Failed to fetch earnings for %s: %v", symbol, err)
	}

	log.Printf("Financial statements collection completed for %s", symbol)
}

// UpdateCompanyOverview met à jour seulement l'overview (léger, fréquent)
func (s *DataCollectionService) UpdateCompanyOverview(symbol string) error {
	overview, err := s.apiClient.CompanyOverview(symbol)
	if err != nil {
		return err
	}

	return s.overviewRepo.Save(*overview)
}

// BatchUpdateOverviews met à jour les overviews pour toutes les entreprises
// À exécuter quotidiennement
func (s *DataCollectionService) BatchUpdateOverviews() error {
	log.Println("Starting batch overview update")

	companies, err := s.companyRepo.List(1000, 0) // Toutes les entreprises
	if err != nil {
		return err
	}

	for _, company := range companies {
		if err := s.UpdateCompanyOverview(company.Symbol); err != nil {
			log.Printf("Failed to update overview for %s: %v", company.Symbol, err)
			// Continue avec les autres
		} else {
			log.Printf("Updated overview for %s", company.Symbol)
		}

		// Petite pause pour respecter les limites API
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("Batch overview update completed")
	return nil
}

// IsDataStale vérifie si les données d'une entreprise sont obsolètes
func (s *DataCollectionService) IsDataStale(symbol string, maxAge time.Duration) (bool, error) {
	ctx := context.Background()

	s.logger.Debug(ctx, "🔍 Vérification de la fraîcheur des données",
		otelog.String("symbol", symbol),
		otelog.Float64("max_age_seconds", maxAge.Seconds()))

	overview, err := s.overviewRepo.GetBySymbol(symbol)
	if err != nil {
		s.logger.Warn(ctx, "⚠️ Erreur lors de la vérification des données en cache",
			otelog.String("symbol", symbol),
			otelog.String("error", err.Error()))
		return true, err // Si pas de données, c'est stale
	}

	if overview == nil {
		s.logger.Info(ctx, "📭 Aucune donnée trouvée en cache - données manquantes",
			otelog.String("symbol", symbol))
		return true, nil // Pas de données
	}

	s.logger.Info(ctx, "✅ Données trouvées en cache - données disponibles",
		otelog.String("symbol", symbol),
		otelog.String("latest_quarter", overview.LatestQuarter))

	// Pour l'instant, on considère comme stale si plus vieux que maxAge
	// TODO: Implémenter une vraie logique de vérification de fraîcheur

	return false, nil // Temporaire
}
