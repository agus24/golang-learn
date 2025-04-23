#!/bin/bash

# 1. Load .env file
set -o allexport
source .env
set +o allexport

# 2. Generate jet models
migrate -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path db/migrations force "$1"

if [ $? -ne 0 ]; then
	echo "❌ Migration failed."
	exit 1
fi

echo "✅ Migration completed successfully."
