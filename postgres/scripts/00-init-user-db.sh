#!/usr/bin/env bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER apiuser;
	CREATE DATABASE api;
	GRANT ALL PRIVILEGES ON DATABASE api TO apiuser;

  \c api
  GRANT ALL PRIVILEGES ON SCHEMA public TO apiuser;
EOSQL
