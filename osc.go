/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

// Handler OSC messege handler
type Handler interface {
	Handle(m *Message)
}
