package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Ошибки парсинга
var (
	ErrInvalidSyntax     = fmt.Errorf("неверный синтаксис")
	ErrUndefinedVariable = fmt.Errorf("неопределенная переменная")
)

// Parser представляет собой парсер конфигурационного языка.
type Parser struct {
	variables map[string]string
}

// NewParser создает новый экземпляр Parser.
func NewParser() *Parser {
	return &Parser{
		variables: make(map[string]string),
	}
}

// Parse читает входной файл и генерирует выходной TOML.
func (p *Parser) Parse(inputPath string, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue // Игнорируем комментарии
		}

		if strings.Contains(line, "|#") {
			// Многострочный комментарий
			continue
		}

		if strings.HasPrefix(line, "var") {
			if err := p.handleVariable(line); err != nil {
				return err
			}
			continue
		}

		if strings.Contains(line, "=>") {
			if err := p.handleDictionary(line, &output); err != nil {
				return err
			}
			continue
		}

		return fmt.Errorf("неизвестная конструкция: %s", line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Запись в выходной файл
	return os.WriteFile(outputPath, []byte(output.String()), 0644)
}

// handleVariable обрабатывает объявление переменной.
func (p *Parser) handleVariable(line string) error {
	re := regexp.MustCompile(`var\s+([_A-Z][_a-zA-Z0-9]*)\s*:=\s*(.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return ErrInvalidSyntax
	}

	varName := matches[1]
	value := strings.TrimSpace(matches[2])
	p.variables[varName] = value
	return nil
}

// handleDictionary обрабатывает словарь.
func (p *Parser) handleDictionary(line string, output *strings.Builder) error {
	re := regexp.MustCompile(`\[(.+?)\s*=>\s*(.+?)\]`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return ErrInvalidSyntax
	}

	dictName := strings.TrimSpace(matches[1])
	dictValues := strings.TrimSpace(matches[2])
	values := strings.Split(dictValues, ",")
	output.WriteString(fmt.Sprintf("[%s]\n", dictName))

	for _, v := range values {
		v = strings.TrimSpace(v)
		if strings.HasPrefix(v, "$") {
			varName := strings.TrimPrefix(v, "$")
			if val, exists := p.variables[varName]; exists {
				output.WriteString(fmt.Sprintf("%s = %s\n", varName, val))
			} else {
				return ErrUndefinedVariable
			}
		} else {
			output.WriteString(fmt.Sprintf("%s = %s\n", v, v))
		}
	}

	return nil
}
