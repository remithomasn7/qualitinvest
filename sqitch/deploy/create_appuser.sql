-- Create application user (same as docker-compose user for consistency)
DO $$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'rthomas') THEN
      CREATE USER rthomas WITH PASSWORD 'rthomas';
   END IF;
END
$$;

GRANT CONNECT ON DATABASE qualitinvest_db TO rthomas;
GRANT USAGE ON SCHEMA qualitinvest TO rthomas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA qualitinvest TO rthomas;

-- Also create qualitinvestuser as alternative (for future multi-user support)
DO $$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'qualitinvestuser') THEN
      CREATE USER qualitinvestuser WITH PASSWORD 'qualitinvest_password';
   END IF;
END
$$;

GRANT CONNECT ON DATABASE qualitinvest_db TO qualitinvestuser;
GRANT USAGE ON SCHEMA qualitinvest TO qualitinvestuser;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA qualitinvest TO qualitinvestuser;