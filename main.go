package main

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"regexp"
)

var (
	blockMessage = "This request has been blocked for security reasons."

	regexRules = []string{
		"(?i)(\\$|%24)(\\{|%7b).*j.*n.*d.*i.*(:|%3a)",
	}
)

type vmContext struct {
	types.DefaultVMContext
}

type pluginContext struct {
	types.DefaultPluginContext
}

type httpContext struct {
	types.DefaultHttpContext
}

func (*httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	if endOfStream {
		headers, err := proxywasm.GetHttpRequestHeaders()
		if err != nil {
			proxywasm.LogCriticalf("failed to get request headers: %v", err)
		}

		fullHeaders := fmt.Sprintf("%v", headers)
		for _, rule := range regexRules {
			regex, err := regexp.Compile(rule)
			if err != nil {
				proxywasm.LogCriticalf("%v", err)
			} else {
				if regex.MatchString(fullHeaders) {
					err := proxywasm.SendHttpResponse(302, nil, []byte(blockMessage), -1)
					if err != nil {
						proxywasm.LogCriticalf("%v", err)
					}

					return types.ActionPause
				}
			}
		}

		return types.ActionContinue
	}
	return types.ActionPause
}

func (*httpContext) OnHttpRequestBody(bodySize int, endOfStream bool) types.Action {
	if endOfStream {
		body, err := proxywasm.GetHttpRequestBody(0 , bodySize)
		if err != nil {
			proxywasm.LogCriticalf("failed to get request body: %v", err)
		}

		for _, rule := range regexRules {
			regex, err := regexp.Compile(rule)
			if err != nil {
				proxywasm.LogCriticalf("%v", err)
			} else {
				if regex.MatchString(string(body)) {
					err := proxywasm.SendHttpResponse(302, nil, []byte(blockMessage), -1)
					if err != nil {
						proxywasm.LogCriticalf("%v", err)
					}

					return types.ActionPause
				}
			}
		}

		return types.ActionContinue
	}
	return types.ActionPause
}

func (*pluginContext) NewHttpContext(contextId uint32) types.HttpContext {
	return &httpContext{}
}

func (*vmContext) NewPluginContext(contextId uint32) types.PluginContext {
	return &pluginContext{}
}

func main() {
	proxywasm.SetVMContext(&vmContext{})
}
