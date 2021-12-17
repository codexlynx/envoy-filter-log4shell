# envoy-filter-log4shell
Plugable Envoy WebAssembly L7 (HTTP) firewall to prevent log4shell vulnerability injections.

[![AUR](https://img.shields.io/github/license/codexlynx/envoy-filter-log4shell)](LICENSE)

## Compilation:
#### Requirements:
* A version of __Docker__ with __BuildKit__ support.
* GNU __make__ utility.

#### Procedure:
* Run: `make`.
* Check the correct creation of `bin/filter-log4shell.wasm` file.

## Resources:
* https://events.istio.io/istiocon-2021/slides/c8p-ExtendingEnvoyWasm-EdSnible.pdf
* https://www.envoyproxy.io/docs/envoy/v1.20.1/configuration/http/http_filters/wasm_filter.html
* https://github.com/istio-ecosystem/wasm-extensions
* https://github.com/proxy-wasm/spec

