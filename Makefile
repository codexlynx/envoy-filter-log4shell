build: bin/filter-log4shell.wasm

bin/filter-log4shell.wasm:
	@DOCKER_BUILDKIT=1 docker build --target binary --output bin .

test:
	$(shell pwd)/test.sh

gofmt:
	gofmt -s -w .

clean:
	rm -rf $(shell pwd)/bin

all: build test
.PHONY: test clean gofmt

