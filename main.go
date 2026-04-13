package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func loadBankData(path string) (Bank, error) {
	parts := strings.Split(path, ",")
	BinTo, err := strconv.Atoi(parts[1])
	if err != nil {
		return Bank{}, err
	}
	BinFrom, err := strconv.Atoi(parts[2])
	if err != nil {
		return Bank{}, err
	}
	return Bank{
		Name:    parts[0],
		BinTo:   BinTo,
		BinFrom: BinFrom,
	}, nil
}

func main() {

	file, err := os.Open("banks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	banks := make([]Bank, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bank, err := loadBankData(line)
		banks = append(banks, bank)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bank)
	}
	fmt.Println("Загружено банков: 5")
	fmt.Println(banks)
}
