/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"reflect"
	"testing"
)

func TestNewReceiver(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReceiver(tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReceiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Receive(t *testing.T) {
	tests := []struct {
		name    string
		s       *Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Receive(); (err != nil) != tt.wantErr {
				t.Errorf("Server.Receive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
