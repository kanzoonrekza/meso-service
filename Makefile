include .env
export

build:
	@go build -o . cmd/main.go

run:
	@go run cmd/main.go

run-bin:
	@./main

air:
	@air cmd/main.go

goose:
	@cmd=$(word 2, $(MAKECMDGOALS)); \
	dir=$(GOOSE_MIGRATION_DIR); \
	driver=$(GOOSE_DRIVER); \
	db_url=$(DB_URL); \
	case $$cmd in \
		create|init) \
			name=$(word 3, $(MAKECMDGOALS)); \
			if [ -z "$$name" ]; then \
				read -p "Migration name: " name; \
			fi; \
			if [ -n "$$name" ]; then \
				echo "💎 Created migration: $$name"; \
				goose create $$name sql -dir $$dir; \
			else \
				echo "❌️ Error: Migration name cannot be empty"; \
				exit 1; \
			fi; \
			;; \
		up|push) \
			echo "🚀 Migrate database to latest..."; \
			goose -dir $$dir $$db_url up; \
			;; \
		down|pull) \
			echo "🪃 Roll back database by one..."; \
			goose -dir $$dir $$db_url down; \
			;; \
		reset) \
			echo "🔄 Reset database..."; \
			goose -dir $$dir $$db_url reset; \
			;; \
		*) \
			echo "❌️ Unknown command: make goose $$cmd"; \
	esac

sqlc:
	@sqlc generate

# Catch undefined second make target
$(wordlist 2, 100, $(MAKECMDGOALS)):
	@: