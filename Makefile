VERSION := $$(cat version)
ARCH := $$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)
BASE_NAME := terraform-provider-spl_${ARCH}_${VERSION}

default: install

format:
	gofmt -w .

build: format
	go build -o $(BASE_NAME)

install: format
	go install .

debug: format
	go build -gcflags="all=-N -l" -o $(BASE_NAME)_debug
