package transform

import (
	"strconv"
	"strings"
	"time"

	"github.com/remithomasn7/qualitinvest/internal/models"
)

// TransformCompanyOverview converts Alpha Vantage CompanyOverview to database format
func TransformCompanyOverview(avOverview models.CompanyOverview) models.CompanyOverview {
	// Return the overview as-is since it's already in the right format
	// The repository will handle the database mapping
	return avOverview
}

// TransformIncomeStatements converts Alpha Vantage income statements to database format
func TransformIncomeStatements(avIncome models.IncomeStatements) ([]models.AnnualReportIncomeStatements, []models.QuarterlyReportIncomeStatements) {
	annual := make([]models.AnnualReportIncomeStatements, len(avIncome.AnnualReports))
	quarterly := make([]models.QuarterlyReportIncomeStatements, len(avIncome.QuarterlyReports))

	// Transform annual reports
	for i, report := range avIncome.AnnualReports {
		annual[i] = models.AnnualReportIncomeStatements{
			FiscalDateEnding:                  report.FiscalDateEnding,
			ReportedCurrency:                  report.ReportedCurrency,
			GrossProfit:                       report.GrossProfit,
			TotalRevenue:                      report.TotalRevenue,
			CostOfRevenue:                     report.CostOfRevenue,
			CostOfGoodsAndServicesSold:        report.CostOfGoodsAndServicesSold,
			OperatingIncome:                   report.OperatingIncome,
			SellingGeneralAndAdministrative:   report.SellingGeneralAndAdministrative,
			ResearchAndDevelopment:            report.ResearchAndDevelopment,
			OperatingExpenses:                 report.OperatingExpenses,
			NetInterestIncome:                 report.NetInterestIncome,
			InterestIncome:                    report.InterestIncome,
			InterestExpense:                   report.InterestExpense,
			NonInterestIncome:                 report.NonInterestIncome,
			OtherNonOperatingIncome:           report.OtherNonOperatingIncome,
			Depreciation:                      report.Depreciation,
			DepreciationAndAmortization:       report.DepreciationAndAmortization,
			IncomeBeforeTax:                   report.IncomeBeforeTax,
			IncomeTaxExpense:                  report.IncomeTaxExpense,
			InterestAndDebtExpense:            report.InterestAndDebtExpense,
			NetIncomeFromContinuingOperations: report.NetIncomeFromContinuingOperations,
			ComprehensiveIncomeNetOfTax:       report.ComprehensiveIncomeNetOfTax,
			Ebit:                              report.Ebit,
			Ebitda:                            report.Ebitda,
			NetIncome:                         report.NetIncome,
		}
	}

	// Transform quarterly reports
	for i, report := range avIncome.QuarterlyReports {
		quarterly[i] = models.QuarterlyReportIncomeStatements{
			FiscalDateEnding:                  report.FiscalDateEnding,
			ReportedCurrency:                  report.ReportedCurrency,
			GrossProfit:                       report.GrossProfit,
			TotalRevenue:                      report.TotalRevenue,
			CostOfRevenue:                     report.CostOfRevenue,
			CostOfGoodsAndServicesSold:        report.CostOfGoodsAndServicesSold,
			OperatingIncome:                   report.OperatingIncome,
			SellingGeneralAndAdministrative:   report.SellingGeneralAndAdministrative,
			ResearchAndDevelopment:            report.ResearchAndDevelopment,
			OperatingExpenses:                 report.OperatingExpenses,
			NetInterestIncome:                 report.NetInterestIncome,
			InterestIncome:                    report.InterestIncome,
			InterestExpense:                   report.InterestExpense,
			NonInterestIncome:                 report.NonInterestIncome,
			OtherNonOperatingIncome:           report.OtherNonOperatingIncome,
			Depreciation:                      report.Depreciation,
			DepreciationAndAmortization:       report.DepreciationAndAmortization,
			IncomeBeforeTax:                   report.IncomeBeforeTax,
			IncomeTaxExpense:                  report.IncomeTaxExpense,
			InterestAndDebtExpense:            report.InterestAndDebtExpense,
			NetIncomeFromContinuingOperations: report.NetIncomeFromContinuingOperations,
			ComprehensiveIncomeNetOfTax:       report.ComprehensiveIncomeNetOfTax,
			Ebit:                              report.Ebit,
			Ebitda:                            report.Ebitda,
			NetIncome:                         report.NetIncome,
		}
	}

	return annual, quarterly
}

