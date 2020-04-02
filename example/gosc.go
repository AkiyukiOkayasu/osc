package main

import (
	"flag"
	"fmt"
	"go-osc"
	"log"
	"os"
	"strconv"
)

func main() {
	// Flag usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "gosc command-line OSC sender/receiver\n")
	}

	flag.Parse()
	switch flag.Arg(0) {
	case "send":
		if len(flag.Args()) < 4 {
			// TODO ヘルプ表示
			flag.Usage()
			return
		}

		ip := flag.Arg(1)
		portStr := flag.Arg(2)
		port, _ := strconv.Atoi(portStr)
		oscAddr := flag.Arg(3)
		send(ip, port, oscAddr)

	case "receive":
		// TODO ヘルプ表示
		portStr := flag.Arg(1)
		port, _ := strconv.Atoi(portStr)
		receive(port, "/test") // TODO OSCアドレス周り追加

	default:
		// TODO add flag usage
		flag.Usage()
	}
}

func send(ip string, port int, oscAddr string) {
	s := osc.CreateSender(ip, port)
	numOSCArgs := len(flag.Args()) - 4
	oscArgs := make([]interface{}, numOSCArgs)
	for i, o := range flag.Args()[4:] {
		oscArgs[i] = o
	}
	if err := s.Send(oscAddr, oscArgs...); err != nil {
		log.Fatalln(err)
	}
}

func receive(port int, oscAddr string) {
	r := osc.CreateReceiver(port)
	go r.Receive(oscAddr)
}
