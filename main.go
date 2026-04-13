package main

import (
	"bufio"
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

	file, err := os.Open("banks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	//content := string(data)
	//fmt.Println("Содержимое файла banks.txt:")
	//fmt.Println(content)
}
