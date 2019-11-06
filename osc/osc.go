package osc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

// Client OSC送信オブジェクト
type Client struct {
	ip    string
	port  int
	laddr *net.UDPAddr
}

// CreateSender Sender作成関数
func CreateSender(ip string, port int) *Client {
	return &Client{ip: ip, port: port, laddr: nil}
}

// Send OSC送信関数
func (c *Client) Send(oscAddr string, args ...interface{}) error {
	if oscAddr[0] != '/' {
		fmt.Println("Error: OSCアドレスは'/'から始まる必要があります")
	}

	data := new(bytes.Buffer)
	oscArgs := new(bytes.Buffer)
	typetags := []byte{','} // stringでもいいかも

	// OSCアドレス
	data.WriteString(oscAddr)
	data.WriteString('0')         //null文字
	padneeded := len(oscAddr) % 4 //TODO 4バイト

	portStr := strconv.Itoa(c.port)
	udpRAddr, err := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udpRAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// typetag, osc arg
	for _, arg := range args {
		fmt.Println("%v", arg)
		switch t := arg.(type) {
		case int32:
			typetags = append(typetags, 'i')
			err := binary.Write(oscArgs, binary.BigEndian, int32(t))

		case float32:
			typetags = append(typetags, 'f')

		case string:
			typetags = append(typetags, 's')
		}
	}

	// TODO 送信処理

	return nil
}
