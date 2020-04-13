/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

const nullChar = '\x00'

// splitOSCPacket split string by OSCstring terminate position
// null文字は4つまでしか連続しない
func splitOSCPacket(str string) (m Message) {
	if str[0] != '/' {
		println("OSC address must start with '/'")
	}

	s := strings.SplitN(str, ",", 2)                      //',' is beginning of OSC typetag
	m.Address = strings.TrimRight(s[0], string(nullChar)) //trim nullChar end of OSC address
	typetagAndArgs := "," + s[1]
	typetag, args := split2OSCStrings(typetagAndArgs)
	if typetag[0] != ',' {
		println("OSC typetag must start with ','")
	}
	typetag = typetag[1:]

	i := 0
	for _, t := range typetag {
		switch t {
		case 'i':
			var v int32
			b := bytes.NewBuffer([]byte(args[i : i+4]))
			if err := binary.Read(b, binary.BigEndian, &v); err != nil {
				fmt.Print("binary.Read failed: ", err)
			}
			fmt.Printf("i: %d\n", v)
			m.AddInt(v)
			i += 4
		case 'f':
			var v float32
			b := bytes.NewBuffer([]byte(args[i : i+4]))
			if err := binary.Read(b, binary.BigEndian, &v); err != nil {
				fmt.Print("binary.Read failed: ", err)
			}
			fmt.Printf("f: %3f\n", v)
			m.AddFloat(v)
			i += 4
		case 's':
			println("s")
		default:
			println("Unexpected OSC typetag:")
			println(t)
		}
	}
	return
}

// terminateOSCString terminate OSC string
func terminateOSCString(str string) string {
	str += string(nullChar) //Add null char at least 1
	for len(str)%4 != 0 {
		str += string(nullChar)
	}
	return str
}

// numNeededNullChar is count nullChar for pad 4bytes
func numNeededNullChar(l int) int {
	n := 0
	if l%4 != 0 {
		n = 4 - (l % 4)
	}
	return n
}

// split2OSCStrings hogehgoe
func split2OSCStrings(s string) (string, string) {
	isNullChar := false
	var splited string
	var remainds string
	for i, r := range s {
		if !isNullChar {
			if r == nullChar {
				isNullChar = true
				continue
			}
		}

		if isNullChar {
			if i%4 == 0 {
				splited = s[:i]
				remainds = s[i:]
				break
			}
		}
	}

	splited = strings.TrimRight(splited, string(nullChar)) //return splited osc string without null char
	return splited, remainds
}
