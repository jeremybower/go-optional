help: ## Show this help message.
	@echo
	@echo 'usage: make [target]'
	@echo
	@echo 'targets:'
	@echo
	@egrep '^(.+)\:\ ##\ (.+)' ${MAKEFILE_LIST} | column -t -c 2 -s ':#'
	@echo
.PHONY: help

init: ## Initialize the project.
	touch docker-compose.local.yml

test: ## Test the project.
	@mkdir -p coverage
	@go test \
		-cover \
		-covermode=atomic \
		-coverprofile coverage/coverage.out \
		-count=1 \
		-failfast \
		./...
	@go tool cover \
		-html=coverage/coverage.out \
		-o coverage/coverage.html

tidy: ## Tidy the go modules.
	@go mod tidy
.PHONY: tidy

update: ## Update the go modules.
	@$(eval GO_VERSION=$(shell yq '.Config.Go.Version' ../../platform.yml))
	@echo setting go.mod version: $(GO_VERSION)
	@go mod edit -go=$(GO_VERSION)
	@go get -u ./...
	@go mod tidy
.PHONY: update