// TransformBalanceSheet converts Alpha Vantage balance sheet to database format
func TransformBalanceSheet(avBalance models.BalanceSheet) ([]models.BalanceSheetAnnualReport, []models.BalanceSheetQuarterlyReport) {
	annual := make([]models.BalanceSheetAnnualReport, len(avBalance.AnnualReports))
	quarterly := make([]models.BalanceSheetQuarterlyReport, len(avBalance.QuarterlyReports))

	// Transform annual reports
	for i, report := range avBalance.AnnualReports {
		annual[i] = models.BalanceSheetAnnualReport{
			FiscalDateEnding:                       report.FiscalDateEnding,
			ReportedCurrency:                       report.ReportedCurrency,
			TotalAssets:                            report.TotalAssets,
			TotalCurrentAssets:                     report.TotalCurrentAssets,
			CashAndCashEquivalentsAtCarryingValue:  report.CashAndCashEquivalentsAtCarryingValue,
			CashAndShortTermInvestments:            report.CashAndShortTermInvestments,
			Inventory:                              report.Inventory,
			CurrentNetReceivables:                  report.CurrentNetReceivables,
			TotalNonCurrentAssets:                  report.TotalNonCurrentAssets,
			PropertyPlantEquipment:                 report.PropertyPlantEquipment,
			AccumulatedDepreciationAmortizationPPE: report.AccumulatedDepreciationAmortizationPPE,
			IntangibleAssets:                       report.IntangibleAssets,
			IntangibleAssetsExcludingGoodwill:      report.IntangibleAssetsExcludingGoodwill,
			Goodwill:                               report.Goodwill,
			Investments:                            report.Investments,
			LongTermInvestments:                    report.LongTermInvestments,
			ShortTermInvestments:                   report.ShortTermInvestments,
			OtherCurrentAssets:                     report.OtherCurrentAssets,
			OtherNonCurrentAssets:                  report.OtherNonCurrentAssets,
			TotalLiabilities:                       report.TotalLiabilities,
			TotalCurrentLiabilities:                report.TotalCurrentLiabilities,
			CurrentAccountsPayable:                 report.CurrentAccountsPayable,
			DeferredRevenue:                        report.DeferredRevenue,
			CurrentDebt:                            report.CurrentDebt,
			ShortTermDebt:                          report.ShortTermDebt,
			TotalNonCurrentLiabilities:             report.TotalNonCurrentLiabilities,
			CapitalLeaseObligations:                report.CapitalLeaseObligations,
			LongTermDebt:                           report.LongTermDebt,
			CurrentLongTermDebt:                    report.CurrentLongTermDebt,
			LongTermDebtNoncurrent:                 report.LongTermDebtNoncurrent,
			ShortLongTermDebtTotal:                 report.ShortLongTermDebtTotal,
			OtherCurrentLiabilities:                report.OtherCurrentLiabilities,
			OtherNonCurrentLiabilities:             report.OtherNonCurrentLiabilities,
			TotalShareholderEquity:                 report.TotalShareholderEquity,
			TreasuryStock:                          report.TreasuryStock,
			RetainedEarnings:                       report.RetainedEarnings,
			CommonStock:                            report.CommonStock,
			CommonStockSharesOutstanding:           report.CommonStockSharesOutstanding,
		}
	}

	// Transform quarterly reports (similar structure)
	for i, report := range avBalance.QuarterlyReports {
		quarterly[i] = models.BalanceSheetQuarterlyReport{
			FiscalDateEnding:                       report.FiscalDateEnding,
			ReportedCurrency:                       report.ReportedCurrency,
			TotalAssets:                            report.TotalAssets,
			TotalCurrentAssets:                     report.TotalCurrentAssets,
			CashAndCashEquivalentsAtCarryingValue:  report.CashAndCashEquivalentsAtCarryingValue,
			CashAndShortTermInvestments:            report.CashAndShortTermInvestments,
			Inventory:                              report.Inventory,
			CurrentNetReceivables:                  report.CurrentNetReceivables,
			TotalNonCurrentAssets:                  report.TotalNonCurrentAssets,
			PropertyPlantEquipment:                 report.PropertyPlantEquipment,
			AccumulatedDepreciationAmortizationPPE: report.AccumulatedDepreciationAmortizationPPE,
			IntangibleAssets:                       report.IntangibleAssets,
			IntangibleAssetsExcludingGoodwill:      report.IntangibleAssetsExcludingGoodwill,
			Goodwill:                               report.Goodwill,
			Investments:                            report.Investments,
			LongTermInvestments:                    report.LongTermInvestments,
			ShortTermInvestments:                   report.ShortTermInvestments,
			OtherCurrentAssets:                     report.OtherCurrentAssets,
			OtherNonCurrentAssets:                  report.OtherNonCurrentAssets,
			TotalLiabilities:                       report.TotalLiabilities,
			TotalCurrentLiabilities:                report.TotalCurrentLiabilities,
			CurrentAccountsPayable:                 report.CurrentAccountsPayable,
			DeferredRevenue:                        report.DeferredRevenue,
			CurrentDebt:                            report.CurrentDebt,
			ShortTermDebt:                          report.ShortTermDebt,
			TotalNonCurrentLiabilities:             report.TotalNonCurrentLiabilities,
			CapitalLeaseObligations:                report.CapitalLeaseObligations,
			LongTermDebt:                           report.LongTermDebt,
			CurrentLongTermDebt:                    report.CurrentLongTermDebt,
			LongTermDebtNoncurrent:                 report.LongTermDebtNoncurrent,
			ShortLongTermDebtTotal:                 report.ShortLongTermDebtTotal,
			OtherCurrentLiabilities:                report.OtherCurrentLiabilities,
			OtherNonCurrentLiabilities:             report.OtherNonCurrentLiabilities,
			TotalShareholderEquity:                 report.TotalShareholderEquity,
			TreasuryStock:                          report.TreasuryStock,
			RetainedEarnings:                       report.RetainedEarnings,
			CommonStock:                            report.CommonStock,
			CommonStockSharesOutstanding:           report.CommonStockSharesOutstanding,
		}
	}

	return annual, quarterly
}

