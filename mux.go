/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

// ServeMux OSC handler mux
type ServeMux struct {
	m map[string]Handler
}

// NewServeMux return ServeMux
func NewServeMux() *ServeMux {
	return &ServeMux{map[string]Handler{}}
}

// Handle TODO add description
func (s *ServeMux) Handle(pattern string, handler HandlerFunc) {
	s.m[pattern] = handler
}

// dispatch handler
func (s *ServeMux) dispatch(m *Message) {
	if h, ok := s.m[m.Address]; ok {
		h.ServeOSC(m)
	}
}
