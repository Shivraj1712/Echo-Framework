.PHONY: run-server tidy migrate-up migrate-down
run-server:
	cd server && air
tidy:
	cd server && go mod tidy
migrate-up:
	cd server && go run ./cmd/migrate-up/main.go
migrate-down:
	cd server && go run ./cmd/migration-down/main.go
