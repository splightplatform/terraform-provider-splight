VERSION := $$(cat version)
ARCH := $$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)


default: install

format:
	gofmt -w .

build: format
	go build -o terraform-provider-spl_${ARCH}_${VERSION}

install: format
	go install .
