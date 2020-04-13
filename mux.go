/*
Copyright 2020 Akiyuki Okayasu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

// ServeMux OSC handler mux
// key: OSC address
// value: handler function
type ServeMux struct {
	m map[string]Handler
}

// NewServeMux return ServeMux
func NewServeMux() *ServeMux {
	return &ServeMux{map[string]Handler{}}
}

// Handle add handler for an OSC address
func (s *ServeMux) Handle(pattern string, handler HandlerFunc) {
	s.m[pattern] = handler
}

// dispatch handler
func (s *ServeMux) dispatch(m *Message) {
	if h, ok := s.m[m.address]; ok {
		h.serveOSC(m)
	}
}
