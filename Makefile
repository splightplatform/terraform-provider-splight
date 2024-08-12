# TODO: trace related stuff
SHELL := /bin/bash
VERSION := $(shell cat version)
ARCH := $(shell uname -m)
BASE_NAME := terraform-provider-splight_${ARCH}_${VERSION}

# ANSI color codes
GREEN = \033[0;32m
YELLOW = \033[0;33m
BLUE = \033[0;34m
RESET = \033[0m

.PHONY: default tidy provider install debug snapshot clean

default: tidy provider

tidy:
	@echo -e "$(BLUE)Running go mod tidy$(RESET)"
	@go generate
	@go mod tidy
	@gofmt -w .

provider: tidy format
	@echo -e "$(GREEN)Building provider$(RESET)"
	@go build -o $(BASE_NAME)

install: provider
	@echo -e "$(GREEN)Installing provider$(RESET)"
	@go install $(BASE_NAME)

debug: format
	@echo -e "$(YELLOW)Building debug version$(RESET)"
	@go build -gcflags="all=-N -l" -o $(BASE_NAME)_debug
	@dlv exec $(BASE_NAME)_debug -- -debug
	@$(MAKE) clean

snapshot: provider
	@echo -e "$(GREEN)Creating snapshot$(RESET)"
	@goreleaser --snapshot --clean

clean:
	@echo -e "$(RED)Cleaning up$(RESET)"
	@rm -f $(BASE_NAME) $(BASE_NAME)_debug
