/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"context"
	"reflect"
	"testing"
)

func TestServer_Receive(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Receive(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.Receive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewReceiver(t *testing.T) {
	type args struct {
		port int
		mux  ServeMux
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
			if got := NewReceiver(tt.args.port, tt.args.mux); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReceiver() = %v, want %v", got, tt.want)
			}
		})
	}
}
