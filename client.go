/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
)

// Client OSC client
type Client struct {
	ip    string
	port  int
	laddr *net.UDPAddr
}

// NewSender Sender作成
func NewSender(ip string, port int) *Client {
	return &Client{ip: ip, port: port, laddr: nil}
}

// Send OSC送信
func (c *Client) Send(m *Message) error {
	if m.Address[0] != '/' {
		fmt.Println("OSC address must start with '/'")
	}

	dataToSend := new(bytes.Buffer)
	m.Address = terminateOSCString(m.Address)
	dataToSend.WriteString(m.Address)

	portStr := strconv.Itoa(c.port)
	udpRAddr, _ := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	conn, _ := net.DialUDP("udp", c.laddr, udpRAddr)
	defer conn.Close()

	// dataToSendにtypetag, OSCアーギュメントを追加
	if _, err := dataToSend.Write(m.Bytes()); err != nil {
		return err
	}

	// OSC送信
	if _, err := conn.Write(dataToSend.Bytes()); err != nil {
		return err
	}
	return nil
}
