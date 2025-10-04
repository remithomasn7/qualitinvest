-- Revert qualitinvest:qualitinvestschema from pg

BEGIN;

DROP SCHEMA IF EXISTS qualitinvest CASCADE;

COMMIT;
