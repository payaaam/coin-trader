models:
	cd scripts && sqlboiler postgres --wipe --output ../db/models

mocks: mock-db mock-exchange mock-manager mock-strategy

mock-db:
	mockgen -source=./db/db.go -destination=./mocks/db.go -package=mocks

mock-exchange:
	mockgen -source=./exchanges/exchanges.go -destination=./mocks/exchanges.go -package=mocks -imports .=github.com/payaaam/coin-trader/exchanges

mock-manager: 
	mockgen -source=./orders/orders.go -destination=./orders/orders_mock.go -package=orders

mock-strategy: 
	mockgen -source=./strategies/strategy.go -destination=./mocks/strategy.go -package=mocks

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