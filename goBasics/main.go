package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// 🎨 Красочное приветствие
	color.Cyan("╔════════════════════════════════════════════╗")
	color.Cyan("║  🚀 ДОБРО ПОЖАЛОВАТЬ В GO BASICS!         ║")
	color.Cyan("║  Изучаем Go легко и интересно!            ║")
	color.Cyan("╚════════════════════════════════════════════╝")
	fmt.Println()

	// 📚 Урок 1: Переменные
	color.Green("📖 УРОК 1: ПЕРЕМЕННЫЕ")
	color.Yellow("────────────────────────────────────────")
	
	// Способ 1: Короткое объявление :=
	name := "Alex"
	age := 25
	fmt.Printf("👤 Имя: %s, Возраст: %d лет\n", name, age)
	
	// Способ 2: var с типом
	var height int = 180
	var isStudent bool = true
	fmt.Printf("📏 Рост: %d см, Студент: %v\n", height, isStudent)
	
	// Способ 3: var без явного типа
	var city = "Moscow"
	fmt.Printf("🏙️ Город: %s\n", city)
	fmt.Println()

	// 📚 Урок 2: Типы данных
	color.Green("📖 УРОК 2: ТИПЫ ДАННЫХ")
	color.Yellow("────────────────────────────────────────")
	
	var intNum int = 42
	var floatNum float64 = 3.14159
	var text string = "Go is awesome!"
	var isFun bool = true
	
	fmt.Printf("🔢 int: %d (тип: %T)\n", intNum, intNum)
	fmt.Printf("🔬 float64: %.2f (тип: %T)\n", floatNum, floatNum)
	fmt.Printf("📝 string: %s (тип: %T)\n", text, text)
	fmt.Printf("✅ bool: %v (тип: %T)\n", isFun, isFun)
	fmt.Println()

	// 📚 Урок 3: Арифметика
	color.Green("📖 УРОК 3: АРИФМЕТИЧЕСКИЕ ОПЕРАЦИИ")
	color.Yellow("────────────────────────────────────────")
	
	a, b := 20, 7
	fmt.Printf("Даны числа: a = %d, b = %d\n\n", a, b)
	
	fmt.Printf("➕ Сложение: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("➖ Вычитание: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("✖️ Умножение: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("➗ Деление: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("📊 Остаток: %d %% %d = %d\n", a, b, a%b)
	fmt.Println()

	// 📚 Урок 4: Строки
	color.Green("📖 УРОК 4: РАБОТА СО СТРОКАМИ")
	color.Yellow("────────────────────────────────────────")
	
	greeting := "Привет"
	target := "Go"
	message := greeting + ", " + target + "!"
	
	fmt.Printf("🎯 Строка 1: %s\n", greeting)
	fmt.Printf("🎯 Строка 2: %s\n", target)
	fmt.Printf("🎯 Конкатенация: %s\n", message)
	fmt.Printf("📏 Длина строки: %d символов\n", len(message))
	fmt.Println()

	// 📚 Урок 5: Константы
	color.Green("📖 УРОК 5: КОНСТАНТЫ")
	color.Yellow("────────────────────────────────────────")
	
	const Pi = 3.14159
	const MaxUsers = 1000
	const AppName = "GoApp"
	
	fmt.Printf("🔵 Pi = %.2f\n", Pi)
	fmt.Printf("🔵 MaxUsers = %d\n", MaxUsers)
	fmt.Printf("🔵 AppName = %s\n", AppName)
	fmt.Println()

	// 🎯 Финальное сообщение
	color.Magenta("╔════════════════════════════════════════════╗")
	color.Magenta("║  ✅ УРОК 1 ЗАВЕРШЁН!                      ║")
	color.Magenta("║  Ты изучил основы переменных и типов!     ║")
	color.Magenta("║  Продолжай в том же духе! 🚀              ║")
	color.Magenta("╚════════════════════════════════════════════╝")
	fmt.Println()
	
	color.Cyan("💡 Совет: Попробуй изменить значения переменных и запусти снова!")
}
