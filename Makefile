build: bin/filter-log4shell.wasm

bin/filter-log4shell.wasm:
	@DOCKER_BUILDKIT=1 docker build --target binary --output bin .

all: build
