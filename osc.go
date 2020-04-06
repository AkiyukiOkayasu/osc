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

const nullChar = '\x00'

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

// Argument OSC argument
type Argument struct {
	typetag  rune
	argument interface{}
}

// ArgumentBuffer OSC argument buffer
type ArgumentBuffer struct {
	buffer []Argument
}

// AddInt int追加
func (buf *ArgumentBuffer) AddInt(arg int32) {
	a := Argument{typetag: 'i', argument: arg}
	buf.buffer = append(buf.buffer, a)
}

// AddFloat float追加
func (buf *ArgumentBuffer) AddFloat(arg float32) {
	a := Argument{typetag: 'f', argument: arg}
	buf.buffer = append(buf.buffer, a)
}

// AddString string追加
func (buf *ArgumentBuffer) AddString(arg string) {
	arg = terminateOSCString(arg)
	a := Argument{typetag: 's', argument: arg}
	buf.buffer = append(buf.buffer, a)
}

// NewSender Sender作成
func NewSender(ip string, port int) *Client {
	return &Client{ip: ip, port: port, laddr: nil}
}

// NewReceiver Receiver作成
func NewReceiver(port int) *Server {
	return &Server{port: port, laddr: nil}
}

// Bytes get OSC typetag and argument in []byte
func (buf *ArgumentBuffer) Bytes() []byte {
	b := new(bytes.Buffer)
	typetag := ","
	for _, m := range buf.buffer {
		typetag += string(m.typetag)
	}
	typetag = terminateOSCString(typetag)
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
func (c *Client) Send(oscAddr string, buf *ArgumentBuffer) error {
	if oscAddr[0] != '/' {
		fmt.Println("Error: OSCアドレスは'/'から始まる必要があります")
	}

	dataToSend := new(bytes.Buffer)
	oscAddr = terminateOSCString(oscAddr)
	dataToSend.WriteString(oscAddr)

	portStr := strconv.Itoa(c.port)
	udpRAddr, _ := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	conn, _ := net.DialUDP("udp", c.laddr, udpRAddr)
	defer conn.Close()

	// dataToSendにtypetag, OSCアーギュメントを追加
	if _, err := dataToSend.Write(buf.Bytes()); err != nil {
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

	var b [512]byte

	for {
		if _, _, err := conn.ReadFromUDP(b[0:]); err != nil {
			return err
		}
		p := string(b[0:])
		addr, _ := splitOSCPacket(p)
		// TODO handler implementation
		fmt.Printf("OSC address: %s\n", addr)
	}
}

// splitOSCPacket split string by OSCstring terminate position
// null文字は4つまでしか連続しない
func splitOSCPacket(str string) (oscAddr string, buf ArgumentBuffer) {
	buf = ArgumentBuffer{}
	if str[0] != '/' {
		println("OSCアドレスは/から始まる必要があります")
	}

	s := strings.SplitN(str, ",", 2) //',' is beginning of OSC typetag
	oscAddr = s[0]
	fmt.Printf("OSC address: %s\n", oscAddr)
	typetagAndArgs := "," + s[1]
	typetag, args := split2OSCStrings(typetagAndArgs)
	fmt.Printf("typetag in String: %s\n", typetag)
	fmt.Printf("args: %x\n", args)
	if typetag[0] != ',' {
		println("OSC typetagは,から始まる必要があります")
	}
	typetag = typetag[1:]

	i := 0
	for _, t := range typetag {
		switch t {
		case 'i':
			var v int32
			b := bytes.NewBuffer([]byte(args[i : i+4]))
			if err := binary.Read(b, binary.BigEndian, &v); err != nil {
				fmt.Print("binary.Read failed: ", err)
			}
			fmt.Printf("i: %d\n", v)
			buf.AddInt(v)
			i += 4
		case 'f':
			var v float32
			b := bytes.NewBuffer([]byte(args[i : i+4]))
			if err := binary.Read(b, binary.BigEndian, &v); err != nil {
				fmt.Print("binary.Read failed: ", err)
			}
			fmt.Printf("f: %3f\n", v)
			buf.AddFloat(v)
			i += 4
		case 's':
			println("s")
		default:
			println("Unexpected OSC typetag:")
			println(t)
		}
	}
	return
}

// TerminateOSCString terminate OSC string
func TerminateOSCString(str string) string {
	str += string(nullChar) //Add null char at least 1
	for len(str)%4 != 0 {
		str += string(nullChar)
	}
	return str
}

// numNeededNullChar is count nullChar for pad 4bytes
func numNeededNullChar(l int) int {
	n := 0
	if l%4 != 0 {
		n = 4 - (l % 4)
	}
	fmt.Printf("%d: %d\n", l, n)
	return n
}

// split2OSCStrings hogehgoe
func split2OSCStrings(s string) (string, string) {
	isNullChar := false
	var splited string
	var remainds string
	for i, r := range s {
		if !isNullChar {
			if r == nullChar {
				isNullChar = true
				continue
			}
		}

		if isNullChar {
			if i%4 == 0 {
				splited = s[:i]
				remainds = s[i:]
				break
			}
		}
	}

	splited = strings.TrimRight(splited, string(nullChar)) //return splited osc string without null char
	return splited, remainds
}
