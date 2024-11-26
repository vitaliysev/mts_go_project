#!/bin/bash
source .env

export MIGRATION_DSN="host=pg-local port=5432 dbname=$POSTGRES_BOOKING_DB user=$POSTGRES_BOOKING_USER password=$POSTGRES_BOOKING_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_BOOKING_DIR}" postgres "${MIGRATION_DSN}" up -v