/*
Package osc is package send and receive OSC(Open Sound Control)
*/
package osc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Client OSC client
type Client struct {
	ip    string
	port  int
	laddr *net.UDPAddr
}

// Server OSC server
type Server struct {
	port  int
	laddr *net.UDPAddr
}

// Message OSC message
type Message struct {
	typetag  rune
	argument interface{}
}

// MessageBuffer hoghoge
type MessageBuffer struct {
	buffer []Message
}

// AddInt int追加
func (buf *MessageBuffer) AddInt(arg int32) {
	m := Message{typetag: 'i', argument: arg}
	buf.buffer = append(buf.buffer, m)
}

// AddFloat float追加
func (buf *MessageBuffer) AddFloat(arg float32) {
	m := Message{typetag: 'f', argument: arg}
	buf.buffer = append(buf.buffer, m)
}

// AddString string追加
func (buf *MessageBuffer) AddString(arg string) {
	padString(&arg)
	m := Message{typetag: 's', argument: arg}
	buf.buffer = append(buf.buffer, m)
}

// CreateSender Sender作成
func CreateSender(ip string, port int) *Client {
	return &Client{ip: ip, port: port, laddr: nil}
}

// CreateReceiver Receiver作成
func CreateReceiver(port int) *Server {
	return &Server{port: port, laddr: nil}
}

// GetBytes get OSC typetag and argument []byte
func (buf *MessageBuffer) GetBytes() []byte {
	b := new(bytes.Buffer)
	typetag := ","
	for _, m := range buf.buffer {
		typetag += string(m.typetag)
	}
	appendNullChar(&typetag)
	padString(&typetag)
	println(typetag)
	b.WriteString(typetag)

	for _, m := range buf.buffer {
		switch m.typetag {
		case 'i':
			if v, ok := m.argument.(int32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 'f':
			if v, ok := m.argument.(float32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 's':
			if v, ok := m.argument.(string); ok {
				b.WriteString(v)
			}
		default:
			println("Unexpected typetag")
		}
	}
	return b.Bytes()
}

// Send OSC送信
func (c *Client) Send(oscAddr string, buf *MessageBuffer) error {
	if oscAddr[0] != '/' {
		fmt.Println("Error: OSCアドレスは'/'から始まる必要があります")
	}

	dataToSend := new(bytes.Buffer)

	// OSCアドレスの末尾にnull文字追加
	appendNullChar(&oscAddr)
	padString(&oscAddr)
	dataToSend.WriteString(oscAddr)

	portStr := strconv.Itoa(c.port)
	udpRAddr, _ := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	conn, _ := net.DialUDP("udp", c.laddr, udpRAddr)
	defer conn.Close()

	// dataToSendにtypetag, OSCアーギュメントを追加
	if _, err := dataToSend.Write(buf.GetBytes()); err != nil {
		return err
	}

	// OSC送信
	if _, err := conn.Write(dataToSend.Bytes()); err != nil {
		return err
	}
	return nil
}

// Receive OSC受信
func (s *Server) Receive(oscAddr string) error {
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

	var buf [512]byte

	for {
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return err
		}
		fmt.Println("Addr: ")
		fmt.Println(addr)
		bufStr := string(buf[0:])
		oscData := strings.SplitN(bufStr, ",", 2)
		oscAddr := oscData[0]
		oscTypesAndArgs := strings.SplitN(oscData[1], "\x00", 2)
		oscTypetag := oscTypesAndArgs[0]
		oscArgs := oscTypesAndArgs[1]
		fmt.Println("OSC address: " + oscAddr)
		println("OSC types: " + oscTypetag)

		argIndexOffset := 4 - ((len(oscTypetag) + 2) % 4) //2は先頭の','と末尾のnull文字
		argIndex := argIndexOffset
		for _, t := range oscTypetag {
			switch t {
			case 'i':
				var i int32
				buf := bytes.NewBuffer([]byte(oscArgs[argIndex : argIndex+4]))
				if err := binary.Read(buf, binary.BigEndian, &i); err != nil {
					fmt.Print("binary.Read failed: ", err)
				}
				argIndex += 4
			case 'f':
				var f float32
				buf := bytes.NewBuffer([]byte(oscArgs[argIndex : argIndex+4]))
				if err := binary.Read(buf, binary.BigEndian, &f); err != nil {
					fmt.Print("binary.Read failed: ", err)
				}
				argIndex += 4
			case 's':
				println("s")
			default:
				println("Unexpected OSC typetag:")
				println(t)
			}
		}
	}
}

// padString stringのサイズを4の倍数に0埋めする
func padString(str *string) {
	for len(*str)%4 != 0 {
		appendNullChar(str)
	}
}

// appendNullChar 末尾にNull文字を追加する
// \x00はnull文字のこと
func appendNullChar(s *string) {
	*s += "\x00"
}
