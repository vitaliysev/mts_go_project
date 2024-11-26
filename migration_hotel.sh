#!/bin/bash
source .env

export MIGRATION_DSN="host=pg port=5432 dbname=$POSTGRES_HOTEL_DB user=$POSTGRES_HOTEL_USER password=$POSTGRES_HOTEL_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_HOTEL_DIR}" postgres "${MIGRATION_DSN}" up -v