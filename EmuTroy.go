package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type VirtualFileSystem struct {
	files map[string][]byte
}

func NewVirtualFileSystem(tarFilePath string) (*VirtualFileSystem, error) {
	vfs := &VirtualFileSystem{files: make(map[string][]byte)}
	err := vfs.loadTar(tarFilePath)
	if err != nil {
		return nil, err
	}
	return vfs, nil
}

func (vfs *VirtualFileSystem) loadTar(tarFilePath string) error {
	file, err := os.Open(tarFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	tarReader := tar.NewReader(file)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // Конец архива
		}
		if err != nil {
			return err
		}

		if header.Typeflag == tar.TypeReg { // Только обычные файлы
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, tarReader); err != nil {
				return err
			}
			vfs.files[header.Name] = buf.Bytes()
		}
	}
	return nil
}

func (vfs *VirtualFileSystem) ListFiles() {
	for fileName := range vfs.files {
		fmt.Println(fileName)
	}
}

func (vfs *VirtualFileSystem) CatFile(fileName string) {
	content, exists := vfs.files[fileName]
	if !exists {
		fmt.Printf("Файл не найден: %s\n", fileName)
		return
	}
	fmt.Println(string(content))
}

func terminalEmulator(vfs *VirtualFileSystem) {
	var command string
	fmt.Println("Эмулятор терминала. Введите команду (или 'exit' для выхода):")

	for {
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch {
		case command == "exit":
			return
		case command == "ls":
			vfs.ListFiles()
		case strings.HasPrefix(command, "cat "):
			fileName := strings.TrimSpace(command[4:])
			vfs.CatFile(fileName)
		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <путь_к_tar_файлу>")
		return
	}

	tarFilePath := os.Args[1]
	vfs, err := NewVirtualFileSystem(tarFilePath)
	if err != nil {
		fmt.Printf("Ошибка загрузки виртуальной файловой системы: %s\n", err)
		return
	}

	terminalEmulator(vfs)
}
