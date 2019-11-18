package osc

// OSC送受信パッケージ

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

	oscAddr = oscAddr + "\x00" //OSCアドレスの末尾にはnull文字('\x00')が必要
	writePaddedString(oscAddr, data)

	portStr := strconv.Itoa(c.port)
	udpRAddr, _ := net.ResolveUDPAddr("udp", c.ip+":"+portStr)
	conn, _ := net.DialUDP("udp", c.laddr, udpRAddr)
	defer conn.Close()

	// typetag, osc argの追加
	for _, arg := range args {
		// TODO 型スイッチでできるかも
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

	typetags = typetags + "\x00"      //末尾にnull文字を追加
	writePaddedString(typetags, data) //typetagをOSCアドレスの末尾に追加

	//その次にOSCアーギュメントを追加
	if _, err := data.Write(oscArgs.Bytes()); err != nil {
		return err
	}

	if _, err := conn.Write(data.Bytes()); err != nil {
		return err
	}

	return nil
}

// writePaddedString stringのサイズ（バイト数）を4の倍数に0埋めする
// \x00はnull文字のこと
func writePaddedString(str string, buf *bytes.Buffer) {
	numPadNeeded := 4 - (len(str) % 4)
	for i := 0; i < numPadNeeded; i++ {
		str = str + "\x00"
	}
	buf.WriteString(str)
}
