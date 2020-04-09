# osc  
OSC(Open Sound Control) package for Go. Implemented in pure Go.  

## Implemented  
- OSC Message  
- int32  
- float32  
- osc string  

## NOT implemented  
- OSC Bundle  
- time tag
- blob  

  
## Example  
### gosc: command line OSC sender, receiver  

#### Send  
```bash
gosc send localhost 8080 /test 123 3.14 foo
```

#### Receive  
```bash
gosc receive 8080
```
