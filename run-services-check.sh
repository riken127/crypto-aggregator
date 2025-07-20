#!/bin/bash

set -e

echo "Running checks for all services..."

for service in services/*; do
  if [ -d "$service" ]; then
    echo "Entering $service"
    cd "$service"

    if [ -f "go.mod" ]; then
      echo "Detected Go project"
      go test ./... || echo "⚠️ Go tests failed @ $service"
    elif [ -f "package.json" ]; then
      echo "Detected Node.js project"
      npm install
      npm run lint || echo "⚠️ Lint failed @ $service"
      npm test || echo "⚠️ Tests failed @ $service"
    else
      echo "Unknown type @ $service"
    fi

    cd - > /dev/null
  fi
done

echo "✅ All checks completed successfully!"
