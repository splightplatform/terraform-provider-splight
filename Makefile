SHELL := /bin/bash
BASE_NAME := terraform-provider-splight
DEBUG_BINARY := $(BASE_NAME)_debug

# ANSI color codes
GREEN = \033[0;32m
RESET = \033[0m

# Go build flags
GCFLAGS := "all=-N -l"
DEBUG_FLAGS := "-debug"

.PHONY: default docs tidy provider debug snapshot clean

default: tidy provider

docs:
	@go generate

tidy:
	@go mod tidy
	@gofumpt -w .

provider: tidy
	@echo -e "$(GREEN)Building provider: $(BASE_NAME)$(RESET)"
	@go build -o $(BASE_NAME)

debug: tidy
	@go build -gcflags=$(GCFLAGS) -o $(DEBUG_BINARY)
	@trap '$(MAKE) clean-debug' INT TERM EXIT; dlv exec $(DEBUG_BINARY) -- $(DEBUG_FLAGS)

clean:
	@rm -f $(BASE_NAME) $(DEBUG_BINARY)

clean-debug:
	@rm -f $(DEBUG_BINARY)
