/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"bytes"
	"encoding/binary"
)

// Argument OSC argument
// Typetag: OSC typetag
// Argument: OSC argument
type Argument struct {
	typetag  rune
	argument interface{}
}

// Message OSC address and Arguments array
type Message struct {
	Address   string
	Arguments []Argument
}

// AddInt add int to message
func (m *Message) AddInt(arg int32) {
	a := Argument{typetag: 'i', argument: arg}
	m.Arguments = append(m.Arguments, a)
}

// AddFloat add float to message
func (m *Message) AddFloat(arg float32) {
	a := Argument{typetag: 'f', argument: arg}
	m.Arguments = append(m.Arguments, a)
}

// AddString add string to message
func (m *Message) AddString(arg string) {
	arg = TerminateOSCString(arg)
	a := Argument{typetag: 's', argument: arg}
	m.Arguments = append(m.Arguments, a)
}

// Bytes return OSC typetag and argument in []byte
func (m *Message) Bytes() []byte {
	b := new(bytes.Buffer)
	typetag := ","
	for _, a := range m.Arguments {
		typetag += string(a.typetag)
	}
	typetag = TerminateOSCString(typetag)
	b.WriteString(typetag)

	for _, a := range m.Arguments {
		switch a.typetag {
		case 'i':
			if v, ok := a.argument.(int32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 'f':
			if v, ok := a.argument.(float32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 's':
			if v, ok := a.argument.(string); ok {
				b.WriteString(v)
			}
		default:
			println("Unexpected typetag")
		}
	}
	return b.Bytes()
}

// Type return OSC typetag
func (a *Argument) Type() rune {
	return a.typetag
}

// Int return value of int type OSC argument
func (a *Argument) Int() (int32, bool) {
	if a.typetag != 'i' {
		return 0, false
	}

	v, ok := a.argument.(int32)
	return v, ok
}

// Float return value of float type OSC argument
func (a *Argument) Float() (float32, bool) {
	if a.typetag != 'f' {
		return 0.0, false
	}

	v, ok := a.argument.(float32)
	return v, ok
}

// String return value of string type OSC argument
func (a *Argument) String() (string, bool) {
	if a.typetag != 's' {
		return "", false
	}

	v, ok := a.argument.(string)
	return v, ok
}
