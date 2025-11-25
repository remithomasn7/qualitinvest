-- Deploy qualitinvest:create_core_tables to pg
-- Comprehensive schema for value investing application

BEGIN;

-- Core company reference
CREATE TABLE companies (
  id SERIAL PRIMARY KEY,
  symbol TEXT NOT NULL UNIQUE,
  name TEXT,
  exchange TEXT,
  sector TEXT,        -- From overview
  industry TEXT,      -- From overview
  country TEXT,       -- From overview
  currency TEXT       -- From overview
);

-- ============================
-- EARNINGS - Essential for value analysis
-- ============================
CREATE TABLE earnings_annual (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  reported_eps NUMERIC,
  CONSTRAINT uq_earnings_annual UNIQUE(company_id, fiscal_date)
);

CREATE TABLE earnings_quarterly (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  reported_date DATE,
  reported_eps NUMERIC,
  estimated_eps NUMERIC,      -- CRITICAL for value analysis
  surprise NUMERIC,           -- CRITICAL for earnings quality
  surprise_percentage NUMERIC,-- CRITICAL for earnings quality
  report_time TEXT,
  CONSTRAINT uq_earnings_quarterly UNIQUE(company_id, fiscal_date)
);

-- ============================
-- INCOME STATEMENTS - Comprehensive but focused
-- ============================
CREATE TABLE income_annual (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  reported_currency TEXT,

  -- Revenue metrics
  total_revenue NUMERIC,
  cost_of_revenue NUMERIC,
  gross_profit NUMERIC,

  -- Operating metrics (key for margins)
  operating_income NUMERIC,
  selling_general_and_administrative NUMERIC,
  research_and_development NUMERIC,

  -- Bottom line
  net_income NUMERIC,
  ebit NUMERIC,
  ebitda NUMERIC,

  -- Tax and interest (important for analysis)
  income_tax_expense NUMERIC,
  interest_expense NUMERIC,

  CONSTRAINT uq_income_annual UNIQUE(company_id, fiscal_date)
);

CREATE TABLE income_quarterly (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  reported_currency TEXT,

  -- Revenue metrics
  total_revenue NUMERIC,
  cost_of_revenue NUMERIC,
  gross_profit NUMERIC,

  -- Operating metrics (key for margins)
  operating_income NUMERIC,
  selling_general_and_administrative NUMERIC,
  research_and_development NUMERIC,

  -- Bottom line
  net_income NUMERIC,
  ebit NUMERIC,
  ebitda NUMERIC,

  -- Tax and interest (important for analysis)
  income_tax_expense NUMERIC,
  interest_expense NUMERIC,

  CONSTRAINT uq_income_quarterly UNIQUE(company_id, fiscal_date)
);

-- ============================
-- BALANCE SHEETS - Key balance sheet items
-- ============================
CREATE TABLE balance_annual (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,

  -- Assets
  total_assets NUMERIC,
  total_current_assets NUMERIC,
  cash_and_cash_equivalents_at_carrying_value NUMERIC,
  inventory NUMERIC,
  current_net_receivables NUMERIC,

  -- Liabilities
  total_liabilities NUMERIC,
  total_current_liabilities NUMERIC,
  current_accounts_payable NUMERIC,
  short_term_debt NUMERIC,
  long_term_debt NUMERIC,

  -- Equity (critical for book value)
  total_shareholder_equity NUMERIC,
  retained_earnings NUMERIC,
  common_stock NUMERIC,

  -- Outstanding shares
  common_stock_shares_outstanding BIGINT,

  CONSTRAINT uq_balance_annual UNIQUE(company_id, fiscal_date)
);

CREATE TABLE balance_quarterly (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,

  -- Assets
  total_assets NUMERIC,
  total_current_assets NUMERIC,
  cash_and_cash_equivalents_at_carrying_value NUMERIC,
  inventory NUMERIC,
  current_net_receivables NUMERIC,

  -- Liabilities
  total_liabilities NUMERIC,
  total_current_liabilities NUMERIC,
  current_accounts_payable NUMERIC,
  short_term_debt NUMERIC,
  long_term_debt NUMERIC,

  -- Equity (critical for book value)
  total_shareholder_equity NUMERIC,
  retained_earnings NUMERIC,
  common_stock NUMERIC,

  -- Outstanding shares
  common_stock_shares_outstanding BIGINT,

  CONSTRAINT uq_balance_quarterly UNIQUE(company_id, fiscal_date)
);

-- ============================
-- CASH FLOW - Essential cash flow analysis
-- ============================
CREATE TABLE cashflow_annual (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,

  -- Operating cash flow (most important)
  operating_cashflow NUMERIC,

  -- Investing activities
  capital_expenditures NUMERIC,
  cashflow_from_investment NUMERIC,

  -- Financing activities
  cashflow_from_financing NUMERIC,
  dividend_payout NUMERIC,

  -- Net change
  change_in_cash_and_cash_equivalents NUMERIC,

  -- Net income for CFO reconciliation
  net_income NUMERIC,

  CONSTRAINT uq_cashflow_annual UNIQUE(company_id, fiscal_date)
);

