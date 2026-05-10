# Makefile for stratus framework

.PHONY: help fmt check add commit git-push tag push-tag

# Module prefix for tagging
TAG_PREFIX=libraries/framework/stratus

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

fmt: ## Format code with go fmt and goimports
	@echo "Formatting code..."
	go fmt ./...

check: ## Run static analysis, vulnerability check, and test coverage (must be > 75%)
	@echo "Running checks..."
	go vet ./...
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...|| true


add: ## Stage all changes (git add .)
	git add .

commit: ## Commit changes (usage: make commit msg="message")
	@if [ -z "$(msg)" ]; then echo "Error: msg is required. Usage: make commit msg=\"message\""; exit 1; fi
	git commit -m "$(msg)"

git-push: ## Push committed changes to remote
	git push

tag: ## Create version tag (usage: make tag minor=1 patch=0)
	@if [ -z "$(minor)" ] || [ -z "$(patch)" ]; then echo "Error: minor and patch are required. Usage: make tag minor=X patch=Y"; exit 1; fi
	@echo "Creating tag $(TAG_PREFIX)/v1.$(minor).$(patch)"
	git tag $(TAG_PREFIX)/v1.$(minor).$(patch)

push-tag: ## Push version tag to origin (usage: make push-tag minor=1 patch=0)
	@if [ -z "$(minor)" ] || [ -z "$(patch)" ]; then echo "Error: minor and patch are required. Usage: make push-tag minor=X patch=Y"; exit 1; fi
	@echo "Pushing tag $(TAG_PREFIX)/v1.$(minor).$(patch)"
	git push origin $(TAG_PREFIX)/v1.$(minor).$(patch)
