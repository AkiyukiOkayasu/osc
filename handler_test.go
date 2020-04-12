/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import "testing"

func TestHandlerFunc_ServeOSC(t *testing.T) {
	type args struct {
		m *Message
	}
	tests := []struct {
		name string
		f    HandlerFunc
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.ServeOSC(tt.args.m)
		})
	}
}
