// Copyright 2020 Takeshi Yoneda(@mathetake)
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

package runtime

//go:export proxy_on_new_connection
func proxyOnNewConnection(contextID uint32) Action {
	ctx, ok := currentState.streamContexts[contextID]
	if !ok {
		panic("invalid context")
	}
	currentState.setActiveContextID(contextID)
	return ctx.OnNewConnection()
}

//go:export proxy_on_downstream_data
func proxyOnDownstreamData(contextID uint32, dataSize int, endOfStream bool) Action {
	ctx, ok := currentState.streamContexts[contextID]
	if !ok {
		panic("invalid context")
	}
	currentState.setActiveContextID(contextID)
	return ctx.OnDownstreamData(dataSize, endOfStream)
}

//go:export on_downstream_connection_close
func proxyOnDownstreamConnectionClose(contextID uint32, pType PeerType) {
	ctx, ok := currentState.streamContexts[contextID]
	if !ok {
		panic("invalid context")
	}
	currentState.setActiveContextID(contextID)
	ctx.OnDownStreamClose(pType)
}

//go:export proxy_on_upstream_data
func proxyOnUpstreamData(contextID uint32, dataSize int, endOfStream bool) Action {
	ctx, ok := currentState.streamContexts[contextID]
	if !ok {
		panic("invalid context")
	}
	currentState.setActiveContextID(contextID)
	return ctx.OnUpstreamData(dataSize, endOfStream)
}

//go:export proxy_on_upstream_connection_close
func proxyOnUpstreamConnectionClose(contextID uint32, pType PeerType) {
	ctx, ok := currentState.streamContexts[contextID]
	if !ok {
		panic("invalid context")
	}
	currentState.setActiveContextID(contextID)
	ctx.OnUpstreamStreamClose(pType)
}
