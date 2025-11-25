-- Revert create_appuser

BEGIN;

-- Revoke permissions and drop users (only drop qualitinvestuser, keep rthomas as it's the main user)
REVOKE ALL PRIVILEGES ON SCHEMA qualitinvest FROM qualitinvestuser;
REVOKE CONNECT ON DATABASE qualitinvest_db FROM qualitinvestuser;

-- Drop the user if it exists
DROP USER IF EXISTS qualitinvestuser;

COMMIT;
