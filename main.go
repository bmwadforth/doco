package main

import (
	"doco/Doco"
	"fmt"
	"log"
)

//TODO:
//Function To Create New Document TypeCatalog
//Function To Create New TypePage Tree
//Function To Create TypePage & Link It To TypePage Tree
//Use Cross-reference streams rather that cross-reference table
//Implement encryption
//

func main(){
	pdf := Doco.New(Doco.A4)
	/*
	val, err := src.ConvertUnit(500, src.UnitPixels, src.Millimeters)
	if err != nil {
		log.Fatalf("An Error Occurred: %v", err)
	}
	fmt.Printf("Conversion From 500 UnitPixels To %f Millimeters\n", val)

	width, height := pdf.GetDimensions()
	fmt.Printf("Dimensions Of Core TypePage. Height: %d \t Width: %d\n", height, width)

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
