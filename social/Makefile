run:
	@air # using air for live reloading

test:
	@go test -v ./...

migration:
	@if [ -z "$(title)" ]; then \
		echo "Error: title is required. Usage: make migration title=your_migration_name"; \
		exit 1; \
	fi
	@mkdir -p ./cmd/migrate/migrations
	@timestamp=$$(date +%Y%m%d%H%M%S); \
	touch "./cmd/migrate/migrations/$${timestamp}_$(title).up.sql"; \
	touch "./cmd/migrate/migrations/$${timestamp}_$(title).down.sql"; \
	echo "Created migration files:"; \
	echo "  - $${timestamp}_$(title).up.sql"; \
	echo "  - $${timestamp}_$(title).down.sql"

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
	