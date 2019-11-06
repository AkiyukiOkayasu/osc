package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"./osc"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: oscsender: osc [flags]\n")
	}

	flag.Parse()
	fmt.Println(flag.Args())
	switch flag.Arg(0) {
	case "send":
		if len(flag.Args()) < 4 {
			flag.Usage()
			return
		}

		ip := flag.Arg(1)
		portstr := flag.Arg(2)
		port, _ := strconv.Atoi(portstr)
		oscAddr := flag.Arg(3)
		s := osc.CreateSender(ip, port)
		numOSCArgs := len(flag.Args()) - 4
		oscArgs := make([]interface{}, numOSCArgs)
		for i, o := range flag.Args()[4:] {
			oscArgs[i] = o
		}
		s.Send(oscAddr, oscArgs...)

	case "receive":
		fmt.Println("receive")
		flag.Usage()
	default:
		fmt.Println("default")
		flag.Usage()

	}
}
