package main

import (
	"fmt"
	"log"
	"os"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func main() {
	//bank := Bank{}
	//lanarbank := Bank{Name: "Lunar Bank", BinFrom: 400000, BinTo: 499999}
	//fmt.Println(bank)
	//fmt.Println(lanarbank)

	data, err := os.ReadFile("banks.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)
	fmt.Println("Содержимое файла banks.txt:")
	fmt.Println(content)
}
