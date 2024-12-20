include .env
export

build:
	@go build -o internal/main.go

run:
	@go run main

air:
	@air internal/main.go

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
		*) \
			echo "❌️ Unknown command: make goose $$cmd"; \
	esac

# Catch undefined second make target
$(wordlist 2, 100, $(MAKECMDGOALS)):
	@: