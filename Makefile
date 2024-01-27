run:
	go run main.go

migrate-up:
	cd migrations/schema && goose postgres postgres://postgres:pass@localhost:5432/rssagg up

migrate-down:
	cd migrations/schema && goose postgres postgres://postgres:pass@localhost:5432/rssagg down

gen:
	sqlc generate