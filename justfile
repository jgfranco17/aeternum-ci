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
build:
	@echo "Building Docker image..."
	docker build -t aeturnum:latest -f .docker/api.Dockerfile .
	@echo "Docker image built successfully!"

# Sync Go modules
tidy:
    cd api && go mod tidy
    go work sync

# Start up Docker Compose server
docker-up:
    @echo "Starting up full Docker suite..."
    docker compose up --build

# Start up Docker Compose server in detached mode
docker-detached:
    @echo "Starting up full Docker suite..."
    docker compose up -d

# Clean up Docker container remnants
docker-down:
    docker compose down
    @echo "Compose server closed successfully!"
