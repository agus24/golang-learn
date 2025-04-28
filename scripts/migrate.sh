#!/bin/bash

# 1. Load .env file
set -o allexport
source .env
set +o allexport

TYPE="$1"
ORDER="$2"

# 2. Generate jet models
if [ -n "$ORDER" ]; then
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path db/migrations "$TYPE" "$ORDER"
else
	migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path db/migrations "$TYPE"
fi

if [ $? -ne 0 ]; then
	echo "❌ Migration failed."
	exit 1
fi

echo "✅ Migration completed successfully."
