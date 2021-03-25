#!/bin/bash

MIGRATE_DIRECTION=$1

cd migrations

# Run the migration for both the application db
# and the test db

echo "Running application migrations"
goose postgres "user=postgres dbname=raspberry_a_pi sslmode=disable" $MIGRATE_DIRECTION
echo -e "Application migrations complete \n \n"

echo "Running test migrations"
goose postgres "user=postgres dbname=raspberry_a_pi_test sslmode=disable" $MIGRATE_DIRECTION

echo -e "test migrations complete \n \n"

# Return to root dir
cd ..
