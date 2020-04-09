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
	typetag  rune
	argument interface{}
}

// Message wraped OSC message
type Message struct {
	Address   string
	arguments []Argument
}

// AddInt int追加
func (m *Message) AddInt(arg int32) {
	a := Argument{typetag: 'i', argument: arg}
	m.arguments = append(m.arguments, a)
}

// AddFloat float追加
func (m *Message) AddFloat(arg float32) {
	a := Argument{typetag: 'f', argument: arg}
	m.arguments = append(m.arguments, a)
}

// AddString string追加
func (m *Message) AddString(arg string) {
	arg = TerminateOSCString(arg)
	a := Argument{typetag: 's', argument: arg}
	m.arguments = append(m.arguments, a)
}

// Bytes get OSC typetag and argument in []byte
func (m *Message) Bytes() []byte {
	b := new(bytes.Buffer)
	typetag := ","
	for _, a := range m.arguments {
		typetag += string(a.typetag)
	}
	typetag = TerminateOSCString(typetag)
	b.WriteString(typetag)

	for _, a := range m.arguments {
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
