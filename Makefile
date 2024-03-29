GOFMT_FILES ?= $$(find . -name '*.go' | grep -v vendor)
VERSION := $$(cat version)
ARCH := $$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)
PLUGIN_PATH := ~/.terraform.d/plugins/local/splight/spl/$(VERSION)/$(ARCH)


default: build

fmt:
	gofmt -w $(GOFMT_FILES)

build: fmt
	./scripts/build.sh

clean-provider-cache:
	rm -rf .terraform.lock.hcl .terraform

install: build
	mkdir -p $(PLUGIN_PATH)
	cp terraform-provider-spl_v$(VERSION) $(PLUGIN_PATH)
