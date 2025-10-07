package models

// ROCE représente le retour sur capital engagé.
type ROCE struct {
	Year  int     `json:"year"`
	Value float64 `json:"value"` // Le ratio ROCE
}

// Margin représente les marges financières.
type Margins struct {
	Year                   int     `json:"year"`
	GrossMargin            float64 `json:"gross_margin"`              // Marge brute
	OperatingMargin        float64 `json:"operating_margin"`          // Marge opérationnelle
	NetMargin              float64 `json:"net_margin"`                // Marge nette
	CAPEXToOperatingIncome float64 `json:"capex_to_operating_income"` // %CAPEX / Résultat opérationnel
}

// Growth représente la croissance financière.
type Growth struct {
	Year                int     `json:"year"`
	RevenueGrowth       float64 `json:"revenue_growth"`       // Croissance du chiffre d'affaires
	PredictabilityIndex float64 `json:"predictability_index"` // Indice de prédictibilité
}

// FreeCashFlow représente le Free Cash Flow par action.
type FreeCashFlow struct {
	Year        int     `json:"year"`
	FCFPerShare float64 `json:"fcf_per_share"` // Free Cash Flow par action
}

// ShareDilution représente la dilution des actions.
type ShareDilution struct {
	Year              int     `json:"year"`
	SharesOutstanding float64 `json:"shares_outstanding"` // Nombre d'actions en circulation
}

// DebtEBITDA représente la dette et l'EBITDA.
type DebtEBITDA struct {
	Debt         float64 `json:"debt"`           // Montant de la dette
	EBITDA       float64 `json:"ebitda"`         // Montant de l'EBITDA
	DebtToEBITDA float64 `json:"debt_to_ebitda"` // Ratio dette/EBITDA
}

// RDExpenses représente les dépenses en recherche et développement.
type RDExpenses struct {
	Year     int     `json:"year"`
	Expenses float64 `json:"expenses"` // Dépenses en R&D
}

type IncomeStatements struct {
	Symbol           string                            `json:"symbol"`
	AnnualReports    []AnnualReportIncomeStatements    `json:"annualReports"`
	QuarterlyReports []QuarterlyReportIncomeStatements `json:"quarterlyReports"`
}
type AnnualReportIncomeStatements struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       AVInt  `json:"grossProfit"`
	TotalRevenue                      AVInt  `json:"totalRevenue"`
	CostOfRevenue                     AVInt  `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        AVInt  `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   AVInt  `json:"operatingIncome"`
	SellingGeneralAndAdministrative   AVInt  `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            AVInt  `json:"researchAndDevelopment"`
	OperatingExpenses                 AVInt  `json:"operatingExpenses"`
	InvestmentIncomeNet               AVInt  `json:"investmentIncomeNet"` // Have to handle case where value is "None"
	NetInterestIncome                 AVInt  `json:"netInterestIncome"`
	InterestIncome                    AVInt  `json:"interestIncome"`
	InterestExpense                   AVInt  `json:"interestExpense"`
	NonInterestIncome                 AVInt  `json:"nonInterestIncome"`
	OtherNonOperatingIncome           AVInt  `json:"otherNonOperatingIncome"`
	Depreciation                      AVInt  `json:"depreciation"`
	DepreciationAndAmortization       AVInt  `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   AVInt  `json:"incomeBeforeTax"`
	IncomeTaxExpense                  AVInt  `json:"incomeTaxExpense"`
	InterestAndDebtExpense            AVInt  `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations AVInt  `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       AVInt  `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              AVInt  `json:"ebit"`
	Ebitda                            AVInt  `json:"ebitda"`
	NetIncome                         AVInt  `json:"netIncome"`
}

type QuarterlyReportIncomeStatements struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       AVInt  `json:"grossProfit"`
	TotalRevenue                      AVInt  `json:"totalRevenue"`
	CostOfRevenue                     AVInt  `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        AVInt  `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   AVInt  `json:"operatingIncome"`
	SellingGeneralAndAdministrative   AVInt  `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            AVInt  `json:"researchAndDevelopment"`
	OperatingExpenses                 AVInt  `json:"operatingExpenses"`
	InvestmentIncomeNet               AVInt  `json:"investmentIncomeNet"`
	NetInterestIncome                 AVInt  `json:"netInterestIncome"`
	InterestIncome                    AVInt  `json:"interestIncome"`
	InterestExpense                   AVInt  `json:"interestExpense"`
	NonInterestIncome                 AVInt  `json:"nonInterestIncome"`
	OtherNonOperatingIncome           AVInt  `json:"otherNonOperatingIncome"`
	Depreciation                      AVInt  `json:"depreciation"`
	DepreciationAndAmortization       AVInt  `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   AVInt  `json:"incomeBeforeTax"`
	IncomeTaxExpense                  AVInt  `json:"incomeTaxExpense"`
	InterestAndDebtExpense            AVInt  `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations AVInt  `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       AVInt  `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              AVInt  `json:"ebit"`
	Ebitda                            AVInt  `json:"ebitda"`
	NetIncome                         AVInt  `json:"netIncome"`
}

