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

func TestArgument_Type(t *testing.T) {
	tests := []struct {
		name string
		a    *Argument
		want rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Type(); got != tt.want {
				t.Errorf("Argument.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgument_Int(t *testing.T) {
	tests := []struct {
		name  string
		a     *Argument
		want  int32
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.Int()
			if got != tt.want {
				t.Errorf("Argument.Int() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Argument.Int() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArgument_Float(t *testing.T) {
	tests := []struct {
		name  string
		a     *Argument
		want  float32
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.Float()
			if got != tt.want {
				t.Errorf("Argument.Float() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Argument.Float() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArgument_String(t *testing.T) {
	tests := []struct {
		name  string
		a     *Argument
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.String()
			if got != tt.want {
				t.Errorf("Argument.String() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Argument.String() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
