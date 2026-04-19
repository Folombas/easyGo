package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func readInt(scanner *bufio.Scanner, prompt string) int {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		color.Red("Ошибка: введите целое число!")
	}
}

func readFloat(scanner *bufio.Scanner, prompt string) float64 {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		// заменяем запятую на точку для удобства
		input = strings.ReplaceAll(input, ",", ".")
		val, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return val
		}
		color.Red("Ошибка: введите число (целое или дробное)!")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	color.Cyan("╔════════════════════════════════════════════╗")
	color.Cyan("║  🧮 GO КАЛЬКУЛЯТОР                        ║")
	color.Cyan("║  Практикуем арифметику в Go!              ║")
	color.Cyan("╚════════════════════════════════════════════╝")
	fmt.Println()

	// Ввод двух целых чисел
	color.Green("Введите два целых числа:")
	num1 := readInt(scanner, "  num1 = ")
	num2 := readInt(scanner, "  num2 = ")
	fmt.Println()

	// Базовые операции
	color.Yellow("┌────────────────────────────────────────┐")
	color.Yellow("│  БАЗОВЫЕ АРИФМЕТИЧЕСКИЕ ОПЕРАЦИИ      │")
	color.Yellow("└────────────────────────────────────────┘")

	sum := num1 + num2
	color.Green("➕ Сложение:")
	fmt.Printf("  %d + %d = %d\n\n", num1, num2, sum)

	diff := num1 - num2
	color.Green("➖ Вычитание:")
	fmt.Printf("  %d - %d = %d\n\n", num1, num2, diff)

	product := num1 * num2
	color.Green("✖️ Умножение:")
	fmt.Printf("  %d * %d = %d\n\n", num1, num2, product)

	if num2 != 0 {
		quotient := num1 / num2
		remainder := num1 % num2
		color.Green("➗ Деление и остаток:")
		fmt.Printf("  %d / %d = %d (целочисленное)\n", num1, num2, quotient)
		fmt.Printf("  %d %% %d = %d (остаток)\n\n", num1, num2, remainder)
	} else {
		color.Red("  Деление на ноль невозможно!\n\n")
	}

	// Ввод двух дробных чисел
	color.Green("Теперь введите два дробных числа (разделитель точка или запятая):")
	float1 := readFloat(scanner, "  float1 = ")
	float2 := readFloat(scanner, "  float2 = ")
	fmt.Println()

	color.Yellow("┌────────────────────────────────────────┐")
	color.Yellow("│  ОПЕРАЦИИ С ПЛАВАЮЩЕЙ ТОЧКОЙ          │")
	color.Yellow("└────────────────────────────────────────┘")

	fmt.Printf("Дано: float1 = %.2f, float2 = %.2f\n\n", float1, float2)

	color.Green("➕ Сложение:")
	fmt.Printf("  %.2f + %.2f = %.2f\n\n", float1, float2, float1+float2)

	color.Green("✖️ Умножение:")
	fmt.Printf("  %.2f * %.2f = %.2f\n\n", float1, float2, float1*float2)

	color.Green("➗ Деление:")
	if float2 != 0 {
		fmt.Printf("  %.2f / %.2f = %.2f\n\n", float1, float2, float1/float2)
	} else {
		color.Red("  Деление на ноль невозможно!\n\n")
	}

	// Возведение в степень (через умножение)
	color.Yellow("┌────────────────────────────────────────┐")
	color.Yellow("│  ВОЗВЕДЕНИЕ В СТЕПЕНЬ                 │")
	color.Yellow("└────────────────────────────────────────┘")

	base := 5
	square := base * base
	cube := base * base * base

	fmt.Printf("Число: %d\n", base)
	fmt.Printf("  Квадрат: %d² = %d\n", base, square)
	fmt.Printf("  Куб: %d³ = %d\n\n", base, cube)

	// Сравнение чисел
	color.Yellow("┌────────────────────────────────────────┐")
	color.Yellow("│  ОПЕРАЦИИ СРАВНЕНИЯ                   │")
	color.Yellow("└────────────────────────────────────────┘")

	fmt.Printf("Сравниваем %d и %d:\n\n", num1, num2)

	isGreater := num1 > num2
	isLess := num1 < num2
	isEqual := num1 == num2

	color.Green("  %d > %d = %v\n", num1, num2, isGreater)
	color.Green("  %d < %d = %v\n", num1, num2, isLess)
	color.Green("  %d == %d = %v\n", num1, num2, isEqual)
	fmt.Println()

	// Инкремент и декремент
	color.Yellow("┌────────────────────────────────────────┐")
	color.Yellow("│  ИНКРЕМЕНТ И ДЕКРЕМЕНТ                │")
	color.Yellow("└────────────────────────────────────────┘")

	counter := 10
	fmt.Printf("Начальное значение: counter = %d\n", counter)

	counter++
	fmt.Printf("После counter++ : counter = %d\n", counter)

	counter--
	fmt.Printf("После counter-- : counter = %d\n", counter)

	counter += 5
	fmt.Printf("После counter += 5 : counter = %d\n", counter)

	counter -= 3
	fmt.Printf("После counter -= 3 : counter = %d\n\n", counter)

	// Финальное сообщение
	color.Magenta("╔════════════════════════════════════════════╗")
	color.Magenta("║  ✅ КАЛЬКУЛЯТОР ЗАВЕРШИЛ РАБОТУ!          ║")
	color.Magenta("║  Ты отлично справился с практикой! 🎉     ║")
	color.Magenta("╚════════════════════════════════════════════╝")
}