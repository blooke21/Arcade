#!/usr/bin/env bash
set -e

echo "Cleaning up ports..."
# Kill anything on ports 8080 and 3000
lsof -t -i:8080 -i:3000 2>/dev/null | xargs kill 2>/dev/null || true

# Start docker compose in background (detached mode)
echo "Starting backend with Docker..."
docker compose up --build -d

# Wait until backend is listening (timeout after 30s)
echo "Waiting for backend..."
start_ts=$(date +%s)
timeout=30
until curl -s http://localhost:8080/api/roms >/dev/null 2>&1; do
    sleep 0.5
    now=$(date +%s)
    elapsed=$((now - start_ts))
    if [ $elapsed -gt $timeout ]; then
        echo "Timeout waiting for backend (>${timeout}s)."
        echo "Backend logs:"
        docker compose logs backend
        exit 1
    fi
done

echo "Backend up! Launching Electron app..."

# Run frontend from sibling front-end folder
cd ../front-end
npm start

# When Electron exits, stop docker containers
echo ""
echo "Electron closed. Stopping Docker containers..."
cd ../back-end
docker compose down