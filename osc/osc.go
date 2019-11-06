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

	oscAddr = oscAddr + "0" //OSCアドレスの末尾にはnull文字('0')が必要
	writePaddedString(oscAddr, data)

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

	// typetag, osc argの追加
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

	//typetagをnull文字埋めし、バイト数を4の倍数にする
	//typetagをOSCアドレスの末尾に追加
	writePaddedString(string(typetags), data)

	//その次にOSCアーギュメントを追加
	if _, err := data.Write(oscArgs.Bytes()); err != nil {
		return err
	}

	// TODO 送信処理

	return nil
}

// writePaddedString バイトサイズが4の倍数になるようにnull文字（'0'）埋めする
func writePaddedString(str string, buf *bytes.Buffer) {
	numPadNeeded := len(str) % 4
	for i := 0; i < numPadNeeded; i++ {
		str = str + "0"
	}
	buf.WriteString(str)
}
