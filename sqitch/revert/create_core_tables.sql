-- Revert qualitinvest:create_core_tables from pg

BEGIN;

DROP TABLE IF EXISTS dividends CASCADE;
DROP TABLE IF EXISTS splits CASCADE;
DROP TABLE IF EXISTS shares_outstanding_quarterly CASCADE;
DROP TABLE IF EXISTS shares_outstanding_annual CASCADE;
DROP TABLE IF EXISTS cashflow_quarterly CASCADE;
DROP TABLE IF EXISTS cashflow_annual CASCADE;
DROP TABLE IF EXISTS balance_quarterly CASCADE;
DROP TABLE IF EXISTS balance_annual CASCADE;
DROP TABLE IF EXISTS income_quarterly CASCADE;
DROP TABLE IF EXISTS income_annual CASCADE;
DROP TABLE IF EXISTS earnings_quarterly CASCADE;
DROP TABLE IF EXISTS earnings_annual CASCADE;
DROP TABLE IF EXISTS companies CASCADE;

COMMIT;
