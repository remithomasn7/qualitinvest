package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/remithomasn7/qualitinvest/internal/alpha_vantage"
	"github.com/remithomasn7/qualitinvest/internal/repository"
	"github.com/remithomasn7/qualitinvest/pkg"
	otelog "go.opentelemetry.io/otel/log"
)

// SearchService gère la recherche d'entreprises
type SearchService struct {
	logger                *pkg.Logger
	apiClient             *alpha_vantage.AlphaVantageClient
	companyRepo           repository.CompanyRepository
	overviewRepo          repository.CompanyOverviewRepository
	dataCollectionService *DataCollectionService
}

// NewSearchService crée un nouveau service de recherche
func NewSearchService(db *sql.DB, apiClient *alpha_vantage.AlphaVantageClient, dataCollectionService *DataCollectionService, logger *pkg.Logger) *SearchService {
	return &SearchService{
		logger:                logger,
		apiClient:             apiClient,
		companyRepo:           repository.NewCompanyRepository(db),
		overviewRepo:          repository.NewCompanyOverviewRepository(db),
		dataCollectionService: dataCollectionService,
	}
}

// SearchResult représente un résultat de recherche enrichi avec les données de la base
type SearchResult struct {
	Symbol      string  `json:"symbol"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Region      string  `json:"region"`
	Currency    string  `json:"currency"`
	MatchScore  float64 `json:"matchScore"`
	// Données enrichies depuis la base (si disponibles)
	Country     string `json:"country,omitempty"`
	Sector      string `json:"sector,omitempty"`
	MarketCap   *int64 `json:"marketCap,omitempty"`
	InDatabase  bool   `json:"inDatabase"` // Indique si l'entreprise est déjà en base
}

// SearchCompanies recherche des entreprises via Alpha Vantage et enrichit avec les données locales
func (s *SearchService) SearchCompanies(ctx context.Context, keywords string) ([]SearchResult, error) {
	s.logger.Info(ctx, "🔍 Recherche d'entreprises",
		otelog.String("keywords", keywords))

	// 1. Rechercher via Alpha Vantage
	searchResponse, err := s.apiClient.SymbolSearch(keywords)
	if err != nil {
		s.logger.Error(ctx, "❌ Erreur lors de la recherche Alpha Vantage",
			otelog.String("keywords", keywords),
			otelog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to search symbols: %w", err)
	}

	if len(searchResponse.BestMatches) == 0 {
		s.logger.Info(ctx, "ℹ️ Aucun résultat trouvé",
			otelog.String("keywords", keywords))
		return []SearchResult{}, nil
	}

	s.logger.Info(ctx, "✅ Résultats Alpha Vantage obtenus",
		otelog.String("keywords", keywords),
		otelog.Int("count", len(searchResponse.BestMatches)))

	// 2. Enrichir chaque résultat avec les données de la base
	results := make([]SearchResult, 0, len(searchResponse.BestMatches))
	
	for _, match := range searchResponse.BestMatches {
		result := SearchResult{
			Symbol:     match.Symbol,
			Name:       match.Name,
			Type:       match.Type,
			Region:     match.Region,
			Currency:   match.Currency,
			MatchScore: match.MatchScore,
			InDatabase: false,
		}

		// Vérifier si l'entreprise existe en base
		company, err := s.companyRepo.GetBySymbol(match.Symbol)
		if err != nil {
			s.logger.Warn(ctx, "⚠️ Erreur lors de la vérification en base",
				otelog.String("symbol", match.Symbol),
				otelog.String("error", err.Error()))
			// Continuer même en cas d'erreur
		} else if company != nil {
			// Entreprise trouvée en base - enrichir les données
			result.InDatabase = true
			result.Country = company.Country
			result.Sector = company.Sector

			// Récupérer le market cap depuis l'overview si disponible
			overview, err := s.overviewRepo.GetBySymbol(match.Symbol)
			if err == nil && overview != nil && overview.MarketCapitalization > 0 {
				marketCap := int64(overview.MarketCapitalization)
				result.MarketCap = &marketCap
			}
		} else {
			// Entreprise non trouvée - déclencher la collecte en arrière-plan
			s.logger.Info(ctx, "🔄 Entreprise non trouvée en base - collecte en arrière-plan",
				otelog.String("symbol", match.Symbol))
			
			// Collecte asynchrone pour ne pas bloquer la réponse
			go func(symbol string) {
				collectCtx := context.Background()
				if err := s.dataCollectionService.CollectCompanyData(symbol); err != nil {
					s.logger.Warn(collectCtx, "⚠️ Échec de la collecte automatique",
						otelog.String("symbol", symbol),
						otelog.String("error", err.Error()))
				} else {
					s.logger.Info(collectCtx, "✅ Collecte automatique réussie",
						otelog.String("symbol", symbol))
				}
			}(match.Symbol)
		}

		results = append(results, result)
	}

	s.logger.Info(ctx, "✅ Recherche terminée",
		otelog.String("keywords", keywords),
		otelog.Int("total_results", len(results)),
		otelog.Int("in_database", countInDatabase(results)))

	return results, nil
}

// countInDatabase compte le nombre de résultats déjà en base
func countInDatabase(results []SearchResult) int {
	count := 0
	for _, r := range results {
		if r.InDatabase {
			count++
		}
	}
	return count
}