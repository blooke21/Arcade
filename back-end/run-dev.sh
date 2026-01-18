#!/usr/bin/env bash
set -e

# free port 8080 if any PIDs are listening
pids="$(lsof -t -i:8080 2>/dev/null || true)"
if [ -n "$pids" ]; then
  echo "killing PIDs on :8080 -> $pids"
  kill $pids || true
fi

echo "starting back-end docker containers..."
cd "$(dirname "$0")"

# If CompileDaemon is installed on the host use it for live reload.
if command -v CompileDaemon >/dev/null 2>&1; then
  echo "Using CompileDaemon for live reload"
  # run in background, capture output
  CompileDaemon -directory=. -command="go run ." > backend.log 2>&1 &
else
  echo "CompileDaemon not found — running 'go run main.go' (no auto-reload)"
  nohup go run . > backend.log 2>&1 &
fi

# wait until backend is listening (timeout after 30s)
echo "waiting for backend..."
start_ts=$(date +%s)
timeout=30
until curl -s http://localhost:8080/api/roms >/dev/null 2>&1; do
  sleep 0.5
  now=$(date +%s)
  if [ $((now - start_ts)) -gt $timeout ]; then
    echo "timeout waiting for backend (>$timeout s). last backend log:"
    tail -n 200 backend.log || true
    exit 1
  fi
done
echo "backend up, launching electron app from host..."

# run frontend start from sibling front-end folder
cd ../front-end
npm start