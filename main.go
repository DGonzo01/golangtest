package main

import "fmt"

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func main() {
	bank := Bank{}
	lanarbank := Bank{Name: "Lunar Bank", BinFrom: 400000, BinTo: 499999}
	fmt.Println(bank)
	fmt.Println(lanarbank)
}
