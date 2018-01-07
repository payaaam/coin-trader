models:
	cd scripts && sqlboiler postgres --wipe --output ../db/models

cli:
	go build -o bin/cli cmd/cli/*.go