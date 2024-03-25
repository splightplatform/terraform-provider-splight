GOFMT_FILES := $$(find . -name '*.go' | grep -v vendor)
VERSION := $$(cat version)
ARCH := $$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)
PROVIDER_PATH := ~/.terraform.d/plugins/local/splight/spl/$(VERSION)/$(ARCH)

default: build

fmt:
	@gofmt -w $(GOFMT_FILES)

build: fmt
	@./scripts/build.sh

install: build
	@mkdir -p $(PROVIDER_PATH)
	@cp terraform-provider-spl_v$(VERSION) $(PROVIDER_PATH)
