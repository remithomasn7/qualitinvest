package pkg

import "github.com/remithomasn7/qualitinvest/internal/models"

// APIResponse représente la structure de réponse standard pour les API.
type APIResponse struct {
	Status  string      `json:"status"`  // Statut de la réponse, e.g., "success" ou "error"
	Message string      `json:"message"` // Message d'information ou d'erreur
	Data    interface{} `json:"data"`    // Données renvoyées par l'API
}

// ErrorResponse représente la structure de réponse pour les erreurs.
type ErrorResponse struct {
	Status  string `json:"status"`  // Statut de la réponse, e.g., "error"
	Message string `json:"message"` // Message d'erreur
	Code    int    `json:"code"`    // Code d'erreur HTTP
}

// OverviewResponse représente la réponse pour l'overview.
type OverviewResponse struct {
	Status string                 `json:"status"` // The answer's status
	Data   models.CompanyOverview `json:"data"`   // The company's overview data
}

// ETFProfileResponse représente la réponse pour le profil d'un ETF.
type ETFProfileResponse struct {
	Status string            `json:"status"` // Statut de la réponse
	Data   models.ETFProfile `json:"data"`   // Données du profil ETF
}

// DividendsResponse représente la réponse pour les dividendes.
type DividendsResponse struct {
	Status string           `json:"status"` // Statut de la réponse
	Data   models.Dividends `json:"data"`   // Données sur les dividendes
}

// SplitsResponse représente la réponse pour les splits.
type SplitsResponse struct {
	Status string        `json:"status"` // Statut de la réponse
	Data   models.Splits `json:"data"`   // Données sur les splits
}

// IncomeStatementsResponse represents the answer when the company's income statements are requested.
type IncomeStatementsResponse struct {
	Status string                  `json:"status"` // The answer's status
	Data   models.IncomeStatements `json:"data"`   // Income statements data
}

// BalanceSheetResponse represents the answer when the company's balance sheet is requested.
type BalanceSheetResponse struct {
	Status string              `json:"status"` // The answer's status
	Data   models.BalanceSheet `json:"data"`   // Balance Sheet data
}

// CashFlowStatementsResponse represents the answer when the company's cash flow statements are requested.
type CashFlowStatementsResponse struct {
	Status string                    `json:"status"` // The answer's status
	Data   models.CashFlowStatements `json:"data"`   // Cash Flow Statements data
}

// SharesOutstandingsResponse représente la réponse pour les actions en circulation.
type SharesOutstandingsResponse struct {
	Status string                    `json:"status"` // The answer's status
	Data   models.SharesOutstandings `json:"data"`   // Shares Outstandings data
}

// EarningsResponse représente la réponse pour les bénéfices.
type EarningsResponse struct {
	Status string          `json:"status"` // Statut de la réponse
	Data   models.Earnings `json:"data"`   // Données sur les bénéfices
}

// ShareDilutionResponse représente la réponse pour la dilution des actions.
type ShareDilutionResponse struct {
	Status string                 `json:"status"` // Statut de la réponse
	Data   []models.ShareDilution `json:"data"`   // Données sur la dilution des actions
}
