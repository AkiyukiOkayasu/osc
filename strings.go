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
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

const nullChar = '\x00'

// splitOSCPacket split string by OSCstring terminate position
// The maximum number of consecutive null characters is 4.
func splitOSCPacket(str string) (m Message) {
	if str[0] != '/' {
		println("OSC address must start with '/'")
	}

	s := strings.SplitN(str, ",", 2)                      //',' is beginning of OSC typetag
	m.address = strings.TrimRight(s[0], string(nullChar)) //trim nullChar end of OSC address
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
			m.AddInt(v)
			i += 4
		case 'f':
			var v float32
			b := bytes.NewBuffer([]byte(args[i : i+4]))
			if err := binary.Read(b, binary.BigEndian, &v); err != nil {
				fmt.Print("binary.Read failed: ", err)
			}
			m.AddFloat(v)
			i += 4
		case 's':
			// TODO split implementation
			splited, _ := split2OSCStrings(args[i:])
			m.AddString(splited)
			i += len(splited) + numNeededNullChar(len(splited))
		default:
			fmt.Printf("Unexpected OSC typetag: %c\n", t)
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

// numNeededNullChar returns the required number of null characters for the length of the string.
// l: length of OSC string
func numNeededNullChar(l int) int {
	return 4 - (l % 4)
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
