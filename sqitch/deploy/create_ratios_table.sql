-- Deploy create_ratios_table

BEGIN;

CREATE TABLE financial_ratios (
  id SERIAL PRIMARY KEY,
  company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  fiscal_date DATE NOT NULL,
  roce NUMERIC,                    -- Return on Capital Employed
  gross_margin NUMERIC,
  operating_margin NUMERIC,
  net_margin NUMERIC,
  revenue_growth NUMERIC,
  eps_growth NUMERIC,
  free_cashflow_ratio NUMERIC,
  share_dilution NUMERIC,
  CONSTRAINT uq_financial_ratios UNIQUE(company_id, fiscal_date)
);

CREATE INDEX idx_financial_ratios_company ON financial_ratios(company_id, fiscal_date);

COMMIT;
