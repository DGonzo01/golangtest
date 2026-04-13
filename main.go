package main

import "fmt"

type Bank struct {
	Name    string
	BioFrom int
	BinTo   int
}

func main() {
	bank := Bank{}
	lanarbank := Bank{Name: "Lunar Bank", BioFrom: 400000, BinTo: 499999}
	fmt.Println(bank)
	fmt.Println(lanarbank)
}
