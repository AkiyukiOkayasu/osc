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
	typetags := "," // OSC typetagの先頭は','

	oscAddr = oscAddr + "0" //OSCアドレスの末尾にはnull文字('0')が必要
	writePaddedString(oscAddr, data)

	portStr := strconv.Itoa(c.port)
	udpRAddr, _ := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	conn, _ := net.DialUDP("udp", c.laddr, udpRAddr)
	defer conn.Close()

	// typetag, osc argの追加
	fmt.Println(args)
	for _, arg := range args {
		switch arg.(type) {
		case int32, int64:
			fmt.Println("int")
			typetags = typetags + "i"
			if err := binary.Write(oscArgs, binary.BigEndian, arg.(int32)); err != nil {
				fmt.Println("Error: endian")
			}

		case float32, float64:
			fmt.Println("float")
			typetags = typetags + "f"
			if err := binary.Write(oscArgs, binary.BigEndian, arg.(float32)); err != nil {
				fmt.Print("Error: endian")
			}

		case string:
			fmt.Println("string")
			typetags = typetags + "s"
			writePaddedString(arg.(string), oscArgs)
		default:
			fmt.Println("default")
		}
	}

	//typetagをOSCアドレスの末尾に追加
	writePaddedString(typetags, data)

	//その次にOSCアーギュメントを追加
	if _, err := data.Write(oscArgs.Bytes()); err != nil {
		return err
	}

	if _, err := conn.Write(data.Bytes()); err != nil {
		return err
	}
	fmt.Println("send, ", len(data.Bytes()), ", ", data)

	return nil
}

// writePaddedString stringのサイズ（バイト数）を4の倍数に0埋めする
// 0はnull文字のこと
func writePaddedString(str string, buf *bytes.Buffer) {
	numPadNeeded := 4 - (len(str) % 4)
	for i := 0; i < numPadNeeded; i++ {
		str = str + "\x00"
	}
	buf.WriteString(str)
}
