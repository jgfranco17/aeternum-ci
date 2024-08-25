# Default command
default:
    @just --list

# Run debug server
run-local port:
    go run ./api/cmd/main.go --port={{port}} --dev=true

# Run production server
run-prod port:
    go run ./api/cmd/main.go --port {{port}} --dev=false

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover ./api/...

# Build Docker image
build tag="latest":
	@echo "Building Docker image (tag={{ tag }})..."
	docker build -t aeternum:{{ tag }} -f ./docker/server.Dockerfile .
	@echo "Docker image built successfully!"

# Sync Go modules
tidy:
    cd api && go mod tidy
    go work sync

# Start Compose with load-balancer
compose-up:
    docker compose -f docker/docker-compose.yml up

# Stop all Compose containers and delete images created
compose-down:
    docker compose -f docker/docker-compose.yml down
    docker rmi $(docker images | grep "aeternum" | awk "{print \$3}")
