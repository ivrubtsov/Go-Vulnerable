.PHONY: help scan scan-json build clean update-deps verify

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

scan: ## Run Trivy vulnerability scan
	@echo "Scanning for vulnerabilities..."
	trivy fs --severity HIGH,CRITICAL .

scan-json: ## Run Trivy scan and output JSON
	@echo "Scanning and generating JSON report..."
	trivy fs --format json --output trivy-results.json .
	@echo "Results saved to trivy-results.json"

scan-all: ## Run comprehensive Trivy scan
	@echo "Running comprehensive vulnerability scan..."
	trivy fs --severity LOW,MEDIUM,HIGH,CRITICAL .

build: ## Build the application
	@echo "Building vulnerable-app..."
	go build -o vulnerable-app main.go

run: build ## Build and run the application
	@echo "Starting vulnerable application on :8080"
	@echo "WARNING: This application has known vulnerabilities!"
	./vulnerable-app

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f vulnerable-app
	rm -f trivy-results.json
	rm -f trivy-report.html

update-deps: ## Update dependencies (for AI agent testing)
	@echo "This target is for your AI agent to implement"
	@echo "It should update go.mod with secure dependency versions"

verify: ## Verify the application builds after updates
	@echo "Verifying application..."
	go mod verify
	go build -o vulnerable-app main.go
	@echo "Build successful!"

test: ## Run tests (placeholder for actual tests)
	@echo "Running tests..."
	go test -v ./...
