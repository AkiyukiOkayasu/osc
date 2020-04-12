/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import "fmt"

// Handler OSC messege handler
type Handler interface {
	ServeOSC(m *Message)
}

// HandlerFunc is created to meet the requirements of Handler interface
// in the definition of the Lambda function.
type HandlerFunc func(*Message)

// ServeOSC is created to meet the requirements of Handler interface
// in the definition of the Lambda function.
func (f HandlerFunc) ServeOSC(m *Message) {
	fmt.Println("HandlerFunc ServeOSC()")
	f(m)
}
