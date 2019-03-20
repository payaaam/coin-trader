test:
	go test ./orders ./strategies ./cmd/cli ./charts

models:
	cd scripts && sqlboiler postgres --wipe --output ../db/models

mocks:
	./scripts/generate-mocks.sh

cli:
	go build -o bin/cli cmd/cli/*.go

cli-production:
	GOOS=linux GOARCH=amd64 go build -o bin/cli cmd/cli/*.go

setup-db: create-db
	sql-migrate up -config=scripts/dbconfig.yml

teardown-db:
	sql-migrate down -config=scripts/dbconfig.yml --limit=2

create-db:
	psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'coins'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE coins"

setup: create-db setup-db cli

migrate-production:
	sql-migrate up -config=scripts/dbconfig.yml -env=production

upload-cli:
	scp bin/cli payam@ec2-35-170-55-23.compute-1.amazonaws.com:/home/payam/cli-02-01

deploy:
	./scripts/deploy.sh

.PHONY: mocks