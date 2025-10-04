CREATE USER qualitinvestuser WITH PASSWORD 'qualitinvest_password';
GRANT CONNECT ON DATABASE qualitinvest_db TO qualitinvestuser;
GRANT USAGE ON SCHEMA qualitinvest TO qualitinvestuser;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA qualitinvest TO qualitinvestuser;