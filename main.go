package main

import (
	"doco/src"
	"fmt"
	"log"
)

//TODO:
//Function To Create New Document Catalog
//Function To Create New Page Tree
//Function To Create Page & Link It To Page Tree
//Use Cross-reference streams rather that cross-reference table
//Implement encryption
//

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
