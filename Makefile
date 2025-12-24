.PHONY: all clean docs build build-server build-ui lint lint-server lint-ui test test-server

# Default target
all: lint test build

docs:
	go run github.com/swaggo/swag/cmd/swag@latest init \
		-g main.go \
		--dir ./server/cmd/api,./server/internal/handlers,./server/internal/models/response,./server/internal/server,./server/internal/services \
		--outputTypes json,yaml \
		-o docs

# Build
build: build-server build-ui

build-server:
	cd server && go build -o bin/api cmd/api/main.go

build-ui:
	cd ui && npm run build

# Clean
clean:
	rm -rf server/bin
	rm -rf ui/dist
	rm -rf ui/node_modules

# Lint
lint: lint-server lint-ui

lint-server:
	cd server && go vet ./...

lint-ui:
	cd ui && npm run lint

# Test
test: test-server

test-server:
	cd server && go test ./...
