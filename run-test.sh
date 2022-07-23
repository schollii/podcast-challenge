#!/usr/bin/env bash

set -eu

# ATTENTION: script assumes $USER is in docker group

echo "Starting services..."
docker-compose up -d --force-recreate --build --scale podcast=1
while ! curl -Is http://localhost:8081/health; do sleep 1; done

echo "-----------------------------------"
echo "Running test"
cd integ-test
go test -v
TEST_RESULT=$?
cd -

echo "-----------------------------------"
echo "Test done, tearing down deployment"

docker-compose down
exit $TEST_RESULT