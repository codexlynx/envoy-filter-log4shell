FROM tinygo/tinygo:0.21.0 AS builder

WORKDIR /go/src/github.com/codexlynx/envoy-filter-log4shell

COPY . .

RUN tinygo build -o filter-log4shell.wasm -scheduler=none -target=wasi main.go

FROM scratch AS binary
COPY --from=builder /go/src/github.com/codexlynx/envoy-filter-log4shell/filter-log4shell.wasm /

