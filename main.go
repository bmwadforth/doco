package main

import (
	src "doco/src"
	"fmt"
	"log"
)

//TODO:
//Support Linearized PDF

func main(){
	pdf := src.New(src.A4)
	val, err := src.ConvertUnit(500, src.Pixels, src.Millimeters)
	if err != nil {
		log.Fatalf("An Error Occurred: %v", err)
	}
	fmt.Printf("Conversion From 500 Pixels To %f Millimeters\n", val)

	width, height := pdf.GetDimensions()
	fmt.Printf("Dimensions Of Doco Page. Height: %d \t Width: %d\n", height, width)

	saveErr := pdf.Save("./myPdf.pdf")
	if saveErr != nil {
		fmt.Printf("An Error Occurred: %v\n", saveErr)
	}
}
