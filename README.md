# osc  
OSC (Open Sound Control) package for Go.  
A part of The [Open Sound Control 1.0 Specification](http://opensoundcontrol.org/spec-1_0) has been implemented in pure Go.  

## Type  
- int32  
- float32  
- string  
The other types are NOT supported.  

## Messages or Bundles  
Only OSC Messages are implemented.  
OSC Bundles are NOT supported.  
  
## How to use  
### Send  
```Go
package main

import (
	"log"

	"github.com/AkiyukiOkayasu/osc"
)

func main() {
	ip := "127.0.0.1" // You can also use "localhost"
	port := 8080
	address := "/test"

	sender := osc.NewSender(ip, port)
	message := osc.NewMessage(address)

	message.AddInt(123)
	message.AddFloat(3.14)
	message.AddString("foo")

	if err := sender.Send(message); err != nil {
		log.Fatalln(err)
	}
}

```

## Command line OSC sender, receiver  
gosc is
### Send  
```bash
gosc send localhost 8080 /test 123 3.14 foo
```

### Receive  
```bash
gosc receive 8080
```

## Contribution  
Contributions are welcome.  
Please send me a pull request if you have any ideas.  


## LICENSE  
Apache 2.0  
See [LICENSE](LICENSE)  
