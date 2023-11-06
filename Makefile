.DEFAULT_GOAL := help

# HOST is only used for API specs generation
HOST ?= localhost:8085
batch_name ?= main.go
# Generates a help message. Borrowed from https://github.com/pydanny/cookiecutter-djangopackage.
help: ## Display this help message
	@echo "Please use \`make <target>' where <target> is one of"
	@perl -nle'print $& if m{^[\.a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

depends: ## Install & build dependencies
	go get ./...
	go build ./...
	go mod tidy

mod.clean:
	go clean -modcache

mod:
	go mod tidy && go mod vendor

run: 
	@go run cmd/server/main.go

build-image:
	@docker build --tag bee-go-demo .