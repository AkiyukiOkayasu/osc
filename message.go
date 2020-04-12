/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"bytes"
	"encoding/binary"
)

// Argument OSC argument
type Argument struct {
	Typetag  rune
	Argument interface{}
}

// Message wraped OSC message
type Message struct {
	Address   string
	Arguments []Argument
}

// AddInt int追加
func (m *Message) AddInt(arg int32) {
	a := Argument{Typetag: 'i', Argument: arg}
	m.Arguments = append(m.Arguments, a)
}

// AddFloat float追加
func (m *Message) AddFloat(arg float32) {
	a := Argument{Typetag: 'f', Argument: arg}
	m.Arguments = append(m.Arguments, a)
}

// AddString string追加
func (m *Message) AddString(arg string) {
	arg = TerminateOSCString(arg)
	a := Argument{Typetag: 's', Argument: arg}
	m.Arguments = append(m.Arguments, a)
}

// Bytes get OSC typetag and argument in []byte
func (m *Message) Bytes() []byte {
	b := new(bytes.Buffer)
	typetag := ","
	for _, a := range m.Arguments {
		typetag += string(a.Typetag)
	}
	typetag = TerminateOSCString(typetag)
	b.WriteString(typetag)

	for _, a := range m.Arguments {
		switch a.Typetag {
		case 'i':
			if v, ok := a.Argument.(int32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 'f':
			if v, ok := a.Argument.(float32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 's':
			if v, ok := a.Argument.(string); ok {
				b.WriteString(v)
			}
		default:
			println("Unexpected typetag")
		}
	}
	return b.Bytes()
}
