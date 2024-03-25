GOFMT_FILES := $$(find . -name '*.go' | grep -v vendor)
VERSION := $$(cat version)
ARCH := $$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)
PLUGIN_PATH := ~/.terraform.d/plugins/local/splight/spl/$(VERSION)/$(ARCH)
LOCAL_PROVIDER_PATH := .terraform/providers/local/splight/spl/$(VERSION)/$(ARCH)


default: build

fmt:
	gofmt -w $(GOFMT_FILES)

build: fmt
	./scripts/build.sh

clean:
	rm -rf $(LOCAL_PROVIDER_PATH)
	rm .terraform.lock.hcl

install: build clean
	mkdir -p $(PLUGIN_PATH)
	cp terraform-provider-spl_v$(VERSION) $(PLUGIN_PATH)

