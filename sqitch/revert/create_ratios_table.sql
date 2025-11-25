-- Revert create_ratios_table

BEGIN;

DROP TABLE IF EXISTS financial_ratios;

COMMIT;