// TransformCashFlow converts Alpha Vantage cash flow to database format
func TransformCashFlow(avCashFlow models.CashFlowStatements) ([]models.CashFlowStatementsAnnualReport, []models.CashFlowStatementsQuarterlyReport) {
	annual := make([]models.CashFlowStatementsAnnualReport, len(avCashFlow.AnnualReports))
	quarterly := make([]models.CashFlowStatementsQuarterlyReport, len(avCashFlow.QuarterlyReports))

	// Transform annual reports
	for i, report := range avCashFlow.AnnualReports {
		annual[i] = models.CashFlowStatementsAnnualReport{
			FiscalDateEnding:                                          report.FiscalDateEnding,
			ReportedCurrency:                                          report.ReportedCurrency,
			OperatingCashflow:                                         report.OperatingCashflow,
			PaymentsForOperatingActivities:                            report.PaymentsForOperatingActivities,
			ProceedsFromOperatingActivities:                           report.ProceedsFromOperatingActivities,
			ChangeInOperatingLiabilities:                              report.ChangeInOperatingLiabilities,
			ChangeInOperatingAssets:                                   report.ChangeInOperatingAssets,
			DepreciationDepletionAndAmortization:                      report.DepreciationDepletionAndAmortization,
			CapitalExpenditures:                                       report.CapitalExpenditures,
			ChangeInReceivables:                                       report.ChangeInReceivables,
			ChangeInInventory:                                         report.ChangeInInventory,
			ProfitLoss:                                               report.ProfitLoss,
			CashflowFromInvestment:                                    report.CashflowFromInvestment,
			CashflowFromFinancing:                                     report.CashflowFromFinancing,
			ProceedsFromRepaymentsOfShortTermDebt:                     report.ProceedsFromRepaymentsOfShortTermDebt,
			PaymentsForRepurchaseOfCommonStock:                        report.PaymentsForRepurchaseOfCommonStock,
			PaymentsForRepurchaseOfEquity:                             report.PaymentsForRepurchaseOfEquity,
			PaymentsForRepurchaseOfPreferredStock:                     report.PaymentsForRepurchaseOfPreferredStock,
			DividendPayout:                                           report.DividendPayout,
			DividendPayoutCommonStock:                                 report.DividendPayoutCommonStock,
			DividendPayoutPreferredStock:                              report.DividendPayoutPreferredStock,
			ProceedsFromIssuanceOfCommonStock:                         report.ProceedsFromIssuanceOfCommonStock,
			ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet: report.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
			ProceedsFromIssuanceOfPreferredStock:                      report.ProceedsFromIssuanceOfPreferredStock,
			ProceedsFromRepurchaseOfEquity:                            report.ProceedsFromRepurchaseOfEquity,
			ProceedsFromSaleOfTreasuryStock:                           report.ProceedsFromSaleOfTreasuryStock,
			ChangeInCashAndCashEquivalents:                            report.ChangeInCashAndCashEquivalents,
			ChangeInExchangeRate:                                      report.ChangeInExchangeRate,
			NetIncome:                                                report.NetIncome,
		}
	}

	// Transform quarterly reports
	for i, report := range avCashFlow.QuarterlyReports {
		quarterly[i] = models.CashFlowStatementsQuarterlyReport{
			FiscalDateEnding:                                          report.FiscalDateEnding,
			ReportedCurrency:                                          report.ReportedCurrency,
			OperatingCashflow:                                         report.OperatingCashflow,
			PaymentsForOperatingActivities:                            report.PaymentsForOperatingActivities,
			ProceedsFromOperatingActivities:                           report.ProceedsFromOperatingActivities,
			ChangeInOperatingLiabilities:                              report.ChangeInOperatingLiabilities,
			ChangeInOperatingAssets:                                   report.ChangeInOperatingAssets,
			DepreciationDepletionAndAmortization:                      report.DepreciationDepletionAndAmortization,
			CapitalExpenditures:                                       report.CapitalExpenditures,
			ChangeInReceivables:                                       report.ChangeInReceivables,
			ChangeInInventory:                                         report.ChangeInInventory,
			ProfitLoss:                                               report.ProfitLoss,
			CashflowFromInvestment:                                    report.CashflowFromInvestment,
			CashflowFromFinancing:                                     report.CashflowFromFinancing,
			ProceedsFromRepaymentsOfShortTermDebt:                     report.ProceedsFromRepaymentsOfShortTermDebt,
			PaymentsForRepurchaseOfCommonStock:                        report.PaymentsForRepurchaseOfCommonStock,
			PaymentsForRepurchaseOfEquity:                             report.PaymentsForRepurchaseOfEquity,
			PaymentsForRepurchaseOfPreferredStock:                     report.PaymentsForRepurchaseOfPreferredStock,
			DividendPayout:                                           report.DividendPayout,
			DividendPayoutCommonStock:                                 report.DividendPayoutCommonStock,
			DividendPayoutPreferredStock:                              report.DividendPayoutPreferredStock,
			ProceedsFromIssuanceOfCommonStock:                         report.ProceedsFromIssuanceOfCommonStock,
			ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet: report.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
			ProceedsFromIssuanceOfPreferredStock:                      report.ProceedsFromIssuanceOfPreferredStock,
			ProceedsFromRepurchaseOfEquity:                            report.ProceedsFromRepurchaseOfEquity,
			ProceedsFromSaleOfTreasuryStock:                           report.ProceedsFromSaleOfTreasuryStock,
			ChangeInCashAndCashEquivalents:                            report.ChangeInCashAndCashEquivalents,
			ChangeInExchangeRate:                                      report.ChangeInExchangeRate,
			NetIncome:                                                report.NetIncome,
		}
	}

	return annual, quarterly
}

// Helper functions for data conversion
func parseDate(dateStr string) time.Time {
	// Alpha Vantage dates are typically in YYYY-MM-DD format
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t
	}
	// Try other common formats
	if t, err := time.Parse("2006-01-02 15:04:05", dateStr); err == nil {
		return t
	}
	// Return zero time if parsing fails
	return time.Time{}
}

func parseFloat(s string) float64 {
	if s == "" || s == "None" {
		return 0
	}
	if f, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err == nil {
		return f
	}
	return 0
}

func parseInt(s string) int64 {
	if s == "" || s == "None" {
		return 0
	}
	if i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64); err == nil {
		return i
	}
	return 0
}
