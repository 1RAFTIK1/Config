// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// // SimpleAssembler представляет собой простой ассемблер.
// type SimpleAssembler struct {
// 	instructions map[string]byte
// 	machineCode  []byte
// }

// // NewSimpleAssembler создает новый экземпляр SimpleAssembler.
// func NewSimpleAssembler() *SimpleAssembler {
// 	return &SimpleAssembler{
// 		instructions: map[string]byte{
// 			"MOV": 0x01,
// 			"ADD": 0x02,
// 			"SUB": 0x03,
// 		},
// 		machineCode: []byte{},
// 	}
// }

// // Assemble принимает строку с ассемблерным кодом и генерирует машинный код.
// func (sa *SimpleAssembler) Assemble(assemblyCode string) error {
// 	lines := strings.Split(assemblyCode, "\n")
// 	for _, line := range lines {
// 		line = strings.TrimSpace(line)
// 		if line == "" || strings.HasPrefix(line, ";") {
// 			continue // Игнорировать пустые строки и комментарии
// 		}
// 		if err := sa.processLine(line); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // processLine обрабатывает одну строку ассемблерного кода.
// func (sa *SimpleAssembler) processLine(line string) error {
// 	re := regexp.MustCompile(`(\w+)\s+(\w+),\s*(\w+)`)
// 	matches := re.FindStringSubmatch(line)
// 	if matches != nil {
// 		instruction, operand1, operand2 := matches[1], matches[2], matches[3]
// 		if opcode, exists := sa.instructions[instruction]; exists {
// 			sa.machineCode = append(sa.machineCode, opcode)
// 			sa.machineCode = append(sa.machineCode, operandToByte(operand1))
// 			sa.machineCode = append(sa.machineCode, operandToByte(operand2))
// 		} else {
// 			return fmt.Errorf("неизвестная инструкция: %s", instruction)
// 		}
// 	} else {
// 		return fmt.Errorf("ошибка синтаксиса: %s", line)
// 	}
// 	return nil
// }

// // operandToByte преобразует операнд в байт (предполагается, что операнды - это R1, R2 и т.д.).
// func operandToByte(operand string) byte {
// 	if len(operand) < 2 || operand[0] != 'R' {
// 		return 0 // Возвращаем 0 для неверного формата
// 	}
// 	var regNum byte
// 	fmt.Sscanf(operand[1:], "%d", &regNum) // Преобразуем R1 в 1
// 	return regNum
// }

// // GetMachineCode возвращает сгенерированный машинный код.
// func (sa *SimpleAssembler) GetMachineCode() []byte {
// 	return sa.machineCode
// }

// func main() {
// 	assembler := NewSimpleAssembler()
// 	assemblyCode := `
// 	; Простой пример ассемблера
// 	MOV R1, R2
// 	ADD R1, R3
// 	SUB R2, R1
// 	`
// 	if err := assembler.Assemble(assemblyCode); err != nil {
// 		fmt.Println("Ошибка:", err)
// 		return
// 	}
// 	fmt.Println("Сгенерированный машинный код:", assembler.GetMachineCode())
// }
