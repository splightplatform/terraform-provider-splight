VERSION := $$(cat version)
ARCH := $$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)
BASE_NAME := terraform-provider-spl_${ARCH}_${VERSION}

default: install

generate-docs:
	go generate

format:
	go mod tidy
	gofmt -w .

build: format
	go build -o $(BASE_NAME)

install: format
	go install .

debug-build: format
	go build -gcflags="all=-N -l" -o $(BASE_NAME)_debug

debug-start: debug-build
	dlv exec $(BASE_NAME)_debug -- -debug
