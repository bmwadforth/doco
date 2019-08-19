package main

import (
	"doco/src/file"
	"fmt"
)

func main() {
	header := file.Header{Version: file.Latest}
	fmt.Println(header.ToBytes())
	fmt.Println(string(header.ToBytes()))
}
