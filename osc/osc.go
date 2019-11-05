package osc

import (
	"fmt"
	"net"
	"strconv"
)

// Send OSC送信関数
func Send(ip string, port int, address string) {
	portStr := strconv.Itoa(port)
	udpRAddr, err := net.ResolveUDPAddr("udp", ip+":"+portStr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udpRAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if address[0] != '/' {
		fmt.Println("Error: OSC address")
	}

	fmt.Println("送信")
}