type Sector struct {
	Sector string `json:"sector"`
	Weight string `json:"weight"`
}

type Holding struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Weight      string `json:"weight"`
}

type ETFProfile struct {
	NetAssets         string    `json:"net_assets"`
	NetExpenseRatio   string    `json:"net_expense_ratio"`
	PortfolioTurnover string    `json:"portfolio_turnover"`
	DividendYield     string    `json:"dividend_yield"`
	InceptionDate     string    `json:"inception_date"`
	Leveraged         string    `json:"leveraged"`
	Sectors           []Sector  `json:"sectors"`
	Holdings          []Holding `json:"holdings"`
}

type Dividend struct {
	ExDividendDate  string `json:"ex_dividend_date"`
	DeclarationDate string `json:"declaration_date"`
	RecordDate      string `json:"record_date"`
	PaymentDate     string `json:"payment_date"`
	Amount          string `json:"amount"`
}

type Dividends struct {
	Symbol string     `json:"symbol"`
	Data   []Dividend `json:"data"`
}

type Split struct {
	EffectiveDate string `json:"effective_date"`
	SplitFactor   string `json:"split_factor"`
}

type Splits struct {
	Symbol string  `json:"symbol"`
	Data   []Split `json:"data"`
}

type SharesOutstanding struct {
	Date                     string `json:"date"`
	SharesOutstandingDiluted string `json:"shares_outstanding_diluted"`
	SharesOutstandingBasic   string `json:"shares_outstanding_basic"`
}

type SharesOutstandings struct {
	Symbol string              `json:"symbol"`
	Status string              `json:"status"`
	Data   []SharesOutstanding `json:"data"`
}

type AnnualEarning struct {
	FiscalDateEnding string `json:"fiscalDateEnding"`
	ReportedEPS      string `json:"reportedEPS"`
}

type QuarterlyEarning struct {
	FiscalDateEnding   string `json:"fiscalDateEnding"`
	ReportedDate       string `json:"reportedDate"`
	ReportedEPS        string `json:"reportedEPS"`
	EstimatedEPS       string `json:"estimatedEPS"`
	Surprise           string `json:"surprise"`
	SurprisePercentage string `json:"surprisePercentage"`
	ReportTime         string `json:"reportTime"`
}

type Earnings struct {
	Symbol            string             `json:"symbol"`
	AnnualEarnings    []AnnualEarning    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarning `json:"quarterlyEarnings"`
}

