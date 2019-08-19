package main

import (
	"doco/src/file"
	"doco/src/pdf"
)

func main() {
	doc, _ := pdf.New(file.Latest)
	bytes, _ := doc.Build()
	println(string(bytes))
}
