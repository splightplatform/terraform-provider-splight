SHELL := /bin/bash
BASE_NAME := terraform-provider-splight
DEBUG_BINARY := $(BASE_NAME)_debug

# Go build flags
GCFLAGS := "all=-N -l"
DEBUG_FLAGS := "-debug"

# Test vars
RESOURCE_DIR := examples
RESOURCES := $(shell find $(RESOURCE_DIR) -name "resource.tf")
TEST_DIR := test
TEST_MAIN := $(TEST_DIR)/main.tf

.PHONY: default docs tidy provider debug snapshot clean clean-debug \
        integration-test-all integration-test-one clean-tests

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

clean:
	@rm -f $(BASE_NAME) $(DEBUG_BINARY)
	@rm -rf $(TEST_DIR)

clean-debug:
	@rm -f $(DEBUG_BINARY)

integration-test-all:
	@if [ -z "$(TF_REATTACH_PROVIDERS)" ]; then \
		echo "Usage: make integration-test-all TF_REATTACH_PROVIDERS='...'"; \
		exit 1; \
	fi
	@mkdir -p $(TEST_DIR)
	@set -e; for file in $(RESOURCES); do \
		echo "Testing resource: $$file"; \
		cp "$$file" $(TEST_MAIN); \
		(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform init >/dev/null 2>&1); \
		(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform apply -auto-approve); \
		(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform refresh); \
		(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform destroy -auto-approve); \
	done

integration-test-one:
	@if [ -z "$(RESOURCE)" ]; then \
		echo "Usage: make integration-test-one RESOURCE=path/to/resource.tf TF_REATTACH_PROVIDERS='...'"; \
		exit 1; \
	fi
	@mkdir -p $(TEST_DIR)
	@cp $(RESOURCE) $(TEST_MAIN)
	@(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform init >/dev/null 2>&1)
	@(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform apply -auto-approve)
	@(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform refresh)
	@(cd $(TEST_DIR) && TF_REATTACH_PROVIDERS='$(TF_REATTACH_PROVIDERS)' terraform destroy -auto-approve)

clean-tests:
	@rm -rf $(TEST_DIR)
