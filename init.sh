#!/bin/bash
set -e
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	  CREATE DATABASE users;
    \connect users;
    create schema V1;
    set search_path TO V1;
    CREATE TABLE users (
        id BIGSERIAL NOT NULL,
        first_name VARCHAR(30) NOT NULL,
        last_name VARCHAR(30) NOT NULL,
        city VARCHAR(10) NOT NULL,
        created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
        PRIMARY KEY (id)
    );
    COMMIT;
EOSQL