// Файл: Go110/gocommands/main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║  🛠️  ДЕМОНСТРАЦИЯ РАБОТЫ КОМАНД GO (go command)              ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Этот проект показывает, как использовать основные подкоманды Go.")
	fmt.Println("Выполните следующие команды в этой директории:\n")

	commands := []struct {
		cmd  string
		desc string
	}{
		{"go run .", "Запустить программу (то, что вы сейчас видите)"},
		{"go build -o demo", "Скомпилировать бинарный файл 'demo' (или demo.exe на Windows)"},
		{"./demo", "Запустить скомпилированный бинарник (или demo.exe)"},
		{"go test -v", "Запустить тесты (файл main_test.go)"},
		{"go fmt ./...", "Отформатировать весь код в соответствии со стандартами Go"},
		{"go vet", "Выполнить статический анализ кода"},
		{"go mod init gocommands", "Создать go.mod (если ещё не создан)"},
		{"go mod tidy", "Подтянуть недостающие зависимости и удалить неиспользуемые"},
		{"go install .", "Установить бинарник в $GOPATH/bin (или ~/go/bin)"},
	}

	for _, c := range commands {
		fmt.Printf("  • %-25s — %s\n", c.cmd, c.desc)
	}

	fmt.Println()
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Println("🧪 Дополнительно: попробуйте 'go help <command>' для справки.")
	fmt.Println("   Например: go help build, go help test, go help mod")
	fmt.Println("════════════════════════════════════════════════════════════════")

	// Вызов функции из demo.go
	DemoFunction()
}
