# osc  

OSC (Open Sound Control) package for Go.  
A part of The [Open Sound Control 1.0 Specification](http://opensoundcontrol.org/spec-1_0) has been implemented in pure Go.  
You can work with OSC in Mac, Win, Linux (including Raspberry Pi).  

## Features  

### Type  
- int32  
- float32  
- string  
The other types are NOT supported.  

### Messages or Bundles  
Only OSC Messages are implemented.  
OSC Bundles are NOT supported.  

## Install  

If you're using Go  
```bash
go get github.com/AkiyukiOkayasu/osc
```
to install it.  

If you're not using Go and only need gosc (command line OSC tool) download the zip from [Release](AkiyukiOkayasu/osc/releases/latest/download/gosc.zip).  

  
## How to use  

### Send  
```Go
package main

import (
	"log"

	"github.com/AkiyukiOkayasu/osc"
)

const (
	ip      string = "127.0.0.1" // You can also use "localhost"
	port    int    = 8080
	address string = "/test"
)

func main() {
	sender := osc.NewSender(ip, port)
	message := osc.NewMessage(address)

	// Add OSC arguments by type
	message.AddInt(123)
	message.AddFloat(3.14)
	message.AddString("foo")

	// Send OSC
	if err := sender.Send(message); err != nil {
		log.Fatalln(err)
	}
}

```

### Receive  
```Go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AkiyukiOkayasu/osc"
)

const port int = 8080

func main() {
	// context to cancel OSC receiving gorouting
	c := context.Background()
	ctx, cancel := context.WithCancel(c)

	// OSC handler for /foo
	mux := osc.NewServeMux()
	mux.Handle("/foo", func(m *osc.Message) {
		fmt.Printf("OSC Address: %s\n", m.Address)
		for _, a := range m.Arguments {
			switch a.Type() {
			case 'i':
				if v, ok := a.Int(); ok {
					fmt.Printf("Foo Int: %d\n", v)
				}
			case 'f':
				if v, ok := a.Float(); ok {
					fmt.Printf("Foo Float: %3f\n", v)
				}
			case 's':
				if v, ok := a.String(); ok {
					fmt.Printf("Foo String: %s\n", v)
				}
			default:
				fmt.Printf("Unexpected type: %v\n", a.Type())
			}
		}
	})

	// Another OSC handler for /bar
	mux.Handle("/bar", func(m *osc.Message) {
		fmt.Printf("OSC Address: %s\n", m.Address)
		for _, a := range m.Arguments {
			switch a.Type() {
			case 'i':
				if v, ok := a.Int(); ok {
					fmt.Printf("Bar Int: %d\n", v)
				}
			case 'f':
				if v, ok := a.Float(); ok {
					fmt.Printf("Bar Float: %3f\n", v)
				}
			case 's':
				if v, ok := a.String(); ok {
					fmt.Printf("Bar String: %s\n", v)
				}
			default:
				fmt.Printf("Unexpected type: %v\n", a.Type())
			}
		}
	})

	r := osc.NewReceiver(port, *mux)
	go r.Receive(ctx) // Start OSC receiving

	// Something to do...
	time.Sleep(30 * time.Second) // Sleep 30 seconds

	cancel() // Stop receiving OSC
}

```

## Command line OSC tool  

gosc is a command line tool to send and receive OSC in the terminal.  

### Send  
```bash
gosc send localhost 8080 /test 123 3.14 foo
```
Any number of OSCC arguments can be added.  
Type is automatically determined (int32, float32, string only).  

### Receive  
```bash
gosc receive 8080
```
Ctrl+C to exit receiving OSC.  

## Contribution  

Contributions are welcome.  
Please send me a pull request if you have any ideas.  

## LICENSE  

Apache 2.0  
See [LICENSE](LICENSE)  
