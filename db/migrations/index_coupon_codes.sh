#!/bin/bash

set -e

POSTGRES_CONTAINER="kart-challenge-postgres-1"
DB_USER="postgres"
DB_NAME="oolio"
TABLE_NAME="couponcodes"

echo "Vacuuming database to reclaim space..."
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "VACUUM FULL;"

echo "Creating index on $TABLE_NAME.code"
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "CREATE INDEX IF NOT EXISTS idx_${TABLE_NAME}_code ON $TABLE_NAME(code) WITH (fillfactor=70);"

echo "DB Indexed successfully."
