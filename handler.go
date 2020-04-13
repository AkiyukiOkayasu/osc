/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

// Handler OSC messege handler
type Handler interface {
	ServeOSC(m *Message)
}

// HandlerFunc is created to meet the requirements of Handler interface
// in the definition of the Lambda function.
type HandlerFunc func(*Message)

// serveOSC is created to meet the requirements of Handler interface
// in the definition of the Lambda function.
// serveOSC() will be called from dispatch() of ServeMux()
func (f HandlerFunc) serveOSC(m *Message) {
	f(m)
}
