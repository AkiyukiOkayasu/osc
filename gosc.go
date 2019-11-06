package main

import (
	"fmt"
)

func main() {
	// fmt.Println(len("abc"))
	hoge("fugafruga")
}

func hoge(args ...interface{}) {
	for _, arg := range args {
		switch t := arg.(type) {
		case int32, int64:
			fmt.Println("int")
			fmt.Println(arg)

		case float32, float64:
			fmt.Println("float")
			fmt.Println(arg)

		case string:
			fmt.Println("string")
			fmt.Println(arg)

		default:
			fmt.Println("default")
			fmt.Println(t)
		}
	}
}
