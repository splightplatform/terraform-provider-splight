GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

fmt:
	gofmt -w $(GOFMT_FILES)

build: fmt
	./scripts/build.sh