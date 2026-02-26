.PHONY: all build run test clean backend frontend dev

# Default target
all: backend frontend

# Backend (Go)
backend:
	cd backend && go build -buildvcs=false -o bin/server ./cmd/server

run-backend: backend
	cd backend && ./bin/server

# Frontend (Vue 3)
frontend:
	cd frontend && npm run build

dev-frontend:
	cd frontend && npm run format && npm run dev

# Run both (backend in background, frontend dev server)
dev: run-backend &
	cd frontend && npm run dev

# Test
test:
	cd backend && go test ./...
	cd frontend && npm run test 2>/dev/null || true

# Format
fmt:
	cd backend && gofmt -s -w .
	cd frontend && npm run format 2>/dev/null || npx prettier --write "src/**/*.{ts,vue,js}" 2>/dev/null || true

# Clean
clean:
	rm -rf backend/bin backend/vendor
	rm -rf frontend/dist frontend/node_modules/.vite

# Install dependencies
deps-backend:
	cd backend && go mod tidy

deps-frontend:
	cd frontend && (pnpm install 2>/dev/null || npm install)
