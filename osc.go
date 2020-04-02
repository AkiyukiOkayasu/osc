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

// CreateSender Sender作成関数
func CreateSender(ip string, port int) *Client {
	return &Client{ip: ip, port: port, laddr: nil}
}

// CreateReceiver Receiver作成関数
func CreateReceiver(port int) *Server {
	return &Server{port: port, laddr: nil}
}

// Send OSC送信関数
func (c *Client) Send(oscAddr string, args ...interface{}) error {
	if oscAddr[0] != '/' {
		fmt.Println("Error: OSCアドレスは'/'から始まる必要があります")
	}

	data := new(bytes.Buffer)
	oscArgs := new(bytes.Buffer)
	typetags := "," // OSC typetagの先頭は','

	oscAddr = appendNullChar(oscAddr)
	writePaddedString(oscAddr, data)

	portStr := strconv.Itoa(c.port)
	udpRAddr, _ := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	conn, _ := net.DialUDP("udp", c.laddr, udpRAddr)
	defer conn.Close()

	// typetag, osc argの追加
	for _, arg := range args {
		argStr := arg.(string)
		if i, err := strconv.Atoi(argStr); err == nil {
			typetags = typetags + "i"
			binary.Write(oscArgs, binary.BigEndian, int32(i))
			continue
		}

		if f, err := strconv.ParseFloat(argStr, 32); err == nil {
			typetags = typetags + "f"
			binary.Write(oscArgs, binary.BigEndian, float32(f))
			continue
		}

		typetags = typetags + "s"
		writePaddedString(argStr, oscArgs)
	}

	typetags = appendNullChar(typetags)
	writePaddedString(typetags, data) // typetagをOSCアドレス末尾に追加

	// OSCアーギュメントを追加
	if _, err := data.Write(oscArgs.Bytes()); err != nil {
		return err
	}

	if _, err := conn.Write(data.Bytes()); err != nil {
		return err
	}

	return nil
}

// Receive OSC受信関数
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
		// fmt.Println(string(buf[0:]))
		bufStr := string(buf[0:])
		oscData := strings.SplitN(bufStr, ",", 2)
		oscAddr := oscData[0]
		fmt.Println("OSC address: " + oscAddr)

		counter := 1
		for {
			i := counter*4 - 1
			if oscData[1][i] == 0 {
				oscTypeTag := oscData[1][0:i]
				strings.TrimRight(oscTypeTag, "\x00") //null文字削除
				fmt.Println("OSC typetag: " + oscTypeTag)
				oscArgs := []byte(oscData[1][i:])
				for pos, c := range oscTypeTag {
					bindex := pos * 4
					switch c {
					case 'i':
						num := 0
						buf := bytes.NewBuffer(oscArgs[bindex : bindex+4])
						fmt.Println(buf)
						binary.Read(buf, binary.BigEndian, &num)
						fmt.Println(num)

					case 'f':
						num := 0.0
						buf := bytes.NewBuffer(oscArgs[bindex : bindex+4])
						fmt.Println(buf)
						binary.Read(buf, binary.BigEndian, &num)
						fmt.Println(num)

					default:
						fmt.Println("default")
					}
				}
				break
			}
			counter++
		}
	}
}

// writePaddedString stringのサイズ（バイト数）を4の倍数に0埋めする
// \x00はnull文字のこと
func writePaddedString(str string, buf *bytes.Buffer) {
	for len(str)%4 != 0 {
		str = str + "\x00"
	}
	buf.WriteString(str)
}

// appendNullChar 末尾にNull文字を追加する
func appendNullChar(s string) string {
	return s + "\x00"
}
