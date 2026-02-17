.PHONY: test-backend test-frontend lint-backend lint-frontend coverage-backend

test-backend:
	cd backend && go test -v -coverprofile=coverage.out ./...
	cd backend && go tool cover -func=coverage.out

# コアロジック(domain, usecase)のカバレッジが100%であることを検証するスクリプト
coverage-backend: test-backend
	@echo "Checking core logic coverage..."
	@cd backend && go tool cover -func=coverage.out | grep -E "internal/domain|internal/usecase" | grep -v "_mock.go" | awk '{if ($$NF != "100.0%") { print "Low coverage in core: " $$0; exit 1 }}'
	@echo "Core logic coverage is 100%!"

test-frontend:
	cd frontend && npm test

lint-backend:
	docker compose run --rm backend golangci-lint run ./...

lint-frontend:
	cd frontend && npm run lint

# 全てのローカルチェックを実行
check-all: lint-backend lint-frontend test-backend test-frontend
