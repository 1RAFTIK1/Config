package main

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	parser := NewParser()

	// Создание временного файла для тестирования
	inputFile, err := os.CreateTemp("", "test.conf")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(inputFile.Name())

	outputFile, err := os.CreateTemp("", "test.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outputFile.Name())

	// Запись тестовых данных
	testData := `
// Это комментарий
var A := 10
var B := "Hello"
[my_dict => A, B]
`
	if _, err := inputFile.WriteString(testData); err != nil {
		t.Fatal(err)
	}
	inputFile.Close()

	// Парсинг
	if err := parser.Parse(inputFile.Name(), outputFile.Name()); err != nil {
		t.Fatalf("Ошибка парсинга: %v", err)
	}

	// Проверка результата
	expected := "[my_dict]\nA = 10\nB = Hello\n"
	outputData, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if string(outputData) != expected {
		t.Errorf("Ожидалось:\n%s\nНо получено:\n%s", expected, string(outputData))
	}
}
