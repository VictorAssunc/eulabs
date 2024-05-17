deps:
	@docker compose up -d --force-recreate

down:
	@docker compose down

run: deps
	@sleep 2
	@go run ./cmd/api.go