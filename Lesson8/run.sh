#!/bin/bash

export PORT=5555
export DB_URL="postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
export JAEGER_URL="http://jaeger:16686"
export SENTRY_URL="http://sentry:9000"
export KAFKA_BROKER="rabbitmq:1234"
export SOME_APP_ID="EnvId"
export SOME_APP_KEY="EnvKey"

echo -e "\033[42m\033[1m> Get default variables\033[0m"
go run main.go -port 8080

echo -e "\033[42m\033[1m> Get variables from environment\033[0m"
go run main.go

echo -e "\033[42m\033[1m> Get variables from flags\033[0m"
go run main.go -port 7777 -db_url "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" \
-jaeger_url "http://jaeger:16686" -sentry_url "http://sentry:9000" -kafka_broker "zeromq:5678" \
-some_app_id "FlagId" -some_app_key "FlagKey"