#!/bin/bash

# 1. Load .env file
set -o allexport
source .env
set +o allexport

# 2. Generate jet models
jet -dsn="mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -schema=dvds -path=./app

if [ $? -ne 0 ]; then
	echo "❌ Jet Generate Failed failed."
	exit 1
fi
