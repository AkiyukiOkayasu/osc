package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AkiyukiOkayasu/osc"
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
		receive(port)

	default:
		// TODO add flag usage
		flag.Usage()
	}
}

func send(ip string, port int, oscAddr string) {
	s := osc.NewSender(ip, port)
	m := osc.Message{Address: oscAddr}
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
		a = osc.TerminateOSCString(a)
		m.AddString(a)
	}

	if err := s.Send(&m); err != nil {
		log.Fatalln(err)
	}
}

func receive(port int) {
	c := context.Background()
	ctx, cancel := context.WithCancel(c)
	mux := osc.NewServeMux()
	mux.Handle("/test", func(m *osc.Message) {
		fmt.Println("/test handler begin")
		for _, a := range m.Arguments {
			fmt.Println(a.Typetag)
		}
		fmt.Println("/test handler end")
	})
	r := osc.NewReceiver(port, *mux)
	go r.Receive(ctx)
	defer fmt.Println("Finished")
	time.Sleep(30 * time.Second)
	cancel() // Stop receiving OSC
}
