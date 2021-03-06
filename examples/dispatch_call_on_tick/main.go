// Copyright 2020 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/rgnu/proxy-wasm-go-sdk/proxywasm"
)

const tickMilliseconds uint32 = 100

func main() {
	proxywasm.SetNewRootContext(newRootContext)
}

type rootContext struct {
	// you must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultRootContext
	contextID uint32
}

func newRootContext(contextID uint32) proxywasm.RootContext {
	return &rootContext{contextID: contextID}
}

// override
func (ctx *rootContext) OnVMStart(vmConfigurationSize int) bool {
	if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
		proxywasm.LogCriticalf("failed to set tick period: %v", err)
	}
	proxywasm.LogInfof("set tick period milliseconds: %d", tickMilliseconds)
	return true
}

func (ctx *rootContext) OnTick() {
	hs := [][2]string{{":method", "GET"}, {":authority", "some_authority"}, {":path", "/path/to/service"}, {"accept", "*/*"}}
	if _, err := proxywasm.DispatchHttpCall("web_service", hs, "", [][2]string{},
		5000, callback); err != nil {
		proxywasm.LogCriticalf("dispatch httpcall failed: %v", err)
	}
}

var cnt int

func callback(numHeaders, bodySize, numTrailers int) {
	cnt++
	proxywasm.LogInfof("called! %d", cnt)
}
