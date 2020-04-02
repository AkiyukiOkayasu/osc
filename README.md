# go-osc  
OSC(Open Sound Control) package for golang.
  
### Example gosc  
Simple command line OSC sender and receiver.  

#### Send  
```bash
gosc send localhost 8080 /test 123 3.14 foo
```

#### Receive  
```bash
gosc receive 8080
```
