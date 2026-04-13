package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func parseBankLine(line string) (Bank, error) {
	parts := strings.Split(line, ",")
	if len(parts) < 3 {
		return Bank{}, fmt.Errorf("недостаточно полей в строке: %s", line)
	}
	binFrom, err := strconv.Atoi(parts[1])
	if err != nil {
		return Bank{}, err
	}
	binTo, err := strconv.Atoi(parts[2])
	if err != nil {
		return Bank{}, err
	}
	return Bank{
		Name:    parts[0],
		BinFrom: binFrom,
		BinTo:   binTo,
	}, nil
}

func loadBankData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	banks := make([]Bank, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank, err := parseBankLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}
	return banks, scanner.Err()
}

func main() {
	banks, err := loadBankData("banks.txt")
	if err != nil {
		slog.Error("ошибка загрузки банков", "err", err)
		os.Exit(1)
	}
	fmt.Println("Загружено банков:", len(banks))
	fmt.Println(banks)
}
