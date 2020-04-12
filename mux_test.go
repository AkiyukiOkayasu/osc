/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"reflect"
	"testing"
)

func TestNewServeMux(t *testing.T) {
	tests := []struct {
		name string
		want *ServeMux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServeMux(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServeMux() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServeMux_Handle(t *testing.T) {
	type args struct {
		pattern string
		handler HandlerFunc
	}
	tests := []struct {
		name string
		s    *ServeMux
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Handle(tt.args.pattern, tt.args.handler)
		})
	}
}

func TestServeMux_dispatch(t *testing.T) {
	type args struct {
		m *Message
	}
	tests := []struct {
		name string
		s    *ServeMux
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.dispatch(tt.args.m)
		})
	}
}
