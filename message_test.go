/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"reflect"
	"testing"
)

func TestMessage_AddInt(t *testing.T) {
	type args struct {
		arg int32
	}
	tests := []struct {
		name string
		m    *Message
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddInt(tt.args.arg)
		})
	}
}

func TestMessage_AddFloat(t *testing.T) {
	type args struct {
		arg float32
	}
	tests := []struct {
		name string
		m    *Message
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddFloat(tt.args.arg)
		})
	}
}

func TestMessage_AddString(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name string
		m    *Message
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddString(tt.args.arg)
		})
	}
}

func TestMessage_Bytes(t *testing.T) {
	tests := []struct {
		name string
		m    *Message
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
