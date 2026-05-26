#!/usr/bin/env bash
set -e

psql -v ON_ERROR_STOP=1 --username "apiuser" --dbname "api" <<-EOSQL
  CREATE TABLE IF NOT EXISTS user_input (
  id SERIAL PRIMARY KEY,
  text VARCHAR(50)
  );
EOSQL
