#!/bin/bash

# 1. Load .env file
set -o allexport
source .env
set +o allexport

# 2. Generate jet models
jet -dsn="mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path=./app

rm -rf ./app/databases/*

mkdir ./app/databases

mv -f ./app/${DB_NAME}/* ./app/databases
echo "✅ Moved: ./app/${DB_NAME} to ./app/databases"

find . -type f -path "./app/databases/table/*.go" -name "*.go" | while read -r file; do
	echo "📝 Processing: $file"
	sed -i.bak "s/new\([A-Za-z]*\)Table(\"$DB_NAME\",/new\1Table(\"\",/" "$file"
	rm "${file}.bak"
done

echo "✅ Done stripping schema: $DB_NAME"

rm -rf ./app/databases/${DB_NAME}
echo "✅ Removed tmp backup: $DB_NAME"

if [ $? -ne 0 ]; then
	echo "❌ Jet Generate Failed failed."
	exit 1
fi
