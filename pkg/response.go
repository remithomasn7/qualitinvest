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

// ROCEResponse représente la réponse pour le ratio ROCE.
type ROCEResponse struct {
	Status string        `json:"status"` // Statut de la réponse
	Data   []models.ROCE `json:"data"`   // Données sur le ROCE
}

// MarginResponse représente la réponse pour les marges financières.
type MarginResponse struct {
	Status string           `json:"status"` // Statut de la réponse
	Data   []models.Margins `json:"data"`   // Données sur les marges
}

// GrowthResponse représente la réponse pour la croissance financière.
type GrowthResponse struct {
	Status string          `json:"status"` // Statut de la réponse
	Data   []models.Growth `json:"data"`   // Données sur la croissance
}

// FreeCashFlowResponse représente la réponse pour le Free Cash Flow.
type FreeCashFlowResponse struct {
	Status string                `json:"status"` // Statut de la réponse
	Data   []models.FreeCashFlow `json:"data"`   // Données sur le Free Cash Flow
}

// ShareDilutionResponse représente la réponse pour la dilution des actions.
type ShareDilutionResponse struct {
	Status string                 `json:"status"` // Statut de la réponse
	Data   []models.ShareDilution `json:"data"`   // Données sur la dilution des actions
}

// DebtEBITDAResponse représente la réponse pour la dette et l'EBITDA.
type DebtEBITDAResponse struct {
	Status string              `json:"status"` // Statut de la réponse
	Data   []models.DebtEBITDA `json:"data"`   // Données sur la dette et l'EBITDA
}

// RDExpensesResponse représente la réponse pour les dépenses en R&D.
type RDExpensesResponse struct {
	Status string              `json:"status"` // Statut de la réponse
	Data   []models.RDExpenses `json:"data"`   // Données sur les dépenses en R&D
}
