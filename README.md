# osc  
OSC (Open Sound Control) package for Go.  
Implemented in pure Go.  
A portion of The [Open Sound Control 1.0 Specification](http://opensoundcontrol.org/spec-1_0) has been implemented.  

## Type  
- int32  
- float32  
- string  
The other types are NOT supported.  

## Message OR Bundle  
Only OSC Messages are implemented.  
OSC Bundles are NOT supported.  
  
## How to use    


## Command line OSC sender, receiver  
gosc is
#### Send  
```bash
gosc send localhost 8080 /test 123 3.14 foo
```

#### Receive  
```bash
gosc receive 8080
```

### Contribution  
Contributions are welcome.  
Please send me a pull request if you have any ideas.  


### LICENSE  
Apache 2.0
See [LICENSE]
