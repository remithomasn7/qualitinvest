-- Verify qualitinvest:create_core_tables on pg

BEGIN;

SELECT 1 FROM pg_tables WHERE tablename = 'companies';
SELECT 1 FROM pg_tables WHERE tablename = 'earnings_annual';
SELECT 1 FROM pg_tables WHERE tablename = 'income_annual';
SELECT 1 FROM pg_tables WHERE tablename = 'balance_annual';
SELECT 1 FROM pg_tables WHERE tablename = 'cashflow_annual';
SELECT 1 FROM pg_tables WHERE tablename = 'dividends';

COMMIT;
