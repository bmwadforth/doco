package main

import (
	"doco/src"
	"fmt"
	"log"
)

//TODO:
//Support Linearized PDF
//Look at cross-reference streams

func main(){
	pdf := src.New(src.A4)
	/*
	val, err := src.ConvertUnit(500, src.Pixels, src.Millimeters)
	if err != nil {
		log.Fatalf("An Error Occurred: %v", err)
	}
	fmt.Printf("Conversion From 500 Pixels To %f Millimeters\n", val)

	width, height := pdf.GetDimensions()
	fmt.Printf("Dimensions Of Doco Page. Height: %d \t Width: %d\n", height, width)

	*/

	saveErr := pdf.Save("./myPdf.pdf")
	if saveErr != nil {
		fmt.Printf("An Error Occurred: %v\n", saveErr)
	}

	/*
	err := pdf.Build()
	if err != nil {
		log.Fatal(fmt.Sprintf("An Error Occurred: %v", err))
	}*/

	log.Println(pdf.Output())
}
