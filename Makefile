models:
	cd scripts && sqlboiler postgres --wipe --output ../db/models

mocks: mock-db mock-exchange

mock-db:
	mockgen -source=./db/db.go -destination=./db/mock-db.go -package=db

mock-exchange:
	mockgen -source=./exchanges/exchanges.go -destination=./exchanges/mock-exchange.go -package=exchanges

cli:
	go build -o bin/cli cmd/cli/*.go

cli-production:
	GOOS=linux GOARCH=amd64 go build -o bin/cli cmd/cli/*.go

setup-db: create-db
	sql-migrate up -config=scripts/dbconfig.yml

teardown-db:
	sql-migrate down -config=scripts/dbconfig.yml

create-db:
	psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'coins'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE coins"

setup: create-db setup-db cli

migrate-production:
	sql-migrate up -config=scripts/dbconfig.yml -env=production
