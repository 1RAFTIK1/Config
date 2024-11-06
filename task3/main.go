package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: go run main.go <input.conf> <output.toml>")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	parser := NewParser()
	if err := parser.Parse(inputPath, outputPath); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("Парсинг завершен успешно.")
}
