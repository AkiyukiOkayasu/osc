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
