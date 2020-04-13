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
	"log"

	"github.com/AkiyukiOkayasu/osc"
)

const (
	ip      string = "127.0.0.1" // You can also use "localhost"
	port    int    = 8080
	address string = "/test"
)

func main() {
	sender := osc.NewSender(ip, port)
	message := osc.NewMessage(address)

	// Add OSC arguments by type
	message.AddInt(123)
	message.AddFloat(3.14)
	message.AddString("foo")

	// Send OSC
	if err := sender.Send(message); err != nil {
		log.Fatalln(err)
	}
}
