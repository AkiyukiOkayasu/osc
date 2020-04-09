/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"fmt"
	"net"
	"strconv"
)

// Server OSC server
type Server struct {
	port  int
	laddr *net.UDPAddr
}

// NewReceiver Receiver作成
func NewReceiver(port int) *Server {
	return &Server{port: port, laddr: nil}
}

// Receive OSC受信
func (s *Server) Receive() error {
	portStr := strconv.Itoa(s.port)
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+portStr)
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	var b [512]byte

	for {
		if _, _, err := conn.ReadFromUDP(b[0:]); err != nil {
			return err
		}
		p := string(b[0:])
		m := splitOSCPacket(p)
		// TODO handler implementation
		fmt.Printf("OSC address: %s\n", m.Address)
	}
}
