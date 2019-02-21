package main

import (
	"doco/Doco"
	"fmt"
	"log"
)

func main(){
	pdf := Doco.New(Doco.A4)

	pdf.Write("Hello World!")


	saveError := pdf.Save("./myPdf.pdf")
	if saveError != nil {
		log.Fatal(fmt.Errorf("an error occurred while saving PDF - %v", saveError))
	}

	log.Println(string(pdf.Output()))
}
