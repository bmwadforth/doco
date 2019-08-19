package main

import (
	"doco/src/pdf"
	"fmt"
)

func main() {
	header := pdf.Header{Version: pdf.Latest}
	fmt.Println(header.ToBytes())
	fmt.Println(string(header.ToBytes()))
}
