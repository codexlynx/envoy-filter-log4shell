package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
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

		for _, header := range headers {
			proxywasm.LogInfof("%s: %s", header[0], header[1])
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
		proxywasm.LogInfof("%s", body)
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
