#!/bin/bash

set -e

POSTGRES_CONTAINER="kart-challenge-postgres-1"
DB_USER="postgres"
DB_NAME="oolio"
TABLE_NAME="couponcodes"

echo "Creating $TABLE_NAME table..."
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS $TABLE_NAME (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL
);"

for file in couponbase1 couponbase2 couponbase3; do
    if [ -f "$file" ]; then
        echo "Importing $file"
        cat "$file" | docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "COPY $TABLE_NAME(code) FROM STDIN;"
        echo "Finished importing $file"
    else
        echo "Warning: $file not found"
    fi
done

echo "Counting imported coupon codes..."
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "SELECT COUNT(*) AS total_coupons FROM $TABLE_NAME;"

echo "Import completed successfully"