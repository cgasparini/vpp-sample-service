// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// HTTPHandlers is an interface that is useful for other plugins that need to register HTTP Handlers.
// Use this interface as type for the field in terms of dependency injection.
type HTTPHandlers interface {
	// RegisterHTTPHandler propagates to Gorilla mux
	RegisterHTTPHandler(path string,
		handler func(formatter *render.Render) http.HandlerFunc,
		methods ...string) *mux.Route
}
