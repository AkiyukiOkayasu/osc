package main

import (
	"log"

	"github.com/AkiyukiOkayasu/osc"
)

func main() {
	ip := "127.0.0.1" // You can also use "localhost"
	port := 8080
	address := "/test"

	sender := osc.NewSender(ip, port)
	message := osc.NewMessage(address)

	message.AddInt(123)
	message.AddFloat(3.14)
	message.AddString("foo")

	if err := sender.Send(message); err != nil {
		log.Fatalln(err)
	}
}
