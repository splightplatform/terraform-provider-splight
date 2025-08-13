SHELL := /bin/bash
BASE_NAME := terraform-provider-splight
DEBUG_BINARY := $(BASE_NAME)_debug

# Go build configuration
GCFLAGS := "all=-N -l"
DEBUG_FLAGS := "-debug"

# Directory structure
RESOURCE_DIR := examples
TEST_DIR := test
TEST_MAIN := $(TEST_DIR)/main.tf

# Dynamic resource discovery
RESOURCES := $(shell find $(RESOURCE_DIR) -name "resource.tf")

.PHONY: default docs tidy provider debug clean clean-debug \
		clean-tests integration-test-all integration-test-one

default: tidy provider

docs:
	@go generate

tidy:
	@go mod tidy
	@go tool gofumpt -w .

provider: tidy
	@go build -o $(BASE_NAME)

debug: tidy
	@go build -gcflags=$(GCFLAGS) -o $(DEBUG_BINARY)
	@trap '$(MAKE) clean-debug' INT TERM EXIT; go tool dlv exec $(DEBUG_BINARY) -- $(DEBUG_FLAGS)

clean: clean-debug clean-tests
	@rm -f $(BASE_NAME)

clean-debug:
	@rm -f $(DEBUG_BINARY)

clean-tests:
	@rm -rf $(TEST_DIR)

# Setup test directory
setup-test-dir:
	@mkdir -p $(TEST_DIR)


# Shortcut to run Terraform in the test directory,
# using the local provider RPC server.
TF_CMD := cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform

# Execute terraform lifecycle for a single resource
terraform-lifecycle:
	@$(TF_CMD) init
	@$(TF_CMD) apply -auto-approve -input=false
	@$(TF_CMD) refresh
	@$(TF_CMD) destroy -auto-approve -input=false

# Cleanup on terraform failure
terraform-cleanup:
	@$(TF_CMD) destroy -auto-approve -input=false

# Test single resource with proper error handling
test-resource:
	@cp $(RESOURCE_FILE) $(TEST_MAIN)
	@if ! $(MAKE) terraform-lifecycle; then \
		$(MAKE) terraform-cleanup; \
		exit 1; \
	fi

# Run integration tests for all resources
integration-test-all: setup-test-dir
	@for resource in $(RESOURCES); do \
		$(MAKE) test-resource RESOURCE_FILE=$$resource; \
	done

# Run integration test for a single resource
integration-test-one: setup-test-dir
	@$(MAKE) test-resource RESOURCE_FILE=$(RESOURCE)
