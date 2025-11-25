-- Verify create_appuser

BEGIN;

-- Verify that users exist and have proper permissions
SELECT 1 FROM pg_catalog.pg_roles WHERE rolname = 'rthomas'
UNION ALL
SELECT 1 FROM pg_catalog.pg_roles WHERE rolname = 'qualitinvestuser';

-- Verify schema permissions
SELECT 1 FROM information_schema.role_table_grants
WHERE grantee IN ('rthomas', 'qualitinvestuser')
AND table_schema = 'qualitinvest'
AND privilege_type IN ('SELECT', 'INSERT', 'UPDATE', 'DELETE');

COMMIT;
