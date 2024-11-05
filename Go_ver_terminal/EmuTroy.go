package main

import (
	"archive/tar"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Config struct {
	Username    string
	Hostname    string
	VfsPath     string
	LogFilePath string
}

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

func (vfs *VirtualFileSystem) ChangeOwner(fileName string, newOwner string) {
	// В реальной системе вы бы изменили владельца файла
	// Здесь просто выводим информацию
	fmt.Printf("Изменение владельца файла %s на %s\n", fileName, newOwner)
}

func (vfs *VirtualFileSystem) FindFile(fileName string) {
	found := false
	for name := range vfs.files {
		if strings.Contains(name, fileName) {
			fmt.Println(name)
			found = true
		}
	}
	if !found {
		fmt.Printf("Файл не найден: %s\n", fileName)
	}
}

func logAction(logFilePath, username, action string) error {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{time.Now().Format(time.RFC3339), username, action}
	return writer.Write(record)
}

func terminalEmulator(vfs *VirtualFileSystem, config Config) {
	var command string
	fmt.Printf("%s@%s:~$ ", config.Username, config.Hostname)

	for {
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch {
		case command == "exit":
			return
		case command == "ls":
			vfs.ListFiles()
			logAction(config.LogFilePath, config.Username, "ls")
		case strings.HasPrefix(command, "chown "):
			parts := strings.Fields(command)
			if len(parts) < 3 {
				fmt.Println("Использование: chown <новый_владелец> <файл>")
				continue
			}
			newOwner := parts[1]
			fileName := parts[2]
			vfs.ChangeOwner(fileName, newOwner)
			logAction(config.LogFilePath, config.Username, fmt.Sprintf("chown %s %s", newOwner, fileName))
		case strings.HasPrefix(command, "find "):
			fileName := strings.TrimSpace(command[5:])
			vfs.FindFile(fileName)
			logAction(config.LogFilePath, config.Username, fmt.Sprintf("find %s", fileName))
		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
		}
	}
}

func loadConfig(configFilePath string) (Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Config{}, err
	}

	if len(records) < 1 || len(records[0]) < 4 {
		return Config{}, fmt.Errorf("недостаточно данных в конфигурационном файле")
	}

	return Config{
		Username:    records[0][0],
		Hostname:    records[0][1],
		VfsPath:     records[0][2],
		LogFilePath: records[0][3],
	}, nil
}

func main() {
	configFilePath := "E:\\Emu_ter_go\\Config\\Go_ver_terminal\\config.csv"  // Путь к конфигурационному файлу
	tarFilePath := "E:\\Emu_ter_go\\Config\\Go_ver_terminal\\filesystem.tar" // Путь к архиву с файлами

	// Загрузка конфигурации
	config, err := loadConfig(configFilePath)
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v\n", err)
		return
	}

	// Инициализация виртуальной файловой системы
	vfs, err := NewVirtualFileSystem(tarFilePath)
	if err != nil {
		fmt.Printf("Ошибка загрузки виртуальной файловой системы: %v\n", err)
		return
	}

	// Запуск эмулятора терминала
	terminalEmulator(vfs, config)
}
