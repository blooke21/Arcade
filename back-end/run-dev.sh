#!/usr/bin/env bash
set -e

# start containers in background
docker compose up -d --build

# wait until backend is healthy or listening (simple wait)
echo "waiting for backend..."
until curl -s http://localhost:8080/api/roms >/dev/null 2>&1; do
  sleep 0.5
done
echo "backend up, launching electron app from host..."

# run frontend start from sibling front-end folder
cd ../front-end
npm start