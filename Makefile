SHELL := /bin/bash
VERSION := $(shell cat version)
ARCH := $(shell uname -m)
BASE_NAME := terraform-provider-splight_${ARCH}_v${VERSION}

# ANSI color codes
GREEN = \033[0;32m
RESET = \033[0m

.PHONY: default tidy docs provider debug snapshot clean

default: docs tidy provider

docs:
	@go generate

tidy:
	@go mod tidy
	@gofmt -w .

provider: tidy
	@echo -e "$(GREEN)Building provider: $(BASE_NAME)$(RESET)"
	@go build -o $(BASE_NAME)

debug: tidy
	@echo -e "$(GREEN)Building debug version: $(BASE_NAME)_debug$(RESET)"
	@go build -gcflags="all=-N -l" -o $(BASE_NAME)_debug
	@echo -e "$(GREEN)Starting debugger...$(RESET)"
	@trap '$(MAKE) clean' INT TERM EXIT; dlv exec $(BASE_NAME)_debug -- -debug
	@$(MAKE) clean

snapshot: provider
	@goreleaser --snapshot --clean

clean:
	@rm -f $(BASE_NAME) $(BASE_NAME)_debug