CREATE TABLE cashflow_quarterly (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,

  -- Operating cash flow (most important)
  operating_cashflow NUMERIC,

  -- Investing activities
  capital_expenditures NUMERIC,
  cashflow_from_investment NUMERIC,

  -- Financing activities
  cashflow_from_financing NUMERIC,
  dividend_payout NUMERIC,

  -- Net change
  change_in_cash_and_cash_equivalents NUMERIC,

  -- Net income for CFO reconciliation
  net_income NUMERIC,

  CONSTRAINT uq_cashflow_quarterly UNIQUE(company_id, fiscal_date)
);

-- ============================
-- SHARES OUTSTANDING
-- ============================
CREATE TABLE shares_outstanding_annual (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  basic BIGINT,
  diluted BIGINT,
  CONSTRAINT uq_shares_out_annual UNIQUE(company_id, fiscal_date)
);

CREATE TABLE shares_outstanding_quarterly (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  basic BIGINT,
  diluted BIGINT,
  CONSTRAINT uq_shares_out_quarterly UNIQUE(company_id, fiscal_date)
);

-- ============================
-- SPLITS
-- ============================
CREATE TABLE splits (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  effective_date DATE NOT NULL,
  split_factor TEXT NOT NULL,
  CONSTRAINT uq_splits UNIQUE(company_id, effective_date)
);

-- ============================
-- DIVIDENDS
-- ============================
CREATE TABLE dividends (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  ex_date DATE,
  declaration_date DATE,
  record_date DATE,
  payment_date DATE,
  amount NUMERIC,
  CONSTRAINT uq_dividends UNIQUE(company_id, ex_date)
);

-- ============================
-- COMPANY OVERVIEW - Key valuation metrics
-- ============================
CREATE TABLE company_overview (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  symbol TEXT NOT NULL,

  -- Valuation ratios (CRITICAL for value investing)
  pe_ratio DECIMAL(10,4),
  peg_ratio DECIMAL(10,4),
  pb_ratio DECIMAL(10,4),  -- Price to Book - essential for value
  price_to_sales_ratio_ttm DECIMAL(10,4),

  -- Profitability
  profit_margin DECIMAL(10,4),
  operating_margin_ttm DECIMAL(10,4),
  return_on_assets_ttm DECIMAL(10,4),
  return_on_equity_ttm DECIMAL(10,4),

  -- Growth
  quarterly_earnings_growth_yoy DECIMAL(10,4),
  quarterly_revenue_growth_yoy DECIMAL(10,4),

  -- Market data
  market_capitalization BIGINT,
  beta DECIMAL(10,4),

  -- Size metrics
  shares_outstanding BIGINT,

  -- Update tracking
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

  CONSTRAINT uq_company_overview_symbol UNIQUE(symbol)
);

-- ============================
-- HISTORICAL PRICES - For technical analysis
-- ============================
CREATE TABLE price_history (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  symbol TEXT NOT NULL,
  date DATE NOT NULL,

  -- OHLC data
  open DECIMAL(10,4),
  high DECIMAL(10,4),
  low DECIMAL(10,4),
  close DECIMAL(10,4),
  adjusted_close DECIMAL(10,4),
  volume BIGINT,

  -- Key technical indicators
  sma_50 DECIMAL(10,4),  -- 50-day moving average
  sma_200 DECIMAL(10,4), -- 200-day moving average

  CONSTRAINT uq_price_history_symbol_date UNIQUE(company_id, date)
);

-- Indexes for performance
CREATE INDEX idx_earnings_annual_company_date ON earnings_annual(company_id, fiscal_date DESC);
CREATE INDEX idx_earnings_quarterly_company_date ON earnings_quarterly(company_id, fiscal_date DESC);
CREATE INDEX idx_income_annual_company_date ON income_annual(company_id, fiscal_date DESC);
CREATE INDEX idx_income_quarterly_company_date ON income_quarterly(company_id, fiscal_date DESC);
CREATE INDEX idx_balance_annual_company_date ON balance_annual(company_id, fiscal_date DESC);
CREATE INDEX idx_balance_quarterly_company_date ON balance_quarterly(company_id, fiscal_date DESC);
CREATE INDEX idx_cashflow_annual_company_date ON cashflow_annual(company_id, fiscal_date DESC);
CREATE INDEX idx_cashflow_quarterly_company_date ON cashflow_quarterly(company_id, fiscal_date DESC);
CREATE INDEX idx_shares_out_annual_company_date ON shares_outstanding_annual(company_id, fiscal_date DESC);
CREATE INDEX idx_shares_out_quarterly_company_date ON shares_outstanding_quarterly(company_id, fiscal_date DESC);
CREATE INDEX idx_splits_company_date ON splits(company_id, effective_date DESC);
CREATE INDEX idx_dividends_company_date ON dividends(company_id, ex_date DESC);
CREATE INDEX idx_price_history_company_date ON price_history(company_id, date DESC);
CREATE INDEX idx_company_overview_pe_ratio ON company_overview(pe_ratio);
CREATE INDEX idx_company_overview_pb_ratio ON company_overview(pb_ratio);
CREATE INDEX idx_company_overview_market_cap ON company_overview(market_capitalization);

COMMIT;
