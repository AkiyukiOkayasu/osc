package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("abc"))
}

func hoge(args ...interface{}) {
	for _, arg := range args {
		switch t := arg.(type) {
		case int32:
			fmt.Println("int32")

		default:
			fmt.Println("default")
			fmt.Println(t)
		}
	}
}
