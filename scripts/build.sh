#!/bin/bash

# 1. Load .env file
set -o allexport
source .env
set +o allexport

VERSION="${1:-latest}"

mkdir -p .build
go build -o .build/"$APP_NAME"_"$VERSION" .
