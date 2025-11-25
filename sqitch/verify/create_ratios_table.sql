-- Verify create_ratios_table

SELECT
  id,
  company_id,
  fiscal_date,
  roce,
  gross_margin,
  operating_margin,
  net_margin,
  revenue_growth,
  eps_growth,
  free_cashflow_ratio,
  share_dilution
FROM financial_ratios
WHERE 1 = 0;
