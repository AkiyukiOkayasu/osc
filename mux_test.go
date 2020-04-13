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
