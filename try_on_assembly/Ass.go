package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Команды УВМ
const (
	CMD_LOAD  = 0x01
	CMD_ADD   = 0x02
	CMD_STORE = 0x03
	CMD_JUMP  = 0x04
	CMD_HALT  = 0x05
)

// SimpleAssembler представляет собой простой ассемблер.
type SimpleAssembler struct {
	machineCode []byte
	log         map[string]interface{}
}

// NewSimpleAssembler создает новый экземпляр SimpleAssembler.
func NewSimpleAssembler() *SimpleAssembler {
	return &SimpleAssembler{
		machineCode: []byte{},
		log:         make(map[string]interface{}),
	}
}

// Assemble принимает путь к файлу с ассемблерным кодом и генерирует бинарный файл.
func (sa *SimpleAssembler) Assemble(inputPath string, outputPath string, logPath string) error {
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") {
			continue // Игнорировать пустые строки и комментарии
		}
		if err := sa.processLine(line, lineNumber); err != nil {
			return err
		}
	}

	// Запись бинарного файла
	if err := ioutil.WriteFile(outputPath, sa.machineCode, 0644); err != nil {
		return err
	}

	// Запись лога в формате JSON
	logData, err := json.Marshal(sa.log)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(logPath, logData, 0644)
}

// processLine обрабатывает одну строку ассемблерного кода.
func (sa *SimpleAssembler) processLine(line string, lineNumber int) error {
	re := regexp.MustCompile(`(\w+)\s+(\w+),?\s*(\w+)?`)
	matches := re.FindStringSubmatch(line)
	if matches != nil {
		instruction, operand1, operand2 := matches[1], matches[2], matches[3]
		var opcode byte
		switch instruction {
		case "LOAD":
			opcode = CMD_LOAD
			sa.machineCode = append(sa.machineCode, opcode)
			sa.machineCode = append(sa.machineCode, byte(parseOperand(operand1)))
			sa.machineCode = append(sa.machineCode, byte(parseValue(operand2)))
		case "ADD":
			opcode = CMD_ADD
			sa.machineCode = append(sa.machineCode, opcode)
			sa.machineCode = append(sa.machineCode, byte(parseOperand(operand1)))
			sa.machineCode = append(sa.machineCode, byte(parseOperand(operand2)))
		case "STORE":
			opcode = CMD_STORE
			sa.machineCode = append(sa.machineCode, opcode)
			sa.machineCode = append(sa.machineCode, byte(parseOperand(operand1)))
			sa.machineCode = append(sa.machineCode, byte(parseValue(operand2)))
		case "JUMP":
			opcode = CMD_JUMP
			sa.machineCode = append(sa.machineCode, opcode)
			sa.machineCode = append(sa.machineCode, byte(parseValue(operand1)))
		case "HALT":
			opcode = CMD_HALT
			sa.machineCode = append(sa.machineCode, opcode)
		default:
			return fmt.Errorf("неизвестная инструкция: %s", instruction)
		}
		sa.log[fmt.Sprintf("line_%d", lineNumber)] = fmt.Sprintf("%s %s, %s", instruction, operand1, operand2)
	} else {
		return fmt.Errorf("ошибка синтаксиса: %s", line)
	}
	return nil
}

// parseOperand преобразует операнд в число.
func parseOperand(operand string) int {
	if strings.HasPrefix(operand, "R") {
		var regNum int
		fmt.Sscanf(operand[1:], "%d", &regNum)
		return regNum
	}
	return 0
}

// parseValue преобразует значение в число.
func parseValue(value string) int {
	var intValue int
	fmt.Sscanf(value, "%d", &intValue)
	return intValue
}

func maiin() {
	if len(os.Args) < 4 {
		fmt.Println("Использование: go run main.go <input.asm> <output.bin> <log.json>")
		return
	}

	assembler := NewSimpleAssembler()
	if err := assembler.Assemble(os.Args[1], os.Args[2], os.Args[3]); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Ассемблирование завершено успешно.")
}
