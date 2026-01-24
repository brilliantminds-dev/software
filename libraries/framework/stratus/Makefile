# Variables
BINARY_NAME=myapp
VERSION_FILE=VERSION
# Reads version from file; defaults to v0.0.0 if file doesn't exist
VERSION=$(shell cat $(VERSION_FILE) 2>/dev/null || echo "v0.0.0")


fmt:
	@echo "Formatting Go code..."
	go fmt ./...

## Help: Show available commands
help:
	@echo "Usage: make [target]"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

## Testing & Security
test: ## Run unit tests
	go test ./...

test-verbose: ## Run unit tests with verbose output and coverage
	go test -v -cover ./...

vuln: ## Check for vulnerabilities in dependencies
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

## Git Automation
add: ## Git add all changes
	git add .

commit: ## Git commit (usage: make commit msg="message")
	git commit -m "$(msg)"

tag: ## Create and push a git tag (usage: make tag v=v1.0.1)
	git tag -a $(v) -m "Release $(v)"
	git push origin $(v)

bump: ## Increment patch version (v0.0.1 -> v0.0.2) and update VERSION file
	@echo $(VERSION) | awk -F. '{print $$1"."$$2"."$$3+1}' > $(VERSION_FILE)
	@echo "Version bumped to $$(cat $(VERSION_FILE))"

release-auto: bump ## Add, Commit, and Tag using the auto-incremented version
	git add .
	git commit -m "release: $$(cat $(VERSION_FILE))"
	git tag -a $$(cat $(VERSION_FILE)) -m "Release $$(cat $(VERSION_FILE))"
	git push origin main --tags

.PHONY: help test test-verbose vuln add commit tag bump release-auto