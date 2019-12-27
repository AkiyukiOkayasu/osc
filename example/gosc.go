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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "gosc command-line OSC sender/receiver\n")
	}

	flag.Parse()
	switch flag.Arg(0) {
	case "send":
		if len(flag.Args()) < 4 {
			// TODO add flag usage
			flag.Usage()
			return
		}

		ip := flag.Arg(1)
		portStr := flag.Arg(2)
		port, _ := strconv.Atoi(portStr)
		oscAddr := flag.Arg(3)
		s := osc.CreateSender(ip, port)
		numOSCArgs := len(flag.Args()) - 4
		oscArgs := make([]interface{}, numOSCArgs)
		for i, o := range flag.Args()[4:] {
			oscArgs[i] = o
		}
		if err := s.Send(oscAddr, oscArgs...); err != nil {
			log.Fatalln(err)
		}

	case "receive":
		// TODO add flag usage
		// flag.Usage()
	default:
		// TODO add flag usage
		flag.Usage()
	}
}
