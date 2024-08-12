SHELL := /bin/bash
VERSION := $(shell cat version)
ARCH := $(shell uname -m)
BASE_NAME := terraform-provider-splight_${ARCH}_v${VERSION}

# ANSI color codes
GREEN = \033[0;32m
RESET = \033[0m

.PHONY: default docs tidy provider debug snapshot clean

default: tidy provider

docs:
	@go generate

tidy:
	@go mod tidy
	@gofmt -w .

provider: tidy
	@echo -e "$(GREEN)Building provider: $(BASE_NAME)$(RESET)"
	@go build -o $(BASE_NAME)

debug: tidy
	# TODO: only output my logs or sdk: see https://developer.hashicorp.com/terraform/plugin/log/managing#enable-logging
	@echo -e "$(GREEN)Building temporary binary: $(BASE_NAME)_debug$(RESET)"
	@go build -gcflags="all=-N -l" -o $(BASE_NAME)_debug
	@echo -e "$(GREEN)Starting debugger...$(RESET)"
	@trap 'rm -f $(BASE_NAME)_debug' INT TERM EXIT; dlv exec $(BASE_NAME)_debug -- -debug
	@rm -f $(BASE_NAME)_debug

clean:
	@rm -f $(BASE_NAME)
