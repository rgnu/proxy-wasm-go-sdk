package main

import (
	"testing"

	"github.com/rgnu/proxy-wasm-go-sdk/proxytest"
	"github.com/rgnu/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpBody_OnHttpRequestBody(t *testing.T) {
	opt := proxytest.NewEmulatorOption().
		WithNewHttpContext(newContext)
	host := proxytest.NewHostEmulator(opt)
	defer host.Done()

	id := host.HttpFilterInitContext()
	host.HttpFilterPutRequestBody(id, []byte(`{ "initial": "request body" }`))

	res := host.HttpFilterGetRequestBody(id)
	assert.Equal(t, `{ "another": "body" }`, string(res))

	logs := host.GetLogs(types.LogLevelInfo)
	require.Greater(t, len(logs), 1)

	assert.Equal(t, "on http request body finished", logs[len(logs)-1])
	assert.Equal(t, `initial request body: { "initial": "request body" }`, logs[len(logs)-2])
	assert.Equal(t, "body size: 29", logs[len(logs)-3])
}
