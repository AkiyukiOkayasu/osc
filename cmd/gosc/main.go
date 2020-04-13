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
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/AkiyukiOkayasu/osc"
)

func main() {
	// Flag usage
	flag.Usage = func() {
		d := "gosc command-line OSC sender/receiver\nSend: gosc send ip port address arguments...\nReceive: gosc receive port\nctrl+c for quit\n"
		fmt.Fprintf(os.Stderr, d)
	}

	flag.Parse()
	switch flag.Arg(0) {
	case "send":
		if len(flag.Args()) < 4 {
			flag.Usage()
			return
		}

		ip := flag.Arg(1)
		portStr := flag.Arg(2)
		port, _ := strconv.Atoi(portStr)
		oscAddr := flag.Arg(3)
		send(ip, port, oscAddr)

	default:
		flag.Usage()
	}
}

func send(ip string, port int, oscAddr string) {
	s := osc.NewSender(ip, port)
	m := osc.NewMessage(oscAddr)

	for i := 4; i < len(flag.Args()); i++ {
		a := flag.Arg(i)
		// int
		if i, err := strconv.Atoi(a); err == nil {
			m.AddInt(int32(i))
			continue
		}

		// float
		if f, err := strconv.ParseFloat(a, 32); err == nil {
			m.AddFloat(float32(f))
			continue
		}

		// string
		m.AddString(a)
	}

	if err := s.Send(m); err != nil {
		log.Fatalln(err)
	}
}
