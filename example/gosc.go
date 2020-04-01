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
		portStr := flag.Arg(1)
		port, _ := strconv.Atoi(portStr)
		r := osc.CreateReceiver(port)
		err := r.Receive("/test")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("begine")
		for {

		}
		fmt.Println("end")

	default:
		// TODO add flag usage
		flag.Usage()
	}
}