// CompanyOverview contains all details needed to get the company overview.
type CompanyOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        string  `json:"CIK"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"` // "2023-03-31"
	MarketCapitalization       int     `json:"MarketCapitalization,string"`
	EBITDA                     int     `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           float64 `json:"DividendPerShare,string"`
	DividendYield              float64 `json:"DividendYield,string"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int     `json:"RevenueTTM,string"`
	GrossProfitTTM             int     `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	Week52High                 float64 `json:"52WeekHigh,string"`
	Week52Low                  float64 `json:"52WeekLow,string"`
	MovingAverage50Day         float64 `json:"50DayMovingAverage,string"`
	MovingAverage200Day        float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int     `json:"SharesOutstanding,string"`
	DividendDate               string  `json:"DividendDate"`   //  "2023-06-10"
	ExDividendDate             string  `json:"ExDividendDate"` // "2023-05-09"
}

type BalanceSheet struct {
	Symbol           string                        `json:"symbol"`
	AnnualReports    []BalanceSheetAnnualReport    `json:"annualReports"`
	QuarterlyReports []BalanceSheetQuarterlyReport `json:"quarterlyReports"`
}
type BalanceSheetAnnualReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            AVInt  `json:"totalAssets"`
	TotalCurrentAssets                     AVInt  `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  AVInt  `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            AVInt  `json:"cashAndShortTermInvestments"`
	Inventory                              AVInt  `json:"inventory"`
	CurrentNetReceivables                  AVInt  `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  AVInt  `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 AVInt  `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE AVInt  `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       AVInt  `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      AVInt  `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               AVInt  `json:"goodwill"`
	Investments                            AVInt  `json:"investments"`
	LongTermInvestments                    AVInt  `json:"longTermInvestments"`
	ShortTermInvestments                   AVInt  `json:"shortTermInvestments"`
	OtherCurrentAssets                     AVInt  `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  AVInt  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       AVInt  `json:"totalLiabilities"`
	TotalCurrentLiabilities                AVInt  `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 AVInt  `json:"currentAccountsPayable"`
	DeferredRevenue                        AVInt  `json:"deferredRevenue"`
	CurrentDebt                            AVInt  `json:"currentDebt"`
	ShortTermDebt                          AVInt  `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             AVInt  `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                AVInt  `json:"capitalLeaseObligations"`
	LongTermDebt                           AVInt  `json:"longTermDebt"`
	CurrentLongTermDebt                    AVInt  `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 AVInt  `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 AVInt  `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                AVInt  `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             AVInt  `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 AVInt  `json:"totalShareholderEquity"`
	TreasuryStock                          AVInt  `json:"treasuryStock"`
	RetainedEarnings                       AVInt  `json:"retainedEarnings"`
	CommonStock                            AVInt  `json:"commonStock"`
	CommonStockSharesOutstanding           AVInt  `json:"commonStockSharesOutstanding"`
}
type BalanceSheetQuarterlyReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            AVInt  `json:"totalAssets"`
	TotalCurrentAssets                     AVInt  `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  AVInt  `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            AVInt  `json:"cashAndShortTermInvestments"`
	Inventory                              AVInt  `json:"inventory"`
	CurrentNetReceivables                  AVInt  `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  AVInt  `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 AVInt  `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE AVInt  `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       AVInt  `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      AVInt  `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               AVInt  `json:"goodwill"`
	Investments                            AVInt  `json:"investments"`
	LongTermInvestments                    AVInt  `json:"longTermInvestments"`
	ShortTermInvestments                   AVInt  `json:"shortTermInvestments"`
	OtherCurrentAssets                     AVInt  `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  AVInt  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       AVInt  `json:"totalLiabilities"`
	TotalCurrentLiabilities                AVInt  `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 AVInt  `json:"currentAccountsPayable"`
	DeferredRevenue                        AVInt  `json:"deferredRevenue"`
	CurrentDebt                            AVInt  `json:"currentDebt"`
	ShortTermDebt                          AVInt  `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             AVInt  `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                AVInt  `json:"capitalLeaseObligations"`
	LongTermDebt                           AVInt  `json:"longTermDebt"`
	CurrentLongTermDebt                    AVInt  `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 AVInt  `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 AVInt  `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                AVInt  `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             AVInt  `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 AVInt  `json:"totalShareholderEquity"`
	TreasuryStock                          AVInt  `json:"treasuryStock"`
	RetainedEarnings                       AVInt  `json:"retainedEarnings"`
	CommonStock                            AVInt  `json:"commonStock"`
	CommonStockSharesOutstanding           AVInt  `json:"commonStockSharesOutstanding"`
}

type CashFlowStatements struct {
	Symbol           string                              `json:"symbol"`
	AnnualReports    []CashFlowStatementsAnnualReport    `json:"annualReports"`
	QuarterlyReports []CashFlowStatementsQuarterlyReport `json:"quarterlyReports"`
}
type CashFlowStatementsAnnualReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         AVInt  `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            AVInt  `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           AVInt  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              AVInt  `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   AVInt  `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      AVInt  `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       AVInt  `json:"capitalExpenditures"`
	ChangeInReceivables                                       AVInt  `json:"changeInReceivables"`
	ChangeInInventory                                         AVInt  `json:"changeInInventory"`
	ProfitLoss                                                AVInt  `json:"profitLoss"`
	CashflowFromInvestment                                    AVInt  `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     AVInt  `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     AVInt  `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        AVInt  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             AVInt  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     AVInt  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            AVInt  `json:"dividendPayout"`
	DividendPayoutCommonStock                                 AVInt  `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              AVInt  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         AVInt  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet AVInt  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      AVInt  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            AVInt  `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           AVInt  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            AVInt  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      AVInt  `json:"changeInExchangeRate"`
	NetIncome                                                 AVInt  `json:"netIncome"`
}
type CashFlowStatementsQuarterlyReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         AVInt  `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            AVInt  `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           AVInt  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              AVInt  `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   AVInt  `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      AVInt  `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       AVInt  `json:"capitalExpenditures"`
	ChangeInReceivables                                       AVInt  `json:"changeInReceivables"`
	ChangeInInventory                                         AVInt  `json:"changeInInventory"`
	ProfitLoss                                                AVInt  `json:"profitLoss"`
	CashflowFromInvestment                                    AVInt  `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     AVInt  `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     AVInt  `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        AVInt  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             AVInt  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     AVInt  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            AVInt  `json:"dividendPayout"`
	DividendPayoutCommonStock                                 AVInt  `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              AVInt  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         AVInt  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet AVInt  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      AVInt  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            AVInt  `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           AVInt  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            AVInt  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      AVInt  `json:"changeInExchangeRate"`
	NetIncome                                                 AVInt  `json:"netIncome"`
}
