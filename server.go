/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"context"
	"net"
	"strconv"
)

// Server OSC server
type Server struct {
	port  int
	laddr *net.UDPAddr
	mux   ServeMux
}

// NewReceiver Receiver作成
func NewReceiver(port int, mux ServeMux) *Server {
	return &Server{port: port, laddr: nil, mux: mux}
}

// Receive OSC受信
func (s *Server) Receive(ctx context.Context) error {
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
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if _, _, err := conn.ReadFromUDP(b[0:]); err != nil {
			return err
		}
		p := string(b[0:])
		m := splitOSCPacket(p)
		s.mux.dispatch(&m)
	}
}
