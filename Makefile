# Makefile

# Targets
.PHONY: run
run:
	@echo "Running the Go application with air..."
	air

.PHONY: build
build:
	@echo "Building the Go application..."
	go build -o ./tmp/main ./cmd/http/*.*

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f main

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  run    - Run the Go application with air"
	@echo "  build  - Build the Go application"
	@echo "  clean  - Clean up the built files"
	@echo "  test   - Run tests"
