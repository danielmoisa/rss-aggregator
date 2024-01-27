run:
	go run main.go

migrate-up:
	cd sql/schema && goose postgres postgres://postgres:pass@localhost:5432/rssagg up

migrate-down:
	cd sql/schema && goose postgres postgres://postgres:pass@localhost:5432/rssagg down

gen:
	sqlc generate