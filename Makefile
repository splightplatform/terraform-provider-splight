SHELL := /bin/bash
VERSION := $(shell cat version)
ARCH := $(shell uname -m)
BASE_NAME := terraform-provider-splight_${ARCH}_v${VERSION}
DEBUG_BINARY := $(BASE_NAME)_debug

# ANSI color codes
GREEN = \033[0;32m
RESET = \033[0m

# Go build flags
GCFLAGS := "all=-N -l"

.PHONY: default docs tidy provider debug snapshot clean clean-debug

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
	@echo -e "$(GREEN)Debug binary: $(DEBUG_BINARY)$(RESET)"
	@go build -gcflags=$(GCFLAGS) -o $(DEBUG_BINARY)
	@trap '$(MAKE) clean-debug' INT TERM EXIT; dlv exec $(DEBUG_BINARY) -- -debug

clean:
	@rm -f $(BASE_NAME)

clean-debug:
	@rm -f $(DEBUG_BINARY)
