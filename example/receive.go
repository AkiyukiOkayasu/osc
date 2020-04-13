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

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AkiyukiOkayasu/osc"
)

const port int = 8080

func main() {
	// context to cancel OSC receiving gorouting
	c := context.Background()
	ctx, cancel := context.WithCancel(c)

	// OSC handler for /foo
	mux := osc.NewServeMux()
	mux.Handle("/foo", func(m *osc.Message) {
		fmt.Printf("OSC Address: %s\n", m.Address)
		for _, a := range m.Arguments {
			switch a.Type() {
			case 'i':
				if v, ok := a.Int(); ok {
					fmt.Printf("Foo Int: %d\n", v)
				}
			case 'f':
				if v, ok := a.Float(); ok {
					fmt.Printf("Foo Float: %3f\n", v)
				}
			case 's':
				if v, ok := a.String(); ok {
					fmt.Printf("Foo String: %s\n", v)
				}
			default:
				fmt.Printf("Unexpected type: %v\n", a.Type())
			}
		}
	})

	// Another OSC handler for /bar
	mux.Handle("/bar", func(m *osc.Message) {
		fmt.Printf("OSC Address: %s\n", m.Address)
		for _, a := range m.Arguments {
			switch a.Type() {
			case 'i':
				if v, ok := a.Int(); ok {
					fmt.Printf("Bar Int: %d\n", v)
				}
			case 'f':
				if v, ok := a.Float(); ok {
					fmt.Printf("Bar Float: %3f\n", v)
				}
			case 's':
				if v, ok := a.String(); ok {
					fmt.Printf("Bar String: %s\n", v)
				}
			default:
				fmt.Printf("Unexpected type: %v\n", a.Type())
			}
		}
	})

	r := osc.NewReceiver(port, *mux)
	go r.Receive(ctx) // Start OSC receiving

	// Something to do...
	time.Sleep(30 * time.Second) // Sleep 30 seconds

	cancel() // Stop receiving OSC
}
