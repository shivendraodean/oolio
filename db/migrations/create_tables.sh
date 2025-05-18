#!/bin/bash

set -e

POSTGRES_CONTAINER="kart-challenge-postgres-1"
DB_USER="postgres"
DB_NAME="oolio"

echo "Creating tables..."

# Enable UUID extension if not already enabled
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "
CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

# Create orders table
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    items JSONB NOT NULL,
    products JSONB NOT NULL,
    coupon_code TEXT
);"

# Create products table
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    category TEXT NOT NULL,
    image JSONB NOT NULL
);"

# Create couponcodes table
docker exec -i $POSTGRES_CONTAINER psql -U $DB_USER -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS couponcodes (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL
);"

echo "Tables created successfully"
