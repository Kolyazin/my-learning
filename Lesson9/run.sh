#!/bin/bash

export PORT=5555
export DB_URL="postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
export JAEGER_URL="http://env:16686"
export SENTRY_URL="http://env:9000"
export KAFKA_BROKER="env:1234"
export SOME_APP_ID="EnvId"
export SOME_APP_KEY="EnvKey"

echo -e "\033[42m\033[1m> Get default variables\033[0m"
go run main.go -port 8080

echo -e "\033[42m\033[1m> Get variables from environment\033[0m"
go run main.go

echo -e "\033[42m\033[1m> Get variables from flags\033[0m"
go run main.go -port 7777 -db_url "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" \
-jaeger_url "http://flag:16686" -sentry_url "http://flag:9000" -kafka_broker "flag:5678" \
-some_app_id "FlagId" -some_app_key "FlagKey"

echo -e "\033[42m\033[1m> Get variables from conf.YAML\033[0m"
go run main.go conf.yaml

echo -e "\033[42m\033[1m> Get variables from conf.JSON\033[0m"
go run main.go conf.json