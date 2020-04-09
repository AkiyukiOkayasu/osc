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

// Argument OSC argument
type Argument struct {
	typetag  rune
	argument interface{}
}

// Message wraped OSC message
type Message struct {
	Address   string
	arguments []Argument
}

// AddInt int追加
func (m *Message) AddInt(arg int32) {
	a := Argument{typetag: 'i', argument: arg}
	m.arguments = append(m.arguments, a)
}

// AddFloat float追加
func (m *Message) AddFloat(arg float32) {
	a := Argument{typetag: 'f', argument: arg}
	m.arguments = append(m.arguments, a)
}

// AddString string追加
func (m *Message) AddString(arg string) {
	arg = TerminateOSCString(arg)
	a := Argument{typetag: 's', argument: arg}
	m.arguments = append(m.arguments, a)
}

// Bytes get OSC typetag and argument in []byte
func (m *Message) Bytes() []byte {
	b := new(bytes.Buffer)
	typetag := ","
	for _, a := range m.arguments {
		typetag += string(a.typetag)
	}
	typetag = TerminateOSCString(typetag)
	b.WriteString(typetag)

	for _, a := range m.arguments {
		switch a.typetag {
		case 'i':
			if v, ok := a.argument.(int32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 'f':
			if v, ok := a.argument.(float32); ok {
				binary.Write(b, binary.BigEndian, v)
			}
		case 's':
			if v, ok := a.argument.(string); ok {
				b.WriteString(v)
			}
		default:
			println("Unexpected typetag")
		}
	}
	return b.Bytes()
}

// Handler OSC messege handler
type Handler interface {
	Handle(m *Message)
}

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

// NewSender Sender作成
func NewSender(ip string, port int) *Client {
	return &Client{ip: ip, port: port, laddr: nil}
}

// NewReceiver Receiver作成
func NewReceiver(port int) *Server {
	return &Server{port: port, laddr: nil}
}

// Send OSC送信
func (c *Client) Send(m *Message) error {
	if m.Address[0] != '/' {
		fmt.Println("Error: OSCアドレスは'/'から始まる必要があります")
	}

	dataToSend := new(bytes.Buffer)
	m.Address = TerminateOSCString(m.Address)
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

// splitOSCPacket split string by OSCstring terminate position
// null文字は4つまでしか連続しない
func splitOSCPacket(str string) (m Message) {
	if str[0] != '/' {
		println("OSCアドレスは/から始まる必要があります")
	}

	s := strings.SplitN(str, ",", 2) //',' is beginning of OSC typetag
	m.Address = s[0]
	typetagAndArgs := "," + s[1]
	typetag, args := split2OSCStrings(typetagAndArgs)
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
			m.AddInt(v)
			i += 4
		case 'f':
			var v float32
			b := bytes.NewBuffer([]byte(args[i : i+4]))
			if err := binary.Read(b, binary.BigEndian, &v); err != nil {
				fmt.Print("binary.Read failed: ", err)
			}
			fmt.Printf("f: %3f\n", v)
			m.AddFloat(v)
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
