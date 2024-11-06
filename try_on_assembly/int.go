package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Виртуальная машина
type VirtualMachine struct {
	memory [256]int
	reg    [4]int
	pc     int // Программный счетчик
}

// NewVirtualMachine создает новый экземпляр виртуальной машины.
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{}
}

// Execute выполняет команды из бинарного файла.
func (vm *VirtualMachine) Execute(binaryPath string, resultPath string, memoryRange []int) error {
	data, err := ioutil.ReadFile(binaryPath)
	if err != nil {
		return err
	}

	for vm.pc < len(data) {
		opcode := data[vm.pc]
		vm.pc++

		switch opcode {
		case CMD_LOAD:
			reg := data[vm.pc]
			vm.pc++
			value := int(data[vm.pc])
			vm.pc++
			vm.reg[reg] = value
		case CMD_ADD:
			reg1 := data[vm.pc]
			vm.pc++
			reg2 := data[vm.pc]
			vm.pc++
			vm.reg[reg1] += vm.reg[reg2]
		case CMD_STORE:
			reg := data[vm.pc]
			vm.pc++
			address := int(data[vm.pc])
			vm.pc++
			vm.memory[address] = vm.reg[reg]
		case CMD_JUMP:
			address := int(data[vm.pc])
			vm.pc = address
		case CMD_HALT:
			return nil
		default:
			return fmt.Errorf("неизвестная команда: %d", opcode)
		}
	}

	// Запись результатов в файл
	resultData := make(map[string]int)
	for _, address := range memoryRange {
		resultData[fmt.Sprintf("memory_%d", address)] = vm.memory[address]
	}

	resultJSON, err := json.Marshal(resultData)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(resultPath, resultJSON, 0644)
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Использование: go run main.go <binary.bin> <result.json> <start_addr> <end_addr>")
		return
	}

	startAddr := 0
	endAddr := 0
	fmt.Sscanf(os.Args[3], "%d", &startAddr)
	fmt.Sscanf(os.Args[4], "%d", &endAddr)

	vm := NewVirtualMachine()
	if err := vm.Execute(os.Args[1], os.Args[2], generateMemoryRange(startAddr, endAddr)); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Выполнение завершено успешно.")
}

// generateMemoryRange создает диапазон адресов памяти для вывода.
func generateMemoryRange(start int, end int) []int {
	rangeSlice := []int{}
	for i := start; i <= end; i++ {
		rangeSlice = append(rangeSlice, i)
	}
	return rangeSlice
}
